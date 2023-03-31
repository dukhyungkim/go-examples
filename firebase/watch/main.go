package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const collection = "humans"

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

	snapshotIterator := client.Collection(collection).Snapshots(ctx)
	for {
		var snapshot *firestore.QuerySnapshot
		snapshot, err = snapshotIterator.Next()
		if status.Code(err) == codes.DeadlineExceeded {
			break
		}
		if err != nil {
			log.Panic(err)
		}

		if snapshot == nil {
			continue
		}
		for {
			var doc *firestore.DocumentSnapshot
			doc, err = snapshot.Documents.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("Current %#+v\n", doc)
		}
	}
}
