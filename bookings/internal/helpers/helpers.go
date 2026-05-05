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
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {

}
