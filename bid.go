package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getBid() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}
		input := strings.TrimSpace(scanner.Text())

		if num, err := strconv.Atoi(input); err == nil {
			return num
		} else {
			fmt.Println("Please enter a valid integer")
		}
	}
}
