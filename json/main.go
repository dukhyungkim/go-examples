package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	filename, err := filepath.Abs("json/test.json")
	if err != nil {
		panic(err)
	}
	log.Printf("filename: %s\n", filename)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var sample Sample

	if err := json.Unmarshal(yamlFile, &sample); err != nil {
		panic(err)
	}

	log.Printf("sample: %+v\n", sample)
}
