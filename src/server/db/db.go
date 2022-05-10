package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Connection *sqlx.DB

func ConnectionString() string {
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgHost := os.Getenv("PGHOST")
	pgDatabase := os.Getenv("PGDATABASE")
	return fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable", pgUser, pgPassword, pgHost, pgDatabase)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	Connection, err = sqlx.Connect("postgres", ConnectionString())

	if err != nil {
		log.Fatal("Error connecting to database!", err)
	}
}
