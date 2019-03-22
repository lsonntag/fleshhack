package stroke

import (
	"math/rand"

	"github.com/lsonntag/fleshhack/delay"
	"github.com/lsonntag/fleshhack/region"
	"github.com/lsonntag/fleshhack/speed"
)

type SimpleStroke struct {
	Speed  int
	Delay  int
	Region region.Region
}

var SimpleStrokes = map[string]SimpleStroke{
	"VerySlowFull": SimpleStroke{
		speed.VerySlow,
		delay.VeryLong,
		region.Full,
	},

	"SlowFull": SimpleStroke{
		speed.Slow,
		delay.Long,
		region.Full,
	},

	"FastFull": SimpleStroke{
		speed.Med,
		delay.Long,
		region.Full,
	},

	"TeaseFull": SimpleStroke{
		speed.Slow,
		delay.VeryLong * 2,
		region.Full,
	},

	"TeaseTop": SimpleStroke{
		speed.VerySlow,
		delay.VeryLong,
		region.Top,
	},

	"FastTop": SimpleStroke{
		speed.Med,
		delay.Med,
		region.Top,
	},

	"TeaseMid": SimpleStroke{
		speed.VerySlow,
		delay.VeryLong,
		region.Mid,
	},

	"FastMid": SimpleStroke{
		speed.Med,
		delay.Med,
		region.Mid,
	},

	"TeaseLow": SimpleStroke{
		speed.VerySlow,
		delay.VeryLong,
		region.Bottom,
	},

	"FastLow": SimpleStroke{
		speed.Med,
		delay.Med,
		region.Bottom,
	},
}

func RandomSimpleStroke() SimpleStroke {
	i := rand.Intn(len(SimpleStrokes))
	for _, k := range SimpleStrokes {
		if i == 0 {
			return k
		}
		i--
	}
	panic("no!")
	return SimpleStrokes["FastLow"]
}
