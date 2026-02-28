package main

import "fmt"

var myName string

func main() {
	fmt.Println("Hello world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world."
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
