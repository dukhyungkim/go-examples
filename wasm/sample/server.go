package main

import (
	"log"
	"net/http"
)

const (
	addr         = ":8080"
	templatesDir = "."
)

func main() {
	fileServer := http.FileServer(http.Dir(templatesDir))
	if err := http.ListenAndServe(addr, fileServer); err != nil {
		log.Panicln(err)
	}
}
