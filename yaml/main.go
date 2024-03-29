package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	filename, err := filepath.Abs("yaml/test.yaml")
	if err != nil {
		panic(err)
	}
	log.Printf("filename: %s\n", filename)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var service Service

	if err := yaml.Unmarshal(yamlFile, &service); err != nil {
		panic(err)
	}

	log.Printf("service: %+v\n", service)
}
