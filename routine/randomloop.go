package routine

import (
	"github.com/lsonntag/fleshhack/delay"
	"github.com/lsonntag/fleshhack/loops"
	"github.com/lsonntag/fleshhack/region"
	"github.com/lsonntag/fleshhack/speed"
)

type RandomLoop struct {
	loops     int
	innerLoop *SimpleLoop
}

func NewRandomLoop() *RandomLoop {
	loop := new(RandomLoop)
	loop.loops = -1
	loop.innerLoop = NewSimpleLoopN(region.Random(), speed.Random(), delay.Random(), loops.Random())

	return loop
}

func NewRandomLoopN(l int) *RandomLoop {
	loop := NewRandomLoop()
	loop.loops = l

	return loop
}

func (rl *RandomLoop) Next() (int, int, int) {
	if rl.innerLoop.Done() {
		rl.innerLoop = NewSimpleLoopN(region.Random(), speed.Random(), delay.Random(), loops.Random())
		if rl.loops > 0 {
			rl.loops--
		}
	}

	return rl.innerLoop.Next()
}

func (rl *RandomLoop) Done() bool {
	if rl.loops != 0 {
		return false
	}

	return true
}
