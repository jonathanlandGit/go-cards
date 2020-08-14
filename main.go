package main

func main() {

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	////////////
	// cards.print()

	// print out single care
	// card := newCard()
	// fmt.Println(card)

	// print mutliple cards -> slice of type deck
	// cards := deck{"Ace of diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()

	// fmt.Println(cards)

	//using normal for loop to print index, etc.
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

//create new card
// func newCard() string {
// 	return "five of diamonds"
// }
