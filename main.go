package main

import (
	"talkhouse/helper"
	"talkhouse/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	helper.CheckError(err, "Error loading .env file")
}

func main() {
	app := server.Create()
	server.Listen(app)
}
