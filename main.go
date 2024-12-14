package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/marzeq/aoc-2024/day1"
	"github.com/marzeq/aoc-2024/day10"
	"github.com/marzeq/aoc-2024/day11"
	"github.com/marzeq/aoc-2024/day12"
	"github.com/marzeq/aoc-2024/day13"
	"github.com/marzeq/aoc-2024/day14"
	"github.com/marzeq/aoc-2024/day2"
	"github.com/marzeq/aoc-2024/day3"
	"github.com/marzeq/aoc-2024/day4"
	"github.com/marzeq/aoc-2024/day5"
	"github.com/marzeq/aoc-2024/day6"
	"github.com/marzeq/aoc-2024/day7"
	"github.com/marzeq/aoc-2024/day8"
	"github.com/marzeq/aoc-2024/day9"
	"github.com/marzeq/aoc-2024/shared"
)

func copyToCB(text string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("clip")
	case "linux":
		if os.Getenv("WAYLAND_DISPLAY") != "" {
			cmd = exec.Command("wl-copy")
		} else {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		}
	}

	cmd.Stdin = strings.NewReader(text)
	cmd.Run()
}

func printRes(res int, tstart time.Time) {
	dt := time.Now().Sub(tstart).Abs().Milliseconds()

	fmt.Println(res)
	fmt.Println("time took:", dt, "ms")
	copyToCB(strconv.Itoa(res))
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go run . [day] [part (1|2)]")
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

	tstart := time.Now()
	switch day {
	case 1:
		printRes(day1.Run(part, lines), tstart)
	case 2:
		printRes(day2.Run(part, lines), tstart)
	case 3:
		printRes(day3.Run(part, lines), tstart)
	case 4:
		printRes(day4.Run(part, lines), tstart)
	case 5:
		printRes(day5.Run(part, lines), tstart)
	case 6:
		printRes(day6.Run(part, lines), tstart)
	case 7:
		printRes(day7.Run(part, lines), tstart)
	case 8:
		printRes(day8.Run(part, lines), tstart)
	case 9:
		printRes(day9.Run(part, lines), tstart)
	case 10:
		printRes(day10.Run(part, lines), tstart)
	case 11:
		printRes(day11.Run(part, lines), tstart)
	case 12:
		printRes(day12.Run(part, lines), tstart)
	case 13:
		printRes(day13.Run(part, lines), tstart)
	case 14:
		printRes(day14.Run(part, lines), tstart)
	default:
		panic("please update main.go")
	}
}
