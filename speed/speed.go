package speed

import "math/rand"

const (
	VerySlow = 15
	Slow     = 20
	Med      = 30
	Fast     = 50
	VeryFast = 70
)

func Random() int {
	return 15 + rand.Intn(55)
}

func Clamp(s int) int {
	if s < VerySlow {
		s = VerySlow
	}

	if s > VeryFast {
		s = VeryFast
	}

	return s
}
