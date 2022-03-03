package main

import (
	"context"
	"fmt"
	"go-examples/common/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
}

const (
	defaultTimeout = 5 * time.Second

	humanCollection = "humans"
)

func NewMongo(cfg *config.MongoDB) (*Mongo, error) {
	const URI = "mongodb://%s:%d"
	connUri := fmt.Sprintf(URI, cfg.Host, cfg.Port)

	credential := options.Credential{
		AuthSource: cfg.Database,
		Username:   cfg.Username,
		Password:   cfg.Password,
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

func (m *Mongo) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := m.client.Disconnect(ctx); err != nil {
		log.Println("failed to close db cleanly: ", err)
	}
}

func (m *Mongo) SaveHuman(human *Human) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if _, err := m.database.Collection(humanCollection).InsertOne(ctx, human); err != nil {
		return err
	}
	return nil
}

func (m *Mongo) FindHumanByName(name string) (*Human, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	filter := bson.M{"name": name}

	var human Human
	if err := m.database.Collection(humanCollection).FindOne(ctx, filter).Decode(&human); err != nil {
		return nil, err
	}
	return &human, nil
}

func (m *Mongo) DeleteHumanByName(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	filter := bson.M{"name": name}

	if _, err := m.database.Collection(humanCollection).DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
