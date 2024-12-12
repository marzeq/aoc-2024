package day11

import (
	"strconv"
	"strings"
)

func parseInput(lines []string) map[int]int {
	stoneCount := make(map[int]int)
	for _, line := range lines {
		for _, num := range strings.Fields(line) {
			stone, _ := strconv.Atoi(num)
			stoneCount[stone]++
		}
	}
	return stoneCount
}

func transform(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	digits := strconv.Itoa(stone)
	if len(digits)%2 == 0 {
		half := len(digits) / 2
		left, _ := strconv.Atoi(digits[:half])
		right, _ := strconv.Atoi(digits[half:])
		return []int{left, right}
	}
	return []int{stone * 2024}
}

func evolve(stoneCount map[int]int, blinks int) map[int]int {
	for i := 0; i < blinks; i++ {
		next := make(map[int]int)
		for stone, count := range stoneCount {
			for _, transformed := range transform(stone) {
				next[transformed] += count
			}
		}
		stoneCount = next
	}
	return stoneCount
}

func part_common(stoneCount map[int]int, blinks int) int {
	stoneCount = evolve(stoneCount, blinks)

	totalStones := 0
	for _, count := range stoneCount {
		totalStones += count
	}
	return totalStones
}

func part1(stoneCount map[int]int) int {
	return part_common(stoneCount, 25)
}

func part2(stoneCount map[int]int) int {
	return part_common(stoneCount, 75)
}

func Run(part int, lines []string) int {
	stoneCount := parseInput(lines)

	if part == 1 {
		return part1(stoneCount)
	} else {
		return part2(stoneCount)
	}
}
