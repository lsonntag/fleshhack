package routine

import (
	"github.com/lsonntag/fleshhack/delay"
	"github.com/lsonntag/fleshhack/loops"
	"github.com/lsonntag/fleshhack/region"
	"github.com/lsonntag/fleshhack/speed"
	"github.com/lsonntag/fleshhack/stroke"
)

type RandomStrokes struct {
	loops     int
	innerLoop *SimpleLoop
	generator *stroke.StrokeGenerator
}

func NewRandomStrokes() *RandomStrokes {
	loop := new(RandomStrokes)
	loop.loops = -1
	loop.generator = &stroke.StrokeGenerator{}

	loop.generator.MinSpeed(speed.Slow).
		MaxSpeed(speed.Fast).
		MinDelay(delay.Short).
		MaxDelay(delay.Long).
		AddRegion(region.Bottom).
		AddRegion(region.Top)

	stroke := loop.generator.Generate()
	loop.innerLoop = NewSimpleLoopN(stroke.Region, stroke.Speed, stroke.Delay, loops.Random())

	return loop
}

func NewRandomStrokesN(l int) *RandomStrokes {
	loop := NewRandomStrokes()
	loop.loops = l

	return loop
}

func (rl *RandomStrokes) Next() (int, int, int) {
	if rl.innerLoop.Done() {
		stroke := rl.generator.Generate()
		rl.innerLoop = NewSimpleLoopN(stroke.Region, stroke.Speed, stroke.Delay, loops.Random())
		if rl.loops > 0 {
			rl.loops--
		}
	}

	return rl.innerLoop.Next()
}

func (rl *RandomStrokes) Done() bool {
	if rl.loops != 0 {
		return false
	}

	return true
}
