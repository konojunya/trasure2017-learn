package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/konojunya/voyage-group/reports/controller"
	"github.com/suzuken/wiki/db"
)

func main() {
	var (
		r      controller.Report
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)
	flag.Parse()

	db := Init(*dbconf, *env)
	r.DB = db

	r.Save()
}

func Init(dbconfig, env string) *sql.DB {
	cs, err := db.NewConfigsFromFile(dbconfig)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	db, err := cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}

	return db
}
