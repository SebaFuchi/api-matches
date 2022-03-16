package main

import (
	"log"
	"os"
	"tinder_matches/internal/server"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		_ = godotenv.Load("/go/bin/.env")
	}

	serv, err := server.New(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	serv.Start()
}
