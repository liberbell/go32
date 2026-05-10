package main

import (
	"database/sql"
	"log"
)

func main() {
	_, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=booking_ope password=pass1234")
	if err != nil {
		log.Fatalf("Anable to connect: %v", err)
	}
}
