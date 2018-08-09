package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"gopkg.linkai.io/v1/repos/am/pkg/secrets"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	_ "gopkg.linkai.io/v1/repos/database/pg_migrations"
)

const (
	driver = "postgres"
)

var (
	env    = os.Getenv("APP_ENV")
	region = os.Getenv("APP_REGION")
	flags  = flag.NewFlagSet("goose", flag.ExitOnError)
	dir    = flags.String("dir", ".", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 {
		flags.Usage()
		return
	}

	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	if args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	command := args[0]

	if err := goose.SetDialect(driver); err != nil {
		log.Fatalf("unable to set postgres dialect: %s\n", err)
	}

	dbsecrets := secrets.NewDBSecrets(env, region)
	dbstring, err := dbsecrets.DBString("postgres")
	if err != nil {
		log.Fatalf("error getting database string: %s\n", err)
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("error opening db connection: %s\n", err)
	}

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("pgm run: %v", err)
	}
}

func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: pgm [OPTIONS] COMMAND
Examples:
	pgm -dir ./am_migrations status
	pgm -dir ./am_migrations create init sql
	pgm -dir ./am_migrations create something_from_go_file go
	pgm -dir ./am_migrations up
	pgm -dir ./am_migrations down
	pgm -dir ./am_migrations redo
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)