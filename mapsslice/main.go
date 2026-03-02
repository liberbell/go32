package main

import "log"

// type User struct {
// 	FirstName string
// 	LastName  string
// }

func main() {
	// var myString string
	// var myInt int

	// myString = "Hi"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println(myString, mySecondString, myInt)

	// myMap := make(map[string]User)
	// myMap["dog"] = "Samuel"
	// myMap["other-dog"] = "Cassie"
	// myMap["dog"] = "fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])
	// myMap["First"] = 1
	// myMap["Second"] = 2
	// log.Println(myMap["First"], myMap["Second"])
	// me := User{
	// 	FirstName: "Elton",
	// 	LastName:  "John",
	// }
	// myMap["me"] = me
	// log.Println(myMap["me"])

	// var myNewVar float32
	// myNewVar = 11.1

	// var mySlice []int

	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)

	// log.Println(mySlice)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])
}
