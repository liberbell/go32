package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liber/myniceprogram/pkg/config"
	"github.com/liber/myniceprogram/pkg/handlers"
	"github.com/liber/myniceprogram/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	// fmt.Println(fmt.Sprintf("Starting Web server on port %s", portNumber))
	fmt.Printf("Starting Web server on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
