package main

import (
	"log"
	"time"
)

var s = "seven"
var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// var s2 = "six"

	// s := "eight"

	// log.Println("s is:", s)
	// log.Println("s2 is:", s2)

	// saySomething("xxx")
	user := User{
		FirstName: "Eric",
		LastName:  "Clapton",
	}
	log.Println(user.FirstName)
}

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from saySomething func is:", s)
// 	return s3, "world"
// }
