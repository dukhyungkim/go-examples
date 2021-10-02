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

	human := Human{
		Name: "test",
		Age:  123,
	}
	if err := mc.SaveHuman(&human); err != nil {
		log.Fatalln(err)
	}

	found, err := mc.FindHumanByName(human.Name)
	if err != nil {
		log.Fatalln(err)
	}

	if human.IsEqual(found) == false {
		log.Println("not same")
	}

	if err := mc.DeleteHumanByName(human.Name); err != nil {
		log.Fatalln(err)
	}
}
