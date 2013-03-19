package main

import (
	"math/rand"
	"time"
)

type Rules interface{}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := newOnePlayerGame()

	for {
		g.takeTurn()
	}
}
