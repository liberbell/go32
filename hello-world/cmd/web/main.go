package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/liber/myniceprogram/pkg/config"
	"github.com/liber/myniceprogram/pkg/handlers"
	"github.com/liber/myniceprogram/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	session := scs.New()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting Web server on port %s", portNumber)

	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
