package main

import (
	"log"
)

func main() {
	// log.Println("Hello world")

	// var myVar helpers.SomeType
	// myVar.TypeName = "somename"
	// myVar.TypeNumber = 100

	// fmt.Println(myVar.TypeName, myVar.TypeNumber)
	PrintText("Hi")
}

func PrintText(s string) {
	log.Println(s)
}
