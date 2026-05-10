package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "host=localhost port=5432 user=bookings_ope password=pass1234 dbname=test_connect sslmode=disable"
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

	err = getAllRows()
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var id int
	var firstName, lastName string

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("User: %d %s %s\n", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	return nil
}
