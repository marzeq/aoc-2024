package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marzeq/aoc-2024/day1"
)

func main() {
	if len(os.Args) < 3 {
		panic(fmt.Sprintf("Usage: %s [day] [part]", os.Args[0]))
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Invalid day arg")
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Invalid part arg")
	}

	switch day {
	case 1:
		day1.Run(part)
	}
}
