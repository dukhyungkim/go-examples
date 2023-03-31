package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
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
		for _, change := range snapshot.Changes {
			switch change.Kind {
			case firestore.DocumentAdded:
				log.Printf("new Human: %#+v\n", change.Doc.Data())
			case firestore.DocumentModified:
				log.Printf("update Human: %#+v\n", change.Doc.Data())
			case firestore.DocumentRemoved:
				log.Printf("removed Human: %#+v\n", change.Doc.Data())
			}
		}
	}
}
