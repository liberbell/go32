package main

import "log"

func main() {
	// var myString string
	// var myInt int

	// myString = "Hi"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]int)
	// myMap["dog"] = "Samuel"
	// myMap["other-dog"] = "Cassie"
	// myMap["dog"] = "fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])
	myMap["First"] = 1
	myMap["Second"] = 2
	log.Println(myMap["First"], myMap["Second"])
}
