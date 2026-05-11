package dbrepo

import (
	"database/sql"

	"github.com/liber/bookings/internal/config"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
