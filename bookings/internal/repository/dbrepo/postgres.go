package dbrepo

import "github.com/liber/bookings/internal/models"

func (m *PostgresDBRepo) AllUsers() bool {
	return true
}

func (m *PostgresDBRepo) InsertReservation(res models.Reservation) error {

	stmt := `insert into reservations
			(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) 
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.Exec(stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.CreatedAt,
		res.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
