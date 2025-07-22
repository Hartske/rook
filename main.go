package main

import "fmt"

func main() {
	fmt.Println("Welcome to Rook CLI")
	ctx := NewGameContext()
	ctx.REPL()
}
