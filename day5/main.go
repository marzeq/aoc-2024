package day5

import (
	"strconv"
	"strings"
)

type (
	Rules   map[int][]int
	Update  []int
	Updates []Update
)

func parseInput(lines []string) (Rules, Updates) {
	rules := Rules{}
	updates := Updates{}
	parsingRules := true

	for _, line := range lines {
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			split := strings.Split(line, "|")
			a, err := strconv.Atoi(split[0])
			if err != nil {
				panic("error parsing input")
			}
			b, err := strconv.Atoi(split[1])
			if err != nil {
				panic("error parsing input")
			}

			rules[a] = append(rules[a], b)
		} else {
			split := strings.Split(line, ",")
			update := Update{}

			for _, num := range split {
				a, err := strconv.Atoi(num)
				if err != nil {
					panic("error parsing input")
				}
				update = append(update, a)
			}

			updates = append(updates, update)
		}
	}

	return rules, updates
}

func isValidUpdate(update Update, rules Rules) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for a, dependencies := range rules {
		if posA, exists := position[a]; exists {
			for _, b := range dependencies {
				if posB, exists := position[b]; exists {
					if posA > posB {
						return false
					}
				}
			}
		}
	}

	return true
}

func part1(rules Rules, updates Updates) int {
	res := 0

	for _, update := range updates {
		if isValidUpdate(update, rules) {
			midIndex := len(update) / 2
			res += update[midIndex]
		}
	}

	return res
}

func fixUpdate(update Update, rules Rules) Update {
	incomingDegree := make(map[int]int)
	graph := make(map[int][]int)
	present := make(map[int]bool)

	for _, page := range update {
		incomingDegree[page] = 0
		graph[page] = []int{}
		present[page] = true
	}

	for a, dependencies := range rules {
		if !present[a] {
			continue
		}
		for _, b := range dependencies {
			if present[b] {
				graph[a] = append(graph[a], b)
				incomingDegree[b]++
			}
		}
	}

	fixed := Update{}
	queue := []int{}

	for page, degree := range incomingDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fixed = append(fixed, current)

		for _, neighbor := range graph[current] {
			incomingDegree[neighbor]--
			if incomingDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return fixed
}

func part2(rules Rules, updates Updates) int {
	res := 0

	for _, update := range updates {
		if isValidUpdate(update, rules) {
			continue
		}

		fixed := fixUpdate(update, rules)

		midIndex := len(fixed) / 2
		res += fixed[midIndex]
	}

	return res
}

func Run(part int, lines []string) int {
	rules, updates := parseInput(lines)

	if part == 1 {
		return part1(rules, updates)
	} else {
		return part2(rules, updates)
	}
}
