package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "Eric",
			"last_name": "Clapton",
			"hair_color": "Gray",
			"has_dog": false
		},
		{
			"first_name": "Bob",
			"last_name": "Bary",
			"hair_color": "Black",
			"has_dog": true
		}
	]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	var mySlice []Person
	var m1 Person
	var m2 Person

	m1.FirstName = "Elton"
	m1.LastName = "John"
	m1.HairColor = "Gold"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	m2.FirstName = "John"
	m2.LastName = "lennon"
	m2.HairColor = "Brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("Error marshalling: ", err)
	}

	fmt.Println(string(newJson))
}
