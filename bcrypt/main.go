package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const originPassword = "my_secret_password"

func main() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(originPassword)); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln("same")
}
