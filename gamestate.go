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
	State    GameState
	Deck     *internal.Deck
	Commands map[GameState]map[string]cliCommand
	Scanner  *bufio.Scanner
	Running  bool
}

func NewGameContext() *GameContext {
	ctx := &GameContext{
		State:    MainMenu,
		Scanner:  bufio.NewScanner(os.Stdin),
		Running:  true,
		Commands: make(map[GameState]map[string]cliCommand),
		Deck:     internal.BuildDeck(),
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

func (ctx *GameContext) setCommands() {
	// Main Menu commands
	ctx.Commands[MainMenu] = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"start": {
			name:        "start",
			description: "Starts a game of Rook",
			callback:    commandStart,
		},
		"exit": {
			name:        "exit",
			description: "Exits the game",
			callback:    commandExit,
		},
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
