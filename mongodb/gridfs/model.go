package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FsFiles struct {
	ID         primitive.ObjectID `bson:"_id"`
	ChunkSize  int32              `bson:"chunkSize"`
	FileName   string             `bson:"fileName"`
	Length     int64              `bson:"length"`
	UpdateDate time.Time          `bson:"updateDate"`
}

type Person struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Age       int    `bson:"age"`
}
