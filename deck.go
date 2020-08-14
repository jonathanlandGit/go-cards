package main

import "fmt"

//create new type of deck which is slice of strings
//new deck type borrows this
type deck []string

//returns a deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//call deck type and print each card within deck
//loop through deck of cards and print out vals
//receive a deck - any var of type deck gets access to this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
