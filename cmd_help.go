package main

import "fmt"

func commandHelp(args ...string) error {
	fmt.Println()
	fmt.Println("This is a list of commands that can be used:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
