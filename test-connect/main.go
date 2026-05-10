package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "host=localhost port=5432 user=booking_ope password=pass1234 dbname=test_connect sslmode=disable"
	conn, err := sql.Open("pgx", dsn)
	if err != nil {

		log.Fatalf("Anable to connect: %v", err)
	}
	defer conn.Close()

	log.Println("Connect to database.")

	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database:", err)
	}

	log.Println("Pinged database.")
}
