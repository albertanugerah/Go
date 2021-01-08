package main

func main() {
	cards := newDeck()

	// inisialisasi nilai (deck,deck)
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}

// func [newfunction]() [type_data]
