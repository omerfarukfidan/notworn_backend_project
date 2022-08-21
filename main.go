package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalln("error was:", envErr)
	}
	db, err := ConnectDataBase()
	if err != nil {
		log.Fatalln("error was:", db.Error)
	}

}
