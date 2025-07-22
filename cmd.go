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

	cmd, exists := getCommands()[commandName]
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

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	clean := strings.Fields(input)
	return clean
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"start": {
			name:        "start",
			description: "Start a game",
			callback:    commandStart,
		},
		"exit": {
			name:        "exit",
			description: "Exit out of Rook",
			callback:    commandExit,
		},
	}
}
