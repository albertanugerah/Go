package main

func main() {
	// func printState di file state.go
	// printState()

	cards := deck{"Ace of Diamonds", newCard()}
	// menambahkan nilai card
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string { // ketika di eksekusi, function mengembalikan tipe data
	return "Five of Diamonds"
}

// func [newfunction]() [type_data]
