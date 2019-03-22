package region

import "math/rand"

type Region struct {
	Low  int
	High int
}

var (
	Top       = Region{65, 95}
	MidTop    = Region{35, 95}
	Mid       = Region{35, 65}
	MidBottom = Region{5, 65}
	Bottom    = Region{5, 35}
	Full      = Region{5, 95}
)

func Random() Region {
	switch rand.Intn(6) {
	case 0:
		return Bottom
	case 1:
		return MidBottom
	case 2:
		return Mid
	case 3:
		return MidTop
	case 4:
		return Top
	case 5:
		return Full
	}

	return Full
}
