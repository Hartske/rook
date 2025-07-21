package internal

import "fmt"

func StartGame() {
	// Build deck
	deck := buildDeck()

	// Shuffle
	shuffleDeck(deck.Cards)

	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println(deck.Cards[i].Name)
	}
}
