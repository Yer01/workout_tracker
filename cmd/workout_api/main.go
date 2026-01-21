package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading environmental variables: %v", err)
	}
}
