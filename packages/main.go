package main

import (
	"log"

	"github.com/liber/myniceprogram/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// log.Println("Hello world")

	// var myVar helpers.SomeType
	// myVar.TypeName = "somename"
	// myVar.TypeNumber = 100

	// fmt.Println(myVar.TypeName, myVar.TypeNumber)
	// PrintText("Hi")
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

// func PrintText(s string) {
// 	log.Println(s)
// }
