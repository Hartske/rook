package main

import (
	"fmt"

	"github.com/hartske/rook/internal"
)

func (ctx *GameContext) preGame() {
	ctx.changeState(PreGame)
	fmt.Println("To begin we must decide the first Dealer.")
	fmt.Println("We'll first give the deck a shuffle...")
	ctx.Deck.Shuffle()
	fmt.Println("'Draw' a card!")
}

func (ctx *GameContext) decideDealer() {
	card, exists := ctx.Deck.Draw(ctx.PlayerOne)
	if !exists {
		return
	}
	fmt.Printf("You drew: %s\n", card.Name)
	ctx.Pot = append(ctx.Pot, card)

	card, exists = ctx.Deck.Draw(ctx.PlayerTwo)
	if !exists {
		return
	}
	fmt.Printf("Player 2 drew: %s\n", card.Name)
	ctx.Pot = append(ctx.Pot, card)

	card, exists = ctx.Deck.Draw(ctx.PlayerThree)
	if !exists {
		return
	}
	fmt.Printf("Player 3 drew: %s\n", card.Name)
	ctx.Pot = append(ctx.Pot, card)

	card, exists = ctx.Deck.Draw(ctx.PlayerFour)
	if !exists {
		return
	}
	fmt.Printf("Player 4 drew: %s\n", card.Name)
	ctx.Pot = append(ctx.Pot, card)

	fmt.Println()

	winner := ctx.checkWin(true)

	fmt.Printf("The starting dealer is: %s\n", winner.Name)

	if !winner.IsDealer {
		winner.IsDealer = true
	}

}

func (ctx *GameContext) checkWin(first bool) internal.Player {
	winner := ctx.Pot[0]
	if first {
		for card := range ctx.Pot {
			if ctx.Pot[card] == winner {
				continue
			}

			if ctx.Pot[card].Value > winner.Value {
				winner = ctx.Pot[card]
				continue
			}
		}
	} else {
		leadSuit := ctx.Pot[0].Suit
		for card := range ctx.Pot {
			if ctx.Pot[card] == winner {
				continue
			}

			if ctx.Pot[card].Suit == leadSuit {
				if ctx.Pot[card].Value > winner.Value {
					winner = ctx.Pot[card]
					continue
				} else {
					continue
				}
			}

			if leadSuit != "black" && ctx.Pot[card].Suit == "black" {
				winner = ctx.Pot[card]
				continue
			}
		}
	}
	switch winner.Owner {
	case "Player One":
		return *ctx.PlayerOne
	case "Player Two":
		return *ctx.PlayerTwo
	case "Player Three":
		return *ctx.PlayerThree
	case "Player Four":
		return *ctx.PlayerFour
	default:
		return internal.Player{}
	}
}
