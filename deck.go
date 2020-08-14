package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//create new type of deck which is slice of strings
//new deck type borrows this
type deck []string

//deal the deck
func deal(d deck, handSize int) (deck, deck) {
	//expect two vals to be returned
	return d[:handSize], d[handSize:]
}

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

//utility to return a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	//anyone can read/write to this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
