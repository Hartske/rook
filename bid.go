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
	bidders := ctx.Play.Bidders

	startIndex := -1
	for i, player := range bidders {
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
		actualIndex := (startIndex + currentIndex) % len(bidders)
		player := bidders[actualIndex]

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
		ctx.comBid(player)
	}
}

func (ctx *GameContext) playerBid() {
	fmt.Println("playerBid started")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Print("Place your bid: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "help" {
			fmt.Println("Help message")
		} else if input == "pass" {
			fmt.Printf("%s passed\n", ctx.PlayerOne.Name)
			return
		} else if num, err := strconv.Atoi(input); err == nil && num%5 == 0 && num > ctx.Play.HighBid {
			ctx.Play.HighBid = num
			ctx.Play.HighBidder = ctx.PlayerOne.Name
			fmt.Printf("%s bid: %d\n", ctx.Play.HighBidder, ctx.Play.HighBid)
			return
		} else {
			fmt.Printf("Please enter a valid integer greater than the current bid: %d\n", ctx.Play.HighBid)
		}
	}
}

func (ctx *GameContext) comBid(player *internal.Player) {
	ctx.Play.HighBid = 10
	ctx.Play.HighBidder = player.Name
	fmt.Printf("%s bid: %d\n", ctx.Play.HighBidder, ctx.Play.HighBid)
}
