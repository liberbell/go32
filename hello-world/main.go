package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Web server on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
