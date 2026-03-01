package main

import "log"

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Bob"

	myVar2 := myStruct{
		FirstName: "Elton",
	}
	log.Println("myVar is set to:", myVar)
	log.Println("myVar2 is set to:", myVar2)

}
