package main

import "log"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println("Loop count is: ", i)
	// }

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}
}
