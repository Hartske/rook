package main

import (
	"fmt"
	"os"
)

func commandExit(args ...string) error {
	fmt.Println("Closing Rook... GoodBye!")
	os.Exit(0)
	return nil
}
