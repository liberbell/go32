package main

import "log"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println("Loop count is: ", i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	animals := make(map[string]string)
	animals["dog"] = "fido"
	animals["cat"] = "fluffy"

	for anymalType, animal := range animals {
		log.Println(anymalType, animal)
	}
}
