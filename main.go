package main

import "fmt"

func main() {
	fmt.Println("         Welcome to Rook CLI!")
	fmt.Println("       ========================")
	fmt.Println("'Start' a game or 'help' to see commands")
	fmt.Println()
	ctx := NewGameContext()
	ctx.REPL()
}
