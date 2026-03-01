package main

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Bob"

	myVar2 := myStruct{
		FirstName: "Elton",
	}
}
