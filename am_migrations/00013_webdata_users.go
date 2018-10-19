package migration

import (
	"database/sql"
	"os"

	"github.com/linkai-io/am/pkg/secrets"
	"github.com/linkai-io/database/pkg/migration"

	"github.com/pressly/goose"
)

var webUsers = []string{"webdataservice"}

func init() {
	goose.AddMigration(Up00013, Down00013)
}

func Up00013(tx *sql.Tx) error {
	dbsecrets := secrets.NewDBSecrets(os.Getenv("APP_ENV"), os.Getenv("APP_REGION"))
	userMap, err := migration.GetServicePasswords(dbsecrets, webUsers)
	if err != nil {
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

func Down00013(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	for _, service := range webUsers {

		if _, err := tx.Exec("DROP USER " + service); err != nil {
			return err
		}
	}

	return nil
}
