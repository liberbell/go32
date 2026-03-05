package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println("Loop count is: ", i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// animals := make(map[string]string)
	// animals["dog"] = "fido"
	// animals["cat"] = "fluffy"
	// var firstLine = "Once upon a midnight breary"
	// firstLine = "x"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstLName string
		LastName   string
		Email      string
		Age        int
	}

	var users []User
	users = append(users, User{"Elton", "John", "elton@example.com", 78})
	users = append(users, User{"Bob", "Maryn", "bob@example.com", 54})
	users = append(users, User{"Eric", "Clapton", "eric@example.com", 74})
	users = append(users, User{"Aretha", "Franclin", "arethaexample.com", 81})

	for _, l := range users {
		fmt.Println(l.FirstLName, l.LastName, l.Email, l.Age)
	}
}
