package main

import (
	"log"
	"os"

	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/am/pkg/secrets"
)

var (
	appConfig        initializers.AppConfig
	loadBalancerAddr string
)

func init() {
	appConfig.Env = os.Getenv("APP_ENV")
	appConfig.Region = os.Getenv("APP_REGION")
	appConfig.SelfRegister = os.Getenv("APP_SELF_REGISTER")
	appConfig.ServiceKey = "linkai_admin"
}

// main starts the CoordinatorService
func main() {
	var err error
	var orgID int
	var userID int

	_, pool := initializers.DB(&appConfig)

	if err = pool.QueryRow("select organization_id, user_id from am.users where organization_id=(select organization_id from am.organizations where organization_name=$1)", "linkai-system").Scan(&orgID, &userID); err != nil {
		log.Fatalf("failed to get system ids: %v\n", err)
	}
	sec := secrets.NewSecretsCache(appConfig.Env, appConfig.Region)

	if err := sec.SetSystemIDs(orgID, userID); err != nil {
		log.Fatalf("failed to add system ids: %v\n", err)
	}
	log.Printf("system ids added successfully\n")
}
