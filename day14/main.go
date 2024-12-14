package day14

import (
	"strconv"
	"strings"

	"github.com/marzeq/aoc-2024/shared"
)

const (
	HEIGHT = 103
	WIDTH  = 101
)

type (
	Point2 = shared.Point2[int]
	Vec2   = shared.Vec2[int]
)

type Robot struct {
	Pos Point2
	Vel Vec2
}

func parseInput(lines []string) []Robot {
	robots := []Robot{}
	for _, line := range lines {
		fields := strings.Fields(line)

		robot := Robot{}

		posSplit := strings.Split(strings.TrimPrefix(fields[0], "p="), ",")

		robot.Pos.X, _ = strconv.Atoi(posSplit[0])
		robot.Pos.Y, _ = strconv.Atoi(posSplit[1])

		velSplit := strings.Split(strings.TrimPrefix(fields[1], "v="), ",")

		robot.Vel.X, _ = strconv.Atoi(velSplit[0])
		robot.Vel.Y, _ = strconv.Atoi(velSplit[1])

		robots = append(robots, robot)
	}

	return robots
}

func calcPositionAfterT(robot Robot, t int) Point2 {
	transformed := robot.Pos.TransformTimes(robot.Vel, t)

	transformed.X = transformed.X % WIDTH
	transformed.Y = transformed.Y % HEIGHT

	if transformed.X < 0 {
		transformed.X = WIDTH + transformed.X
	}

	if transformed.Y < 0 {
		transformed.Y = HEIGHT + transformed.Y
	}

	return transformed
}

func calcNextPosition(robot Robot) Point2 {
	return calcPositionAfterT(robot, 1)
}

func getQuadrants(poss []Point2) (q1 []Point2, q2 []Point2, q3 []Point2, q4 []Point2) {
	midWidth := WIDTH / 2
	midHeight := HEIGHT / 2

	for _, pos := range poss {
		if pos.Y < midHeight && pos.X < midWidth {
			q1 = append(q1, pos)
		} else if pos.Y < midHeight && pos.X > midWidth {
			q2 = append(q2, pos)
		} else if pos.Y > midHeight && pos.X < midWidth {
			q3 = append(q3, pos)
		} else if pos.Y > midHeight && pos.X > midWidth {
			q4 = append(q4, pos)
		}
	}

	return q1, q2, q3, q4
}

func part1(robots []Robot) int {
	newPoss := []Point2{}

	for _, robot := range robots {
		newPoss = append(newPoss, calcPositionAfterT(robot, 100))
	}

	q1, q2, q3, q4 := getQuadrants(newPoss)

	return len(q1) * len(q2) * len(q3) * len(q4)
}

func part2(robots []Robot) int {
	for t := 1; t < 100_000; t++ {
		newPoss := []Point2{}
		for _, robot := range robots {
			newPoss = append(newPoss, calcPositionAfterT(robot, t))
		}

		if hasTree(newPoss) {
			return t
		}
	}

	return -1
}

func Run(part int, lines []string) int {
	if WIDTH%2 == 0 || HEIGHT%2 == 0 {
		panic("WIDTH and HEIGHT mustn't be even")
	}

	robots := parseInput(lines)

	if part == 1 {
		return part1(robots)
	} else {
		return part2(robots)
	}
}
