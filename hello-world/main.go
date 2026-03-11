package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	var x float32
// 	var y float32
// 	x = 100.0
// 	y = 0.0
// 	f, err := DivideValuse(x, y)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by zero")
// 		return
// 	}
// 	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %.02f ", x, y, f))
// }

// func AddValuse(x, y int) int {
// 	var sum int
// 	sum = x + y
// 	return sum
// }

// func DivideValuse(x, y float32) (float32, error) {
// 	var result float32
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result = x / y
// 	return result, nil
// }

func main() {
	// fmt.Println("Hello-world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Web server on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
