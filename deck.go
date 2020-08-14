package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		//take whatever is at new positoin and assigne to i
		newPosition := r.Intn(len(d) - 1)
		//now swap els
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
