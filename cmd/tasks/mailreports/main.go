package main

import (
	"os"

	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/am/pkg/mail/sesmailer"
	"github.com/linkai-io/database/cmd/tasks/mailreports/reports"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	appConfig initializers.AppConfig
	runType   string
	weekly    bool
)

func init() {
	appConfig.Env = os.Getenv("APP_ENV")
	appConfig.Region = os.Getenv("APP_REGION")
	appConfig.SelfRegister = os.Getenv("APP_SELF_REGISTER")
	appConfig.ServiceKey = "linkai_admin"
	runType = os.Getenv("APP_RUN_TYPE")
	if runType == "weekly" {
		weekly = true
	}
}

func main() {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("service", "Email Reporting Task").Logger()

	mailer := sesmailer.New(appConfig.Env, appConfig.Region)
	if err := mailer.Init(nil); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize mailer")
	}
	reports.GenerateReports(appConfig, weekly, mailer)
}
