package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is the home page.")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	n, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValuse(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func AddValuse(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}
