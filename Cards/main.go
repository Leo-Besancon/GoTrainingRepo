package main

// var card string

func main() {

	// card := "Ace of Spades"

	// card = "Five of Diamonds"

	// cards := newDeck()

	// cards, hand := deal(cards, 3)

	// cards.print()
	// hand.print()

	// cards.saveDeckToFile("card")

	cards := newDeck()

	cards.shuffle()

	cards.print()

}
