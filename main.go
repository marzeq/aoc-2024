package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marzeq/aoc-2024/day1"
	"github.com/marzeq/aoc-2024/day2"
	"github.com/marzeq/aoc-2024/day3"
	"github.com/marzeq/aoc-2024/shared"
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

	var res int
	switch day {
	case 1:
		res = day1.Run(part, shared.GetLines(day))
	case 2:
		res = day2.Run(part, shared.GetLines(day))
	case 3:
		res = day3.Run(part, shared.GetLines(day))
	default:
		res = 0
	}

	fmt.Println(res)
}
