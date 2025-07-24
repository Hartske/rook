package main

import (
	"fmt"

	"github.com/hartske/rook/internal"
)

func (ctx *GameContext) preGame() {
	ctx.changeState(PreGame)
	fmt.Println()
	fmt.Println("To begin we must decide the first Dealer.")
	fmt.Println("We'll first give the deck a shuffle...")
	fmt.Println()
	ctx.Deck.Shuffle()
	fmt.Println("'Draw' a card!")
	fmt.Println()
}

func (ctx *GameContext) decideDealer() {
	var winner *internal.Player
	var hasTie bool

	for {
		card := ctx.Deck.Draw(ctx.PlayerOne)
		fmt.Printf("You drew: %s\n", card.Name)
		ctx.Pot = append(ctx.Pot, card)

		card = ctx.Deck.Draw(ctx.PlayerTwo)
		fmt.Printf("Player 2 drew: %s\n", card.Name)
		ctx.Pot = append(ctx.Pot, card)

		card = ctx.Deck.Draw(ctx.PlayerThree)
		fmt.Printf("Player 3 drew: %s\n", card.Name)
		ctx.Pot = append(ctx.Pot, card)

		card = ctx.Deck.Draw(ctx.PlayerFour)
		fmt.Printf("Player 4 drew: %s\n", card.Name)
		ctx.Pot = append(ctx.Pot, card)

		if winner, hasTie = ctx.dealerWin(); !hasTie {
			fmt.Println()
			fmt.Printf("The starting dealer is: %s\n", winner.Name)
			fmt.Println()
			if !winner.IsDealer {
				winner.IsDealer = true
			}
			ctx.potReset()
			ctx.Deck.Reset()
			ctx.Deck.Shuffle()
			break
		}

		fmt.Println()
		fmt.Println("Tie! Let's draw again!")
		fmt.Println()

		ctx.potReset()
	}
	ctx.preDeal(winner)
}

func (ctx *GameContext) preDeal(winner *internal.Player) {
	if winner == ctx.PlayerOne {
		ctx.changeState(Dealer)
		fmt.Println("It's your 'deal'!")
		fmt.Println()
	} else {
		ctx.deal()
	}
}

func (ctx *GameContext) deal() {
	players := []*internal.Player{
		ctx.PlayerOne,
		ctx.PlayerTwo,
		ctx.PlayerThree,
		ctx.PlayerFour,
	}
	startIndex := -1
	for i, player := range players {
		if !player.IsDealer {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		startIndex = 0
	}

	fmt.Println("Your Hand")
	fmt.Println("=========")
	currentIndex := 0
	for len(ctx.Deck.Cards) > 0 {
		actualIndex := (startIndex + currentIndex) % len(players)
		player := players[actualIndex]
		card := ctx.Deck.Draw(player)
		player.Hand = append(player.Hand, card)
		if player == ctx.PlayerOne {
			fmt.Printf("%s", card.Name)
		}
		currentIndex++
	}
	fmt.Println()
}

func (ctx *GameContext) dealerWin() (*internal.Player, bool) {
	winner := ctx.Pot[0]
	tie := false
	for i := range ctx.Pot {
		card := ctx.Pot[i]
		if card == winner {
			continue
		}
		if card.Value > winner.Value {
			if tie {
				tie = false
			}
			winner = card
			continue
		} else if card != winner && card.Value == winner.Value {
			tie = true
			continue
		} else {
			continue
		}
	}
	switch winner.Owner {
	case "Player One":
		return ctx.PlayerOne, tie
	case "Player Two":
		return ctx.PlayerTwo, tie
	case "Player Three":
		return ctx.PlayerThree, tie
	case "Player Four":
		return ctx.PlayerFour, tie
	default:
		return nil, tie
	}
}

func (ctx *GameContext) checkWin() internal.Player {
	winner := ctx.Pot[0]
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
