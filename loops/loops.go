package loops

import "math/rand"

const (
	One  = 1
	Few  = 3
	Some = 5
	Many = 7
	Ten  = 10
)

func Random() int {
	return 1 + rand.Intn(9)
}
