package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Rules interface{}

func main() {
	rand.Seed(time.Now().UnixNano())
	if g, err := NewGame(1); err != nil {
		for {
			g.turn()
		}
	} else {
		fmt.Println(err)
	}
}
