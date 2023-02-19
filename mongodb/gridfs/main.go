package main

import (
	"log"
	"mongodb/config"
)

var (
	p1 = Person{
		FirstName: "Jhon",
		LastName:  "123",
		Age:       10,
	}

	p2 = Person{
		FirstName: "Sam",
		LastName:  "456",
		Age:       20,
	}
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Panicln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Panicln(err)
	}

	db, err := NewMongo(cfg.MongoDB)
	if err != nil {
		log.Panicln(err)
	}

	if err = db.UploadPerson(&p1); err != nil {
		log.Panicln(err)
	}
	if err = db.UploadPerson(&p2); err != nil {
		log.Println(err)
	}

	people, err := db.DownloadPeople()
	if err != nil {
		log.Panicln(err)
	}

	for _, person := range people {
		log.Println(person)
	}
}
