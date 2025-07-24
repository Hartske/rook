package internal

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Cards []*Card
}

func BuildDeck() *Deck {
	var newCards []*Card
	for i := 0; i < 58; i++ {
		if i < 14 {
			card := &Card{
				Name:  fmt.Sprintf("|Red %d|", i+1),
				Value: i + 1,
				Suit:  "red",
			}
			newCards = append(newCards, card)
		} else if i > 13 && i < 28 {
			card := &Card{
				Name:  fmt.Sprintf("|Green %d|", i-13),
				Value: i - 13,
				Suit:  "green",
			}
			newCards = append(newCards, card)
		} else if i > 28 && i < 43 {
			card := &Card{
				Name:  fmt.Sprintf("|Yellow %d|", i-28),
				Value: i - 28,
				Suit:  "yellow",
			}
			newCards = append(newCards, card)
		} else if i > 43 && i <= 58 {
			card := &Card{
				Name:  fmt.Sprintf("|Black %d|", i-43),
				Value: i - 43,
				Suit:  "black",
			}
			newCards = append(newCards, card)
		}
	}
	rook := &Card{
		Name:  "|Rook|",
		Value: 0,
		Suit:  "black",
	}
	newCards = append(newCards, rook)
	return &Deck{
		Cards: newCards,
	}
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) Draw(player *Player) *Card {
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]

	card.Owner = player.Name

	return card
}

func (d *Deck) Reset() {
	*d = *BuildDeck()
}
