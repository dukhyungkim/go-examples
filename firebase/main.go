package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
)

type Human struct {
	Name string
	Age  int
}

const collection = "humans"

var tom = Human{
	Name: "Tom",
	Age:  10,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Panicln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	write(ctx, client)
	readAll(ctx, client)
	remove(ctx, client)
}

func write(ctx context.Context, client *firestore.Client) {
	_, err := client.Collection(collection).Doc(tom.Name).Set(ctx, tom)
	if err != nil {
		log.Panicln(err)
	}
}

func readAll(ctx context.Context, client *firestore.Client) {
	all, err := client.Collection(collection).Documents(ctx).GetAll()
	if err != nil {
		log.Panicln(err)
	}

	var human Human
	for _, snapshot := range all {
		err = snapshot.DataTo(&human)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("%+v\n", human)
	}
}

func remove(ctx context.Context, client *firestore.Client) {
	_, err := client.Collection(collection).Doc(tom.Name).Delete(ctx)
	if err != nil {
		log.Panicln(err)
	}
}
