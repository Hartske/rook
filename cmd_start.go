package main

import "fmt"

func commandStart(ctx *GameContext, args ...string) error {
	fmt.Println("Starting a new game...")
	ctx.preGame()

	return nil
}
