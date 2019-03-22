package jacker

import (
	"log"
	"time"

	"github.com/funjack/golaunch"
	"github.com/lsonntag/fleshhack/config"
	"github.com/lsonntag/fleshhack/routine"
)

type Jacker struct {
	launch      golaunch.Launch
	RoutineList *routine.RoutineList
	paused      bool
}

func NewJacker(start routine.Routine, l golaunch.Launch) *Jacker {
	j := new(Jacker)
	j.RoutineList = new(routine.RoutineList)
	j.RoutineList.Push(start)
	j.launch = l

	return j
}

func (j *Jacker) Play() {
	for !j.paused {
		routine := j.RoutineList.Pop()
		if routine == nil {
			break
		}

		for !routine.Done() {
			position, speed, delay := routine.Next()

			log.Printf("goto: %d, speed: %d, delay: %d\n", position, speed, delay)
			if !config.GetBool("debug") {
				j.launch.Move(position, speed)
			}

			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}

func (j *Jacker) Pause() {
	j.paused = true
}
