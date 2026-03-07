package main

import (
	"fmt"
	"log"

	"github.com/liber/myniceprogram/helpers"
)

func main() {
	log.Println("Hello world")

	var myVar helpers.SomeType
	myVar.TypeName = "somename"
	myVar.TypeNumber = 100

	fmt.Println(myVar.TypeName, myVar.TypeNumber)
}
