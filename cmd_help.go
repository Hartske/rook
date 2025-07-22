package main

import "fmt"

func commandHelp(ctx *GameContext, args ...string) error {
	fmt.Println()
	fmt.Println("This is a list of commands that can be used:")
	fmt.Println()
	for _, cmd := range ctx.Commands[ctx.State] {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
