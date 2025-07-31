package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hartske/rook/internal"
)

type GameState int

const (
	MainMenu GameState = iota
	PreGame
	Dealer
	Bet
	InGame
)

type cliCommand struct {
	name        string
	description string
	callback    func(*GameContext, ...string) error
}

type GameContext struct {
	State       GameState
	Deck        *internal.Deck
	Commands    map[GameState]map[string]cliCommand
	Scanner     *bufio.Scanner
	Running     bool
	PlayerOne   *internal.Player
	PlayerTwo   *internal.Player
	PlayerThree *internal.Player
	PlayerFour  *internal.Player
	Play        *internal.Play
}

func NewGameContext() *GameContext {
	ctx := &GameContext{
		State:    MainMenu,
		Scanner:  bufio.NewScanner(os.Stdin),
		Running:  true,
		Commands: make(map[GameState]map[string]cliCommand),
		Deck:     internal.BuildDeck(),
		PlayerOne: &internal.Player{
			Name:     "Player One",
			Score:    0,
			Bid:      0,
			Hand:     make([]*internal.Card, 0),
			IsDealer: false,
		},
		PlayerTwo: &internal.Player{
			Name:     "Player Two",
			Score:    0,
			Bid:      0,
			Hand:     make([]*internal.Card, 0),
			IsDealer: false,
		},
		PlayerThree: &internal.Player{
			Name:     "Player Three",
			Score:    0,
			Bid:      0,
			Hand:     make([]*internal.Card, 0),
			IsDealer: false,
		},
		PlayerFour: &internal.Player{
			Name:     "Player Four",
			Score:    0,
			Bid:      0,
			Hand:     make([]*internal.Card, 0),
			IsDealer: false,
		},
		Play: &internal.Play{
			HighBid:    0,
			HighBidder: "",
			Trump:      "",
			Pot:        make([]*internal.Card, 0),
		},
	}
	ctx.setCommands()
	return ctx
}

func (ctx *GameContext) changeState(newState GameState) {
	ctx.State = newState
}

func (ctx *GameContext) stop() {
	ctx.Running = false
}

func (ctx *GameContext) gameReset() {
	ctx.changeState(MainMenu)
	ctx.Deck.Reset()
	ctx.playerReset()
	ctx.playReset()
	fmt.Println("         Welcome to Rook CLI!")
	fmt.Println("       ========================")
	fmt.Println("'Start' a game or 'help' to see commands")
	fmt.Println()
}

func (ctx *GameContext) playReset() {
	ctx.Play.HighBid = 0
	ctx.Play.HighBidder = ""
	ctx.Play.Trump = ""
	ctx.Play.ResetPot()
}

func (ctx *GameContext) playerReset() {
	ctx.PlayerOne.Score = 0
	ctx.PlayerOne.Bid = 0
	ctx.PlayerOne.Hand = make([]*internal.Card, 0)
	ctx.PlayerOne.IsDealer = false
	ctx.PlayerTwo.Score = 0
	ctx.PlayerTwo.Bid = 0
	ctx.PlayerTwo.Hand = make([]*internal.Card, 0)
	ctx.PlayerTwo.IsDealer = false
	ctx.PlayerThree.Score = 0
	ctx.PlayerThree.Bid = 0
	ctx.PlayerThree.Hand = make([]*internal.Card, 0)
	ctx.PlayerThree.IsDealer = false
	ctx.PlayerFour.Score = 0
	ctx.PlayerFour.Bid = 0
	ctx.PlayerFour.Hand = make([]*internal.Card, 0)
	ctx.PlayerFour.IsDealer = false
}

func (ctx *GameContext) checkDealer() (*internal.Player, bool) {
	if ctx.PlayerOne.IsDealer {
		return ctx.PlayerOne, true
	}
	if ctx.PlayerTwo.IsDealer {
		return ctx.PlayerTwo, true
	}
	if ctx.PlayerThree.IsDealer {
		return ctx.PlayerThree, true
	}
	if ctx.PlayerFour.IsDealer {
		return ctx.PlayerFour, true
	} else {
		return &internal.Player{}, false
	}
}

func (ctx *GameContext) REPL() {
	for ctx.Running {
		fmt.Print("Rook > ")

		if !ctx.Scanner.Scan() {
			break
		}

		input := ctx.Scanner.Text()
		runCommand(ctx, input)
	}
}
