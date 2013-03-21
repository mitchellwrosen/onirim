package main

import (
	"errors"
	"fmt"
)

// A convenience class to encapsulate behavior related with a player. All game
// logic still exists in |game|. Doesn't own any data.
type Player struct {
	sharedResources *[]LabyrinthCard

	hand      *[]LabyrinthCard
	labyrinth *[]LabyrinthCard
	doors     *[]DoorCard
}

func (p Player) printPersonalResources() {
	for _, card := range *p.hand {
		fmt.Println(card)
	}
}

func (p Player) printLabyrinth() {
	for _, card := range *p.labyrinth {
		fmt.Println(card)
	}
}

// Removes the card at index |i| from this player's hand.
//
// requires			0 <= i < 5
//
// param		i	The index of the card to remove.
//
// returns			The removed card.
func (p Player) removeCard(i int) LabyrinthCard {
	card, pile := p.cardAt(i)
	*pile = append(*pile, card)
	return card
}

// Gets this player's card at index |i|. Also returns the pile from which this
// card came from, to make removal easier.
//
// requires			0 <= i < 5
//
// param	i		The index of the card.
//
// returns	card	The card.
//			pile	The address of the pile the card came from (either player's
//						hand or shared)
func (p Player) cardAt(i int) (card LabyrinthCard, pile *[]LabyrinthCard) {
	if i < len(*p.hand) {
		pile = p.hand
		card = (*pile)[i]
	} else {
		pile = p.sharedResources
		card = (*pile)[i-3]
	}

	return card, pile
}

// Plays this player's card at index |i| to his/her Labyrinth.
//
// param	i	The index of the card to play.
//
// returns	an error if |card| could not be played, nil otherwise
func (p Player) playCardAt(i int) error {
	card, _ := p.cardAt(i)

	if end, ok := p.labyrinthEnd(); ok {
		if card.symbol == end.symbol {
			return errors.New("Cannot place adjacent symbols.")
		}
	}

	p.removeCard(i)
	*p.labyrinth = append(*p.labyrinth, card)

	return nil
}

// Gets the end of this Player's Labyrinth.
//
// returns	the LabyrinthCard at the end of this Player's Labyrinth, and a
//			boolean representing whether or not there were any cards in this
//			Player's Labyrinth to get
func (p Player) labyrinthEnd() (LabyrinthCard, bool) {
	if len(*p.labyrinth) == 0 {
		return LabyrinthCard{}, false
	}

	return (*p.labyrinth)[len(*p.labyrinth)-1], true
}

// Gets whether or not the Player can discover a Door (the last three cards
// played were of the same class).
//
// returns	a boolean representing whether or not this Player can discover a
//			Door, and an integer representing the class of this Door, if true.
func (p *Player) canDiscoverDoor() (int, bool) {
	if len(*p.labyrinth) >= 3 {
		class1 := (*p.labyrinth)[len(*p.labyrinth)-1].class
		class2 := (*p.labyrinth)[len(*p.labyrinth)-2].class
		class3 := (*p.labyrinth)[len(*p.labyrinth)-3].class

		if class1 == class2 && class2 == class3 {
			return class1, true
		}
	}

	return -1, false
}

// Adds |door| to this Player's doors.
func (p *Player) addDoor(door DoorCard) {
	*p.doors = append(*p.doors, door)
}

// Determines if the player has a Key of class |class|.
//
// param	class	The class of Key to look for.
//
// returns	The index of the Key, if found, and a bool representing whether or
//			not the key was found.
func (p *Player) hasKey(class int) (int, bool) {
	for i, card := range *p.hand {
		if card.symbol == KEY && card.class == class {
			return i, true
		}
	}

	for i, card := range *p.sharedResources {
		if card.symbol == KEY && card.class == class {
			return i + 3, true
		}
	}

	return -1, false
}
