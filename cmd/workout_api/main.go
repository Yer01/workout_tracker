package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading environmental variables: %v", err)
	}

	dburl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dburl)

	if err != nil {
		log.Fatalf("Error in setting up database: %v", err)
	}

	defer db.Close()

}
