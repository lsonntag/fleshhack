package routine

import (
	"github.com/lsonntag/fleshhack/region"
)

type SimpleLoop struct {
	region region.Region
	speed  int
	delay  int

	flip  bool
	loops int
}

func NewSimpleLoop(r region.Region, s int, d int) *SimpleLoop {
	loop := new(SimpleLoop)
	loop.region = r
	loop.speed = s
	loop.delay = d
	loop.loops = -1

	return loop
}

func NewSimpleLoopN(r region.Region, s int, d int, l int) *SimpleLoop {
	loop := NewSimpleLoop(r, s, d)
	loop.loops = l

	return loop
}

func (sl *SimpleLoop) Next() (int, int, int) {
	pos := sl.region.Low
	if sl.flip {
		pos = sl.region.High
	}

	sl.flip = !sl.flip
	if sl.loops > 0 {
		sl.loops--
	}

	return pos, sl.speed, sl.delay
}

func (sl *SimpleLoop) Done() bool {
	if sl.loops != 0 {
		return false
	}

	return true
}
