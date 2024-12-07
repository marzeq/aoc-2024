package day7

import (
	"math"
	"strconv"
	"strings"
)

type Equation struct {
	testValue int
	numbers   []int
}

func parseInput(lines []string) []Equation {
	equations := []Equation{}

	for _, line := range lines {
		equation := Equation{}
		split := strings.Split(line, ": ")

		testValue, err := strconv.Atoi(split[0])
		if err != nil {
			panic("failed to parse input")
		}

		equation.testValue = testValue

		numbersStrs := strings.Split(split[1], " ")

		numbers := []int{}
		for _, nS := range numbersStrs {
			n, err := strconv.Atoi(nS)
			if err != nil {
				panic("failed to parse input")
			}

			numbers = append(numbers, n)
		}

		equation.numbers = numbers

		equations = append(equations, equation)
	}

	return equations
}

func canSolve(equation Equation, currentIndex int, currentValue int, part2 bool) bool {
	if currentValue > equation.testValue {
		return false
	}

	if currentIndex == len(equation.numbers) {
		return currentValue == equation.testValue
	}

	if canSolve(equation, currentIndex+1, currentValue+equation.numbers[currentIndex], part2) {
		return true
	}

	if canSolve(equation, currentIndex+1, currentValue*equation.numbers[currentIndex], part2) {
		return true
	}

	if part2 && canSolve(equation, currentIndex+1, concat(currentValue, equation.numbers[currentIndex]), part2) {
		return true
	}

	return false
}

func getWaysToSolve(equation Equation, part2 bool) int {
	if len(equation.numbers) == 0 {
		return 0
	}

	if canSolve(equation, 1, equation.numbers[0], part2) {
		return equation.testValue
	}

	return 0
}

func part1(equations []Equation) int {
	total := 0
	for _, eq := range equations {
		total += getWaysToSolve(eq, false)
	}
	return total
}

func concat(a, b int) int {
	if b == 0 {
		return a * 10
	}

	bDigits := int(math.Log10(float64(b))) + 1

	return a*int(math.Pow(10, float64(bDigits))) + b
}

func part2(equations []Equation) int {
	total := 0
	for _, eq := range equations {
		total += getWaysToSolve(eq, true)
	}
	return total
}

func Run(part int, lines []string) int {
	equations := parseInput(lines)

	if part == 1 {
		return part1(equations)
	} else {
		return part2(equations)
	}
}
