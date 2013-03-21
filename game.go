package main

import (
	"errors"
	"fmt"
)

// Player turns
const (
	PLAYER_ONE = iota
	PLAYER_TWO
)

// Game phases
const (
	PLAY = iota
	DRAW
	SHUFFLE
)

// Game over conditions
const (
	WIN = iota
	LOSS
)

type Game struct {
	deck        Deck
	limbo       Deck
	discardPile Deck

	sharedResources []LabyrinthCard

	hand      [2][]LabyrinthCard
	labyrinth [2][]LabyrinthCard
	doors     [2][]DoorCard

	players []Player

	numPlayers     int
	curPlayerIndex int

	phase int
}

func NewGame(numPlayers int) (*Game, error) {
	if numPlayers != 1 && numPlayers != 2 {
		return nil, errors.New("Num players must be 1 or 2.")
	}

	game := Game{}

	game.deck = NewBasicDeck()
	game.deck.shuffle()

	game.limbo = make(Deck, 0, 10)
	game.discardPile = make(Deck, 0, 76)

	game.hand[PLAYER_ONE] = game.deck.drawHand()
	game.labyrinth[PLAYER_ONE] = MakeLabyrinth()
	game.doors[PLAYER_ONE] = MakeDoors()

	game.sharedResources = make([]LabyrinthCard, 0, 2)

	game.players = make([]Player, numPlayers, numPlayers)
	game.players[0] = Player{&game.sharedResources, &game.hand[PLAYER_ONE],
		&game.labyrinth[PLAYER_ONE], &game.doors[PLAYER_ONE]}

	if numPlayers == 2 {
		game.hand[PLAYER_TWO] = game.deck.drawHand()
		game.labyrinth[PLAYER_TWO] = MakeLabyrinth()
		game.doors[PLAYER_TWO] = MakeDoors()

		game.players[1] = Player{&game.sharedResources, &game.hand[PLAYER_TWO],
			&game.labyrinth[PLAYER_TWO], &game.doors[PLAYER_TWO]}
	}

	game.numPlayers = numPlayers
	game.curPlayerIndex = PLAYER_ONE

	game.phase = PLAY

	return &game, nil
}

func MakeLabyrinth() []LabyrinthCard {
	return make([]LabyrinthCard, 0, 20)
}

func MakeDoors() []DoorCard {
	return make([]DoorCard, 0, 8)
}

func (g *Game) curPlayer() Player {
	return g.players[g.curPlayerIndex]
}

func (g *Game) turn() {
	var choice int

	switch g.phase {
	case PLAY:
		fmt.Println("1 - Discard")
		fmt.Println("2 - Play")
		fmt.Println("3 - Show Hand(s)")
		fmt.Println("4 - Show Labyrinth(s)")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			if err := g.discardPhase(); err != nil {
				fmt.Println(err)
			}
			g.phase = DRAW
		case 2:
			if err := g.playPhase(); err != nil {
				fmt.Println(err)
			}
			g.phase = DRAW
		case 3:
			g.printHands()
		case 4:
			g.printLabyrinths()
		}
	case DRAW:
	case SHUFFLE:
		g.deck.shuffle()
		g.phase = PLAY
	default:
		panic("NOTREACHED")
	}
}

func (g *Game) printHands() {
	for i, player := range g.players {
		fmt.Printf("Player %d Personal Resources\n", i)
		fmt.Println("----------------------------")
		player.printPersonalResources()
	}

	if g.numPlayers == 2 {
		fmt.Println()
		fmt.Println("Shared Resources")
		fmt.Println("----------------")
		for _, card := range g.sharedResources {
			fmt.Println(card)
		}
	}

	fmt.Println()
}

func (g *Game) printLabyrinths() {
	for i, player := range g.players {
		fmt.Printf("Player %d Labyrinth\n", i)
		fmt.Println("-------------------")
		player.printLabyrinth()
	}

	fmt.Println()
}

func (g *Game) discardPhase() error {
	var choice int

	fmt.Printf("Discard [0-5]: ")

	fmt.Scan(&choice)
	if choice < 0 || choice >= 5 {
		return errors.New("Invalid index.")
	}

	card := g.curPlayer().removeCard(choice)
	g.discardPile = append(g.discardPile, card)

	if card.symbol == KEY {
		g.prophecy()
	}

	return nil
}

func (g *Game) prophecy() {
	var i int

	bound := 5
	if len(g.deck) < bound {
		bound = len(g.deck)
	}

	top := g.deck[0:bound]

	fmt.Println("Top 5 Cards - Discard One")
	for i, card := range top {
		fmt.Printf("%d - %s", i, card)
	}

	for {
		fmt.Scan(&i)
		if i >= 0 && i < bound {
			break
		}
	}

	g.deck = append(g.deck[:i], g.deck[i+1:]...)

	fmt.Println("Old order:")
	for i, card := range top[:bound-1] {
		fmt.Printf("%d - %s", i, card)
	}

	if bound == 5 {
		var n1, n2, n3, n4 int
		for {
			fmt.Print("New order (space-separated): ")
			if num, err := fmt.Scanf("%d %d %d %d", &n1, &n2, &n3, &n4); err != nil && num == 4 {
				g.deck[0], g.deck[1], g.deck[2], g.deck[3] =
					g.deck[n1], g.deck[n2], g.deck[n3], g.deck[n4]
			}
		}
	} else if bound == 4 {
		var n1, n2, n3 int
		for {
			fmt.Print("New order (space-separated): ")
			if num, err := fmt.Scanf("%d %d %d", &n1, &n2, &n3); err != nil && num == 3 {
				g.deck[0], g.deck[1], g.deck[2] = g.deck[n1], g.deck[n2], g.deck[n3]
			}
		}
	} else if bound == 3 {
		var n1, n2 int
		for {
			fmt.Print("New order (space-separated): ")
			if num, err := fmt.Scanf("%d %d", &n1, &n2); err != nil && num == 2 {
				g.deck[0], g.deck[1] = g.deck[n1], g.deck[n2]
			}
		}
	}
}

// Attempts to play card at index |i|.
//
// param	i	the index of the card
//
// returns	an error if an invalid move was attempted, nil otherwise
func (g Game) playPhase() error {
	var i int
	curPlayer := g.curPlayer()

	fmt.Printf("Play [0-5]: ")

	fmt.Scan(&i)
	if i < 0 || i >= 5 {
		return errors.New("Invalid index")
	}

	if err := curPlayer.playCardAt(i); err != nil {
		return err
	}

	if class, ok := curPlayer.canDiscoverDoor(); ok {
		if doorCard, ok := g.deck.removeDoor(class); ok {
			fmt.Printf("%s discovered for player %d", doorCard,
				g.curPlayerIndex+1)

			curPlayer.addDoor(doorCard)
			g.deck.shuffle()
		}
	}

	return nil
}
