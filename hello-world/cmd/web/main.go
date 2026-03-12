package main

import (
	"fmt"
	"net/http"

	"github.com/liber/myniceprogram/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Web server on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
