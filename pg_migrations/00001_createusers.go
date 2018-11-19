package migration

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/linkai-io/am/pkg/secrets"

	"github.com/pressly/goose"
)

var users = []string{"bigdataservice", "webdataservice", "eventservice", "orgservice", "userservice", "tagservice", "scangroupservice", "addressservice", "findingsservice"}

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
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

	if os.Getenv("APP_REGION") == "" {
		log.Printf("creating linkai_admin superuser")
		if _, err := tx.Exec("CREATE USER linkai_admin WITH LOGIN SUPERUSER INHERIT ENCRYPTED PASSWORD '" + adminPassword + "'"); err != nil {
			return err
		}
	} else {
		// if region is set, assume AWS
		log.Printf("creating linkai_admin")
		if _, err := tx.Exec("CREATE ROLE linkai_admin WITH LOGIN INHERIT CREATEROLE ENCRYPTED PASSWORD '" + adminPassword + "' "); err != nil {
			return err
		}
		log.Printf("granting linkai_admin rds_superuser role")
		if _, err := tx.Exec("GRANT rds_superuser to linkai_admin"); err != nil {
			return err
		}
		if _, err := tx.Exec("GRANT linkai_admin to postgres"); err != nil {
			return err
		}
		log.Printf("linkai_admin user created")
	}

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

func Down00001(tx *sql.Tx) error {
	var err error
	// This code is executed when the migration is rolled back.
	for _, service := range users {

		if _, err := tx.Exec("DROP USER " + service); err != nil {
			return err
		}
	}
	if _, err := tx.Exec("REVOKE linkai_admin from postgres"); err != nil {
		return err
	}
	if os.Getenv("APP_REGION") == "" {
		_, err = tx.Exec("DROP USER linkai_admin")
	} else {
		_, err = tx.Exec("DROP ROLE linkai_admin")
	}
	return err
}
