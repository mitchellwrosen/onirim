package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Deck []Card

// Fisher-Yates shuffle
func (d *Deck) shuffle() {
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}

	fmt.Println("Deck shuffled")
}

func (d *Deck) draw() (Card, error) {
	numCards := len(*d)

	if numCards == 0 {
		return nil, errors.New("Cannot draw from empty deck")
	}

	card := (*d)[0]

	if numCards == 1 {
		*d = nil
	}

	*d = (*d)[1:]

	return card, nil
}

func (d *Deck) drawHand() []LabyrinthCard {
	var card Card

	hand := make([]LabyrinthCard, 0, 5)
	putBack := make(Deck, 0, 5)

	for len(hand) < 5 {
		if card, err := d.draw(); err != nil {
			panic("Tried to draw a hand from an empty deck")
		}

		switch card.(type) {
		case DoorCard:
			putBack = append(putBack, card)
		case DreamCard:
			putBack = append(putBack, card)
		default:
			hand = append(hand, card.(LabyrinthCard))
		}
	}

	*d = append(*d, putBack)
	d.shuffle()

	return hand
}

// Removes a Door card of class |class|. Shuffles the deck afterward.
// param	class	The class of the Door to remove.
// returns	The removed Door card, and a boolean indicating success (it may be
//			the case that there were no more Door cards of that class)
func (d *Deck) removeDoor(class int) (DoorCard, bool) {
	for i, card := range *d {
		if doorCard, ok := card.(DoorCard); ok && doorCard.class == class {
			*d = append((*d)[:i], (*d)[i+1:]...)
			d.shuffle()
			return doorCard, true
		}
	}

	return DoorCard{}, false
}

// 8 Door (2 of each color)
// 16 Observatory (9 Sun, 4 Moon, 3 Key)
// 15 Aquarium (8 Sun, 4 Moon, 3 Key)
// 14 Garden (7 Sun, 4 Moon, 3 Key)
// 13 Library (6 Sun, 4 Moon, 3 Key)
// 10 Dream (all Nightmare)
func NewBasicDeck() Deck {
	d := make(Deck, 0, 76)

	for i := 0; i < 2; i++ {
		d = append(d, DoorCard{OBSERVATORY})
		d = append(d, DoorCard{AQUARIUM})
		d = append(d, DoorCard{GARDEN})
		d = append(d, DoorCard{LIBRARY})
	}

	for i := 0; i < 6; i++ {
		d = append(d, LabyrinthCard{OBSERVATORY, SUN})
		d = append(d, LabyrinthCard{AQUARIUM, SUN})
		d = append(d, LabyrinthCard{GARDEN, SUN})
		d = append(d, LabyrinthCard{LIBRARY, SUN})
	}
	d = append(d, LabyrinthCard{OBSERVATORY, SUN})
	d = append(d, LabyrinthCard{OBSERVATORY, SUN})
	d = append(d, LabyrinthCard{OBSERVATORY, SUN})
	d = append(d, LabyrinthCard{AQUARIUM, SUN})
	d = append(d, LabyrinthCard{AQUARIUM, SUN})
	d = append(d, LabyrinthCard{GARDEN, SUN})

	for i := 0; i < 4; i++ {
		d = append(d, LabyrinthCard{OBSERVATORY, MOON})
		d = append(d, LabyrinthCard{AQUARIUM, MOON})
		d = append(d, LabyrinthCard{GARDEN, MOON})
		d = append(d, LabyrinthCard{LIBRARY, MOON})
	}

	for i := 0; i < 3; i++ {
		d = append(d, LabyrinthCard{OBSERVATORY, KEY})
		d = append(d, LabyrinthCard{AQUARIUM, KEY})
		d = append(d, LabyrinthCard{GARDEN, KEY})
		d = append(d, LabyrinthCard{LIBRARY, KEY})
	}

	for i := 0; i < 10; i++ {
		d = append(d, DreamCard{NIGHTMARE})
	}

	return d
}
