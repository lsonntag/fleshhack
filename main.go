package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/funjack/golaunch"
	"github.com/lsonntag/fleshhack/config"
	"github.com/lsonntag/fleshhack/jacker"
	"github.com/lsonntag/fleshhack/routine"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	config.Set("debug", false)

	lCtx := context.Background()
	l := golaunch.NewLaunch()
	if !config.GetBool("debug") {
		log.Printf("Connecting...")
		l.HandleDisconnect(func() {
			os.Exit(0)
		})
		ctx, cancel := context.WithTimeout(lCtx, time.Second*30)
		err := l.Connect(ctx)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connected to Launch")
	}

	//r := routine.NewSimpleLoop(region.High, speed.VerySlow, delay.Long)
	r := routine.NewRandomStrokes()
	j := jacker.NewJacker(r, l)

	j.Play()
}
