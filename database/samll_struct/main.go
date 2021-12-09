package main

import (
	"go-examples/common/config"
	"go-examples/database/samll_struct/repository"
	"log"
)

func main() {
	opts, err := config.ParseDBFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("opts: %+v\n", opts)

	var storage *repository.Storage
	if opts.MariaDB {
		db, err := repository.NewMariaDB(cfg.RDB)
		if err != nil {
			log.Fatalln(err)
		}
		storage = db
	} else if opts.Postgres {
		db, err := repository.NewPostgreSQL(cfg.RDB)
		if err != nil {
			log.Fatalln(err)
		}
		storage = db
	} else {
		log.Fatalln("not selected db")
	}

	if opts.Migration {
		if err := storage.Migration(); err != nil {
			log.Fatalln(err)
		}
	}
}
