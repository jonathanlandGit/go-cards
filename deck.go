package main

import "fmt"

//create new type of deck which is slice of strings
//new deck type borrows this
type deck []string

//call deck type and print each card within deck
//loop through deck of cards and print out vals
//receive a deck - any var of type deck gets access to this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
