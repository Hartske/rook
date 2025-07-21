package internal

import "fmt"

func StartGame() {
	deck := buildDeck()
	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println(deck.Cards[i].Name)
	}
}
