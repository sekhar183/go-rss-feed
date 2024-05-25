package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func postgres() {
	// Database connection parameters
	//connStr := "user=user password=user dbname=rss-feed sslmode=disable host=localhost port=5432"

	// Open the database connection
	db, err := sql.Open("postgres", "postgres://user:user@localhost:5432/rss-feed?sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")
}
