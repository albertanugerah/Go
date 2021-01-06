package main

import "fmt"

// create a new type of 'deck'
// which is  a slice of strings
type deck []string

func newDeck() deck {
	// inisialisasi kartu kosong
	cards := deck{}

	jenisKartu := []string{"Sekop", "Hati", "Wajik"}
	noKartu := []string{"As", "Dua", "Tiga", "Empat"}

	for _, jenis := range jenisKartu {
		for _, no := range noKartu {
			cards = append(cards, no+" dari "+jenis)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
