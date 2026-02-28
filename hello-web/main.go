package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("mystring set to:", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
}
