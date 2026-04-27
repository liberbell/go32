package main

import (
	"testing"

	"github.com/liber/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case condition:

	}
}
