package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		fmt.Println(fmt.Sprintf("Number of bytes written: " + n))
	})
}
