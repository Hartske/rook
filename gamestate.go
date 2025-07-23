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
}

func NewGameContext() *GameContext {
	ctx := &GameContext{
		State:    MainMenu,
		Scanner:  bufio.NewScanner(os.Stdin),
		Running:  true,
		Commands: make(map[GameState]map[string]cliCommand),
		Deck:     internal.BuildDeck(),
		PlayerOne: &internal.Player{
			Name:  "Player One",
			Score: 0,
			Hand:  make([]internal.Card, 0),
		},
		PlayerTwo: &internal.Player{
			Name:  "Player Two",
			Score: 0,
			Hand:  make([]internal.Card, 0),
		},
		PlayerThree: &internal.Player{
			Name:  "Player Three",
			Score: 0,
			Hand:  make([]internal.Card, 0),
		},
		PlayerFour: &internal.Player{
			Name:  "Player Four",
			Score: 0,
			Hand:  make([]internal.Card, 0),
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
