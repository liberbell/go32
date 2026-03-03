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
	// myNum := 100
	// isTrue := true

	// if myNum > 99 && !isTrue {
	// 	log.Println("My num is greater than 99 and isTrue is set to false")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("My num is less than 100 and isTrue is set to true")
	// } else if myNum == 101 || isTrue {
	// 	log.Println("Mynum is set to 101 or isTrue is set to true")
	// } else if myNum > 1000 && isTrue == false {
	// 	log.Println("Mynum is greater than 1000 and isTrue is set to false")
	// }

	myVar := "elephant"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "fish":
		log.Println("myVar is set to fish")
	default:
		log.Println("myVar is set to undefine")
	}

}
