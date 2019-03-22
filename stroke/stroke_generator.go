package stroke

import (
	"log"
	"math/rand"

	"github.com/lsonntag/fleshhack/delay"
	"github.com/lsonntag/fleshhack/region"
	"github.com/lsonntag/fleshhack/speed"
)

type StrokeGenerator struct {
	minSpeed       int
	maxSpeed       int
	minDelay       int
	maxDelay       int
	allowedRegions []region.Region
}

func (sg *StrokeGenerator) MinSpeed(s int) *StrokeGenerator {
	sg.minSpeed = speed.Clamp(s)
	return sg
}

func (sg *StrokeGenerator) MaxSpeed(s int) *StrokeGenerator {
	sg.maxSpeed = speed.Clamp(s)
	return sg
}

func (sg *StrokeGenerator) MinDelay(d int) *StrokeGenerator {
	sg.minDelay = delay.Clamp(d)
	return sg
}

func (sg *StrokeGenerator) MaxDelay(d int) *StrokeGenerator {
	sg.maxDelay = delay.Clamp(d)
	return sg
}

func (sg *StrokeGenerator) AddRegion(r region.Region) *StrokeGenerator {
	for _, v := range sg.allowedRegions {
		if v == r {
			return sg
		}
	}

	sg.allowedRegions = append(sg.allowedRegions, r)

	return sg
}

func (sg *StrokeGenerator) DelRegion(r region.Region) *StrokeGenerator {
	for i, v := range sg.allowedRegions {
		if v == r {
			sg.allowedRegions = append(sg.allowedRegions[:i], sg.allowedRegions[i+1:]...)
		}
	}

	return sg
}

func (sg *StrokeGenerator) Generate() SimpleStroke {
	if sg.maxSpeed == 0 {
		sg.maxSpeed = speed.VeryFast
	}

	if sg.maxDelay == 0 {
		sg.maxDelay = delay.VeryLong
	}

	if sg.maxSpeed < sg.minSpeed {
		panic("maxspeed < minspeed")
	}
	if sg.maxDelay < sg.minDelay {
		panic("maxdelay < mindelay")
	}

	dSpeed := sg.maxSpeed - sg.minSpeed
	speed := sg.minSpeed + rand.Intn(dSpeed)

	dDelay := sg.maxDelay - sg.minDelay
	delay := sg.minDelay + rand.Intn(dDelay)

	var reg region.Region

	if len(sg.allowedRegions) == 0 {
		log.Println("len 0")
		reg = region.Random()
	} else if len(sg.allowedRegions) == 1 {
		log.Println("len 1")
		reg = sg.allowedRegions[0]
	} else {
		log.Println("len >1")
		i := rand.Intn(len(sg.allowedRegions))
		reg = sg.allowedRegions[i]
	}

	return SimpleStroke{
		Speed:  speed,
		Delay:  delay,
		Region: reg,
	}
}
