package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hartske/rook/internal"
)

func (ctx *GameContext) placeBids() {
	players := []*internal.Player{
		ctx.PlayerOne,
		ctx.PlayerTwo,
		ctx.PlayerThree,
		ctx.PlayerFour,
	}
	startIndex := -1
	for i, player := range players {
		if player.IsDealer {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		startIndex = 0
	}

	currentIndex := 0
	fmt.Println("begin bid loop:")
	for bids := 0; bids < 4; bids++ {
		fmt.Printf("bid loop: %d\n", (bids + 1))
		actualIndex := (startIndex + currentIndex) % len(players)
		player := players[actualIndex]

		ctx.getBid(player)

		currentIndex++
	}

}

func (ctx *GameContext) getBid(player *internal.Player) {
	switch player.Name {
	case "Player One":
		fmt.Println("call playerBid")
		ctx.playerBid()
	case "Player Two", "Player Three", "Player Four":
		fmt.Println("call comBid")
		comBid(player)
	}
}

func (ctx *GameContext) playerBid() {
	fmt.Println("playerBid started")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Print("Place your bid: ")

	for {
		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "help" {
			fmt.Println("Help message")
		} else if num, err := strconv.Atoi(input); err == nil {
			ctx.PlayerOne.Bid = num
			fmt.Printf("%s bid: %d\n", ctx.PlayerOne.Name, ctx.PlayerOne.Bid)
			return
		} else {
			fmt.Println("Please enter a valid integer")
		}
	}
}

func comBid(player *internal.Player) {
	player.Bid = 5
	fmt.Printf("%s bid: %d\n", player.Name, player.Bid)
}
