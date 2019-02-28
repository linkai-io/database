package main

import (
	"fmt"
	"os"
	"time"

	"github.com/linkai-io/am/pkg/initializers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	appConfig        initializers.AppConfig
	runType          string
	loadBalancerAddr string
)

func init() {
	appConfig.Env = os.Getenv("APP_ENV")
	appConfig.Region = os.Getenv("APP_REGION")
	appConfig.SelfRegister = os.Getenv("APP_SELF_REGISTER")
	appConfig.ServiceKey = "linkai_admin"
	runType = os.Getenv("APP_RUN_TYPE")
}

// main starts the aggregation methods
func main() {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("service", "Aggregation Task").Logger()
	var err error
	var aggregates []string

	_, pool := initializers.DB(&appConfig)

	if runType == "day" {
		aggregates = []string{"am.do_daily_discovered_aggregation", "am.do_daily_seen_aggregation", "am.do_daily_scanned_aggregation"}
	} else {
		aggregates = []string{"am.do_trihourly_discovered_aggregation", "am.do_trihourly_seen_aggregation", "am.do_trihourly_scanned_aggregation"}
	}
	for _, agg := range aggregates {
		var start int
		var end int
		then := time.Now()
		if err = pool.QueryRow(fmt.Sprintf("select * from %s()", agg)).Scan(&start, &end); err != nil {
			log.Fatal().Err(err).Msg("failed to run aggregation functions")
		}
		log.Info().Str("aggregation", agg).TimeDiff("completed_in", time.Now(), then).Msg("aggregation completed")
	}
}
