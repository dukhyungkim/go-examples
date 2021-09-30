package main

import (
	"go-examples/common/config"
	"log"
)

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	mc, err := NewMongo(cfg.MongoDB)
	if err != nil {
		log.Fatalln(err)
	}
	defer mc.Close()

}
