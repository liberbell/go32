package helpers

import (
	"net/http"

	"github.com/liber/bookings/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w)
}

func ServerError(w http.ResponseWriter, err error) {

}
