package main

import (
	"fmt"
	"os"
)

func commandExit(ctx *GameContext, args ...string) error {
	fmt.Println("Closing Rook... GoodBye!")
	ctx.stop()
	os.Exit(0)
	return nil
}
