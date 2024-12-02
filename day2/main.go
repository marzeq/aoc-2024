package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/marzeq/aoc-2024/shared"
)

func parseInput(lines []string) [][]int {
	reports := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		report := []int{}

		split := strings.Split(line, " ")

		for _, n := range split {
			conv, _ := strconv.Atoi(n)
			report = append(report, conv)
		}

		reports = append(reports, report)
	}

	return reports
}

func check(a int, b int, isfirst bool, shouldincrease bool) (bool, int) {
	diff := b - a

	if diff < 0 {
		if !isfirst && shouldincrease {
			return false, diff
		}
	} else {
		if !isfirst && !shouldincrease {
			return false, diff
		}
	}

	abs := shared.Abs(diff)

	if abs == 0 || abs > 3 {
		return false, diff
	}

	return true, diff
}

func isSafe(report []int) bool {
	increasing := true
	for i := 0; i < len(report)-1; i++ {
		ok, diff := check(report[i], report[i+1], i == 0, increasing)
		if !ok {
			return false
		}
		if diff < 0 {
			increasing = false
		}
	}
	return true
}

func part1(reports [][]int) int {
	res := 0

	for _, report := range reports {
		increasing := true
		valid := true

		for i, n := range report {
			if i == len(report)-1 {
				break
			}

			ok, diff := check(n, report[i+1], i == 0, increasing)

			if !ok {
				valid = false
				break
			}

			if diff < 0 {
				increasing = false
			}
		}

		if valid {
			res++
		}
	}

	return res
}

func part2(reports [][]int) int {
	res := 0

	for _, report := range reports {
		if isSafe(report) {
			res++
			continue
		}

		valid := false
		for i := range report {
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)

			if isSafe(modifiedReport) {
				valid = true
				break
			}
		}

		if valid {
			res++
		}
	}

	return res
}

func Run(part int, lines []string) {
	reports := parseInput(lines)

	var res int

	if part == 1 {
		res = part1(reports)
	} else {
		res = part2(reports)
	}

	fmt.Println(res)
}
