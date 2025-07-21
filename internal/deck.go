package internal

import (
	"fmt"
)

type Deck struct {
	Cards []Card
}

func buildDeck() Deck {
	var newCards []Card
	for i := 0; i < 56; i++ {
		if i < 14 {
			card := Card{
				Name:  fmt.Sprintf("Red | %d", i+1),
				Value: i + 1,
				Suit:  "red",
			}
			newCards = append(newCards, card)
		} else if i > 13 && i < 27 {
			card := Card{
				Name:  fmt.Sprintf("Green | %d", i-13),
				Value: i - 13,
				Suit:  "green",
			}
			newCards = append(newCards, card)
		} else if i > 27 && i < 41 {
			card := Card{
				Name:  fmt.Sprintf("Yellow | %d", i-27),
				Value: i - 27,
				Suit:  "yellow",
			}
			newCards = append(newCards, card)
		} else if i > 41 && i <= 55 {
			card := Card{
				Name:  fmt.Sprintf("Black | %d", i-41),
				Value: i - 41,
				Suit:  "black",
			}
			newCards = append(newCards, card)
		}
	}
	rook := Card{
		Name:  "Rook",
		Value: 0,
		Suit:  "black",
	}
	newCards = append(newCards, rook)
	return Deck{
		Cards: newCards,
	}
}
