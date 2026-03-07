package helpers

import "math/rand/v2"

// type SomeType struct {
// 	TypeName   string
// 	TypeNumber int
// }

func RandomNumber(n int) int {
	value := rand.IntN(n)
	return value
}
