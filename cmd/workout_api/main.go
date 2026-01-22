package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Yer01/workout_tracker/internal/api/application"
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

	app := application.New()

	if err = app.Start(context.TODO()); err != nil {
		log.Fatalf("Failed to start application: %w", err)
	}

	os.Exit(0)

}
