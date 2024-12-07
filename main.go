package main

import (
	"os"
	"strconv"
	"time"

	"github.com/marzeq/aoc-2024/day1"
	"github.com/marzeq/aoc-2024/day2"
	"github.com/marzeq/aoc-2024/day3"
	"github.com/marzeq/aoc-2024/day4"
	"github.com/marzeq/aoc-2024/day5"
	"github.com/marzeq/aoc-2024/day6"
	"github.com/marzeq/aoc-2024/day7"
	"github.com/marzeq/aoc-2024/shared"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go run . [day] [part]")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Invalid day arg")
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Invalid part arg")
	}

	lines := shared.GetLines(day)

	var res int
	tstart := time.Now()
	switch day {
	case 1:
		res = day1.Run(part, lines)
	case 2:
		res = day2.Run(part, lines)
	case 3:
		res = day3.Run(part, lines)
	case 4:
		res = day4.Run(part, lines)
	case 5:
		res = day5.Run(part, lines)
	case 6:
		res = day6.Run(part, lines)
	case 7:
		res = day7.Run(part, lines)
	default:
		panic("please update main.go")
	}
	dt := time.Now().Sub(tstart).Abs().Milliseconds()

	println(res)
	println("time took:", dt, "ms")
}
