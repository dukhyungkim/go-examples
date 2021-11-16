package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const originPassword = "my_secret_password"

func main() {
	log.Println("originPassword:", originPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("hashedPassword:", string(hashedPassword))

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(originPassword)); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln("same")
}
