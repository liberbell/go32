package main

import "log"

func main() {
	// var isTrue bool
	// isTrue = false

	// if isTrue == true {
	// 	log.Println("isTrue is: ", isTrue)
	// } else {
	// 	log.Println("isTrue is: ", isTrue)
	// }

	// cat := "dog"
	// if cat == "cat" {
	// 	log.Println("cat is cat")
	// } else {
	// 	log.Println("cat is not cat")
	// }
	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("My num is greater than 99 and isTrue is set to false")
	} else if myNum < 100 && isTrue {
		log.Println(("My num is less than 100 and isTrue is set to true"))
	}

}
