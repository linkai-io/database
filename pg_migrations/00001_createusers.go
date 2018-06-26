package migration

import (
	"database/sql"
	"errors"
	"os"

	"github.com/pressly/goose"
)

var (
	secret00001_admin        = os.Getenv("GOOSE_SECRET_00001_admin")
	secret00001_jobservice   = os.Getenv("GOOSE_SECRET_00001_jobservice")
	secret00001_inputservice = os.Getenv("GOOSE_SECRET_00001_inputservice")
	secret00001_frontend     = os.Getenv("GOOSE_SECRET_00001_frontend")
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	if secret00001_admin == "" {
		return errors.New("GOOSE_SECRET_00001_admin not set")
	}

	if secret00001_jobservice == "" {
		return errors.New("GOOSE_SECRET_00001_jobservice not set")
	}

	if secret00001_inputservice == "" {
		return errors.New("GOOSE_SECRET_00001_inputservice not set")
	}

	if secret00001_frontend == "" {
		return errors.New("GOOSE_SECRET_00001_frontend not set")
	}

	if _, err := tx.Exec("CREATE USER linkai_admin WITH LOGIN SUPERUSER INHERIT ENCRYPTED PASSWORD '" + secret00001_admin + "'"); err != nil {
		return err
	}

	// ugh, unavoidable, CREATE USER can not be executed with a prepared statement.
	if _, err := tx.Exec("CREATE USER linkai_jobservice WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + secret00001_jobservice + "'"); err != nil {
		return err
	}

	if _, err := tx.Exec("CREATE USER linkai_inputservice WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + secret00001_inputservice + "'"); err != nil {
		return err
	}

	_, err := tx.Exec("CREATE USER linkai_frontend WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + secret00001_frontend + "'")

	return err
}

func Down00001(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	if _, err := tx.Exec("DROP USER linkai_frontend"); err != nil {
		return err
	}

	if _, err := tx.Exec("DROP USER linkai_inputservice"); err != nil {
		return err
	}

	if _, err := tx.Exec("DROP USER linkai_jobservice"); err != nil {
		return err
	}

	_, err := tx.Exec("DROP USER linkai_admin")
	return err
}
