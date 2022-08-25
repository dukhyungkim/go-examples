package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sqlc/tutorial/sample"

	_ "github.com/lib/pq"
)

const (
	HOST     = "192.168.0.219"
	DATABASE = "dvdrental"
	USER     = "user01"
	PASSWORD = "user01_password!"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)
	defer func() {
		err = db.Close()
		checkError(err)
	}()

	err = db.Ping()
	checkError(err)
	log.Println("Successfully created connection to database")

	ctx := context.Background()

	queries := sample.New(db)

	log.Println("query a Actor")
	actor, err := queries.GetActor(ctx, 1)
	checkError(err)
	log.Println(actor)

	log.Println("query 10 Actors")
	listActorsParam := sample.ListActorsParams{
		Limit:  10,
		Offset: 20,
	}
	actors, err := queries.ListActors(ctx, listActorsParam)
	checkError(err)
	log.Println("len(actors) =", len(actors))
	for i, a := range actors {
		log.Println(i, a)
	}

	log.Println("query all Actors")
	actors, err = queries.ListAllActors(ctx)
	checkError(err)
	for i, a := range actors {
		log.Println(i, a)
	}
}
