package migration

import (
	"database/sql"
	"errors"
	"os"

	"gopkg.linkai.io/v1/repos/am/pkg/secrets"

	"github.com/pressly/goose"
)

var users = []string{"eventservice", "orgservice", "userservice", "tagservice", "scangroupservice", "addressservice", "findingsservice"}

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	dbsecrets := secrets.NewDBSecrets(os.Getenv("APP_ENV"), os.Getenv("APP_REGION"))
	adminPassword, err := dbsecrets.ServicePassword("linkai_admin")
	userMap, err := getServicePasswords(dbsecrets, users)
	if err != nil {
		return err
	}

	if adminPassword == "" {
		return errors.New("linkai_admin not set")
	}

	if _, err := tx.Exec("CREATE USER linkai_admin WITH LOGIN SUPERUSER INHERIT ENCRYPTED PASSWORD '" + adminPassword + "'"); err != nil {
		return err
	}

	// ugh, unavoidable, CREATE USER can not be executed with a prepared statement.
	for service, passwd := range userMap {
		if _, err := tx.Exec("CREATE USER " + service + " WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + passwd + "'"); err != nil {
			return err
		}
	}

	return nil
}

func getServicePasswords(dbsecrets *secrets.DBSecrets, users []string) (map[string]string, error) {
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
	// This code is executed when the migration is rolled back.
	for _, service := range users {
		if _, err := tx.Exec("DROP USER " + service); err != nil {
			return err
		}
	}

	_, err := tx.Exec("DROP USER linkai_admin")
	return err
}
