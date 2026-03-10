package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is the home page.")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	n, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	_ = http.ListenAndServe(":8080", nil)
}
