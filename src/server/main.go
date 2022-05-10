package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sheodox/seating-chart/migrate"
	"github.com/sheodox/seating-chart/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	err = migrate.RunMigrations()
	if err != nil {
		log.Fatal("Error running migrations ", err)
	}

	if err != nil {
		log.Fatal("Error connecting to database ", err)
	}

	s := server.NewServer()
	s.Start()
}
