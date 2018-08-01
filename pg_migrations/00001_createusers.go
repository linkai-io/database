package migration

import (
	"database/sql"
	"errors"
	"os"

	"github.com/pressly/goose"
)

var users map[string]string

var (
	secret00001_admin = os.Getenv("GOOSE_SECRET_00001_admin")
)

func init() {
	goose.AddMigration(Up00001, Down00001)
	users = make(map[string]string, 0)
	users["jobservice"] = os.Getenv("GOOSE_SECRET_00001_jobservice")
	users["inputservice"] = os.Getenv("GOOSE_SECRET_00001_inputservice")
	users["orgservice"] = os.Getenv("GOOSE_SECRET_00001_orgservice")
	users["userservice"] = os.Getenv("GOOSE_SECRET_00001_userservice")
	users["tagservice"] = os.Getenv("GOOSE_SECRET_00001_tagservice")
	users["scangroupservice"] = os.Getenv("GOOSE_SECRET_00001_scangroupservice")
	users["hostservice"] = os.Getenv("GOOSE_SECRET_00001_hostservice")
}

func Up00001(tx *sql.Tx) error {

	if secret00001_admin == "" {
		return errors.New("GOOSE_SECRET_00001_admin not set")
	}

	if _, err := tx.Exec("CREATE USER linkai_admin WITH LOGIN SUPERUSER INHERIT ENCRYPTED PASSWORD '" + secret00001_admin + "'"); err != nil {
		return err
	}

	// ugh, unavoidable, CREATE USER can not be executed with a prepared statement.
	for service, passwd := range users {
		if _, err := tx.Exec("CREATE USER " + service + " WITH LOGIN NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION ENCRYPTED PASSWORD '" + passwd + "'"); err != nil {
			return err
		}
	}

	return nil
}

func Down00001(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	for service := range users {
		if _, err := tx.Exec("DROP USER " + service); err != nil {
			return err
		}
	}

	_, err := tx.Exec("DROP USER linkai_admin")
	return err
}
