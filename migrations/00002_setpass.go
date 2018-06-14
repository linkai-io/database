package migration

import (
	"database/sql"
	"errors"
	"os"

	"github.com/pressly/goose"
)

var (
	secret00002 = os.Getenv("GOOSE_SECRET_00002")
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	if secret00002 == "" {
		return errors.New("GOOSE_SECRET_00002 not set")
	}
	// ugh, unavoidable, CREATE USER can not be executed with a prepared statement.
	_, err := tx.Exec("CREATE USER linkai_writer WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + secret00002 + "'")
	return err
}

func Down00002(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP USER linkai_writer")
	return err
}
