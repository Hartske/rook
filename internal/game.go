package internal

import "fmt"

func StartGame() {
	// Build deck
	deck := BuildDeck()

	// Shuffle
	deck.shuffleDeck()

	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println(deck.Cards[i].Name)
	}
}
