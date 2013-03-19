package main

type CardSlice []Card

func (c *CardSlice) remove(i int) Card {
	card := (*c)[i]
	*c = append((*c)[:i], (*c)[i+1:]...)
	return card
}
