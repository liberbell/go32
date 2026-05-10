package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	_, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=booking_ope password=pass1234")
	if err != nil {
		log.Fatalf("Anable to connect: %v", err)
	}
}
