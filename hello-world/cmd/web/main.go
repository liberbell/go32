package main

import (
	"fmt"
	"net/http"

	"github.com/liber/myniceprogram/pkg/config"
	"github.com/liber/myniceprogram/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Web server on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
