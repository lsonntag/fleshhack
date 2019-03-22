package delay

import "math/rand"

const (
	VeryShort = 150
	Short     = 300
	Med       = 600
	Long      = 1000
	VeryLong  = 1500
)

func Random() int {
	return VeryShort + rand.Intn(VeryLong-VeryShort)
}

func Clamp(d int) int {
	if d < VeryShort {
		d = VeryShort
	}

	if d > VeryLong {
		d = VeryLong
	}

	return d
}
