package main

import (
	"log"
	"talkhouse/server"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.Create()

	server.Listen(app)

}
