package main

import (
	"fmt"
)

// Card classes
const (
	OBSERVATORY = iota
	AQUARIUM
	GARDEN
	LIBRARY
)

// Labyrinth card symbols
const (
	SUN = iota
	MOON
	KEY
)

// Dream card types
const (
	NIGHTMARE = iota
)

type Card interface{}

type DoorCard struct {
	class int
}

type DreamCard struct {
	class int
}

type LabyrinthCard struct {
	class  int
	symbol int
}

func (c DoorCard) String() string {
	var class string

	switch c.class {
	case OBSERVATORY:
		class = "Observatory"
	case AQUARIUM:
		class = "Aquarium"
	case GARDEN:
		class = "Garden"
	case LIBRARY:
		class = "Library"
	default:
		panic("NOTREACHED")
	}

	return fmt.Sprintf("%s Door Card", class)
}

func (c DreamCard) String() string {
	var class string

	switch c.class {
	case NIGHTMARE:
		class = "Nightmare"
	default:
		panic("NOTREACHED")
	}

	return fmt.Sprintf("%s Dream Card", class)
}

func (c LabyrinthCard) String() string {
	var class, symbol string

	switch c.class {
	case OBSERVATORY:
		class = "Observatory"
	case AQUARIUM:
		class = "Aquarium"
	case GARDEN:
		class = "Garden"
	case LIBRARY:
		class = "Library"
	default:
		panic("NOTREACHED")
	}

	switch c.symbol {
	case SUN:
		symbol = "Sun"
	case MOON:
		symbol = "Moon"
	case KEY:
		symbol = "Key"
	default:
		panic("NOTREACHED")
	}

	return fmt.Sprintf("%s %s Labyrinth Card", class, symbol)
}
