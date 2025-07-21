package main

import "github.com/hartske/rook/internal"

func commandStart(args ...string) error {
	internal.StartGame()
	return nil
}
