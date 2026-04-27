package handlers

import (
	"encoding/gob"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/liber/bookings/internal/config"
	"github.com/liber/bookings/internal/models"
	"github.com/liber/bookings/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes(t *testing.T) http.Handler {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := NewRepo(&app)
	NewHandlers(repo)
	render.NewTemplates(&app)
}
