package day8

import (
	"github.com/marzeq/aoc-2024/shared"
)

type (
	Point2 = shared.Point2[int]
	Vec2   = shared.Vec2[int]
)

type AntennasMap map[rune]([]Point2)

func parseInput(lines []string) (AntennasMap, int, int) {
	antennasMap := make(AntennasMap)

	for y, line := range lines {
		for x, freq := range line {
			if freq != '.' && freq != '#' {
				antennasMap[freq] = append(antennasMap[freq], Point2{X: x, Y: y})
			}
		}
	}

	sizeY := len(lines)
	sizeX := len(lines[0])

	return antennasMap, sizeX, sizeY
}

func getTwoAntinodes(p1, p2 Point2) (Point2, Point2) {
	vec := p1.DeltaVector(p2)

	return p1.Transform(vec), p2.Transform(vec.Negate())
}

func getAllAntinodes(p1, p2 Point2, sizeX, sizeY int, existingAnodes []Point2) []Point2 {
	anodes := []Point2{}

	vec := p1.DeltaVector(p2)

	np := p1.Transform(vec)

	for isInBounds(np, sizeX, sizeY) {
		if !isPointUsed(existingAnodes, np) {
			anodes = append(anodes, np)
		}
		np = np.Transform(vec)
	}

	nvec := vec.Negate()

	np = p2.Transform(nvec)

	for isInBounds(np, sizeX, sizeY) {
		if !isPointUsed(existingAnodes, np) {
			anodes = append(anodes, np)
		}
		np = np.Transform(nvec)
	}

	return anodes
}

func isPointUsed(anodes []Point2, point Point2) bool {
	used := false

	for _, anode := range anodes {
		if anode.X == point.X && anode.Y == point.Y {
			used = true
			break
		}
	}

	return used
}

func isInBounds(p Point2, sizeX int, sizeY int) bool {
	return p.X >= 0 && p.X < sizeX && p.Y >= 0 && p.Y < sizeY
}

func part1(antennasMap AntennasMap, sizeX int, sizeY int) int {
	anodes := []Point2{}

	for _, points := range antennasMap {
		pairs := shared.UniqueCombinations(points, 2)

		for _, pair := range pairs {
			a1, a2 := getTwoAntinodes(pair[0], pair[1])

			if isInBounds(a1, sizeX, sizeY) && !isPointUsed(anodes, a1) {
				anodes = append(anodes, a1)
			}

			if isInBounds(a2, sizeX, sizeY) && !isPointUsed(anodes, a2) {
				anodes = append(anodes, a2)
			}
		}
	}

	return len(anodes)
}

func part2(antennasMap AntennasMap, sizeX int, sizeY int) int {
	anodes := []Point2{}

	for _, points := range antennasMap {
		pairs := shared.UniqueCombinations(points, 2)

		for _, pair := range pairs {
			p1, p2 := pair[0], pair[1]

			for _, anode := range getAllAntinodes(p1, p2, sizeX, sizeY, anodes) {
				anodes = append(anodes, anode)
			}

			if !isPointUsed(anodes, p1) {
				anodes = append(anodes, p1)
			}
			if !isPointUsed(anodes, p2) {
				anodes = append(anodes, p2)
			}
		}
	}

	return len(anodes)
}

func Run(part int, lines []string) int {
	antennasMap, sizeX, sizeY := parseInput(lines)

	if part == 1 {
		return part1(antennasMap, sizeX, sizeY)
	} else {
		return part2(antennasMap, sizeX, sizeY)
	}
}
