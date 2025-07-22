package main

import "github.com/hartske/rook/internal"

func commandStart(ctx *GameContext, args ...string) error {
	internal.StartGame()
	return nil
}
