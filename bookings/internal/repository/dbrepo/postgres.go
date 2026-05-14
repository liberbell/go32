package dbrepo

import "github.com/liber/bookings/internal/models"

func (m *PostgresDBRepo) AllUsers() bool {
	return true
}

func (m *PostgresDBRepo) InsertReservation(res models.Reservation) error {
	return nil
}
