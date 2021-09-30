package main

type Human struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
