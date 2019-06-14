package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/linkai-io/am/pkg/filestorage"
	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/am/pkg/webtech"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// WappalyzerURL is the link to the latest apps.json file in the Wappalyzer repo
const WappalyzerURL = "https://raw.githubusercontent.com/AliasIO/Wappalyzer/master/src/apps.json"

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

func main() {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("service", "WebTech Task").Logger()
	var err error

	_, pool := initializers.DB(&appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to initialize db connection")
	}

	appJSON := downloadLatestWappalyzer()
	storage := filestorage.NewS3Storage(appConfig.Env, appConfig.Region)
	if err := storage.Init(); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize s3 storage")
	}

	if err := storage.PutInfraFile(context.Background(), "linkai-infra", appConfig.Env+"/web/apps.json", appJSON); err != nil {
		log.Fatal().Err(err).Msg("failed to put latest wappalyzer apps.json file in storage")
	}

	query := "insert into am.web_techtypes (techname, category_id, category, website, icon) values ($1, $2, $3, $4, $5) on conflict (techname, category) do nothing"
	tx, err := pool.Begin()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create transaction")
	}

	defer tx.Rollback() // safe to call as no-op on success

	if err != nil {
		log.Fatal().Err(err).Msg("unable to create prepared stmt")
	}

	wapp := webtech.NewWappalyzer()
	if err := wapp.Init(appJSON); err != nil {
		log.Fatal().Err(err).Msg("unable to initialize wappalyzer data")
	}

	defs := wapp.AppDefinitions()
	for techName, app := range defs.Apps {
		for i := 0; i < len(app.CatNames); i++ {
			cid, _ := strconv.Atoi(app.Cats[i])
			log.Info().Msgf("%s %d %s %s %s", techName, cid, app.CatNames[i], app.Website, app.Icon)
			if _, err := tx.Exec(query, techName, cid, app.CatNames[i], app.Website, app.Icon); err != nil {
				log.Fatal().Err(err).Msg("failed to insert tech")
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("failed to commit tech updates")
	}
}

func downloadLatestWappalyzer() []byte {
	resp, err := http.Get(WappalyzerURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to download latest wappalyzer apps.json")
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read latest wappalyzer apps.json")
	}
	return data
}
