package migration

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/linkai-io/am/pkg/secrets"

	"github.com/pressly/goose"
)

var users = []string{"customwebflowservice", "integrationservice", "techfingerservice", "typosquatterservice"}

func init() {
	goose.AddMigration(Up00038, Down00038)
}

func Up00038(tx *sql.Tx) error {
	dbsecrets := secrets.NewSecretsCache(os.Getenv("APP_ENV"), os.Getenv("APP_REGION"))
	adminPassword, err := dbsecrets.ServicePassword("linkai_admin")
	userMap, err := getServicePasswords(dbsecrets, users)
	if err != nil {
		return err
	}

	if adminPassword == "" {
		return errors.New("linkai_admin not set")
	}

	log.Printf("Creating users...\n")

	// ugh, unavoidable, CREATE USER can not be executed with a prepared statement.
	for service, passwd := range userMap {
		log.Printf("creating user: %s\n", service)
		if _, err := tx.Exec("CREATE USER " + service + " WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + passwd + "'"); err != nil {
			return err
		}
	}

	return nil
}

func getServicePasswords(dbsecrets *secrets.SecretsCache, users []string) (map[string]string, error) {
	userMap := make(map[string]string, 0)
	for _, user := range users {
		password, err := dbsecrets.ServicePassword(user)
		if password == "" {
			return nil, errors.New("empty password for user: " + user)
		}
		userMap[user] = password
		if err != nil {
			return nil, err
		}
	}
	return userMap, nil
}

func Down00038(tx *sql.Tx) error {
	var err error
	// This code is executed when the migration is rolled back.
	for _, service := range users {

		if _, err := tx.Exec("DROP USER " + service); err != nil {
			return err
		}
	}
	return err
}
