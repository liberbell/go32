package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"

	log.Println(s)
	log.Println(s2)
}

func saySomething(s string) (string, string) {
	return s, "world"
}
