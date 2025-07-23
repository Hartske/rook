package main

import (
	"fmt"
)

func commandQuit(ctx *GameContext, args ...string) error {
	fmt.Println("Quitting current game...")
	ctx.gameReset()
	return nil
}
