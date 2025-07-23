package main

import (
	"fmt"
	"strings"
)

func runCommand(ctx *GameContext, input string) {
	words := cleanInput(input)
	if len(words) == 0 {
		return
	}

	commandName := words[0]
	args := []string{}
	if len(words) > 1 {
		args = words[1:]
	}

	cmd, exists := ctx.Commands[ctx.State][commandName]
	if exists {
		err := cmd.callback(ctx, args...)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		fmt.Printf("Unknown command %s. Type 'help' for available commands. \n", commandName)
		return
	}
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
			description: "Exits Rook",
			callback:    commandExit,
		},
		"test": {
			name:        "test",
			description: "Shows test results",
			callback:    showTest,
		},
		"players": {
			name:        "players",
			description: "Shows players",
			callback:    showPlayers,
		},
	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	clean := strings.Fields(input)
	return clean
}
