package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	newInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Rook > ")
		newInput.Scan()

		words := cleanInput(newInput.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	clean := strings.Fields(input)
	return clean
}

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
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
