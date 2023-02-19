package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mongodb/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
}

const (
	defaultTimeout = 5 * time.Second
	bucketName     = "people"
)

func NewMongo(cfg *config.MongoDB) (*Mongo, error) {
	const URI = "mongodb://%s:%d"
	connUri := fmt.Sprintf(URI, cfg.Host, cfg.Port)

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	connCtx, connCancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer connCancel()

	client, err := mongo.Connect(connCtx, options.Client().ApplyURI(connUri).SetAuth(credential))
	if err != nil {
		return nil, fmt.Errorf("failed to connect db; %w", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer pingCancel()
	if err = client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping db; %w", err)
	}

	return &Mongo{
		client:   client,
		database: client.Database(cfg.Database),
	}, nil
}

func (m *Mongo) UploadPerson(p *Person) error {
	opts := &options.BucketOptions{}
	bucket, err := gridfs.NewBucket(m.database, opts.SetName(bucketName))
	if err != nil {
		return err
	}

	uploadStream, err := bucket.OpenUploadStream(p.LastName)
	if err != nil {
		return err
	}
	defer func() {
		if err = uploadStream.Close(); err != nil {
			log.Println(err)
		}
	}()

	b, err := bson.Marshal(p)
	if err != nil {
		return err
	}
	size, err := uploadStream.Write(b)
	if err != nil {
		return err
	}

	log.Printf("Write file to DB was successful. size: %d\n", size)
	return nil
}

func (m *Mongo) DownloadPeople() ([]*Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	people := m.database.Collection(bucketName)
	cursor, err := people.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Println(err)
		}
	}()

	var fsFiles []*FsFiles
	if err = cursor.All(ctx, &fsFiles); err != nil {
		return nil, err
	}

	results := make([]*Person, len(fsFiles))

	bucket, err := gridfs.NewBucket(m.database)
	if err != nil {
		return nil, err
	}
	for i, result := range fsFiles {
		log.Println(result)

		var buf bytes.Buffer
		size, err := bucket.DownloadToStream(result.ID, &buf)
		if err != nil {
			return nil, err
		}
		fmt.Printf("size to download: %v\n", size)

		p := Person{}
		if err = bson.Unmarshal(buf.Bytes(), &p); err != nil {
			return nil, err
		}
		results[i] = &p
	}

	return results, nil
}
