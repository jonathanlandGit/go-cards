package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//read from file
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//type conversion
	// greeting := "hey"
	// fmt.Println([]byte(greeting))

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
