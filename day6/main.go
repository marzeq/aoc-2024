package day6

import (
	"sync"

	"github.com/marzeq/aoc-2024/shared"
)

type (
	Map [][]bool
	Pos struct {
		X, Y int
	}
	VisitedMap map[Pos]bool
)

func parseInput(lines []string) (Map, Pos) {
	labmap := Map{}
	startpos := Pos{}

	for y, line := range lines {
		mapline := []bool{}

		for x, ch := range line {
			mapline = append(mapline, ch != '#')
			if ch == '^' {
				startpos.X = x
				startpos.Y = y
			}
		}

		labmap = append(labmap, mapline)
	}

	return labmap, startpos
}

func isPosInsideMap(pos Pos, labmap Map) bool {
	return pos.Y >= 0 && pos.Y < len(labmap) &&
		pos.X >= 0 && pos.X < len(labmap[pos.Y])
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func move(labmap Map, currentPos Pos, currentDirection Direction) (bool, Pos, Direction) {
	newPos := getNewPos(currentPos, currentDirection)
	newDirection := currentDirection

	if !isPosInsideMap(newPos, labmap) {
		return false, Pos{}, NORTH
	}

	if !labmap[newPos.Y][newPos.X] {
		newDirection = getNextDirection(currentDirection)
		newPos = getNewPos(currentPos, newDirection)

		if !labmap[newPos.Y][newPos.X] {
			return move(labmap, currentPos, newDirection)
		}
	}

	return true, newPos, newDirection
}

func getNextDirection(direction Direction) Direction {
	switch direction {
	case NORTH:
		return EAST
	case EAST:
		return SOUTH
	case SOUTH:
		return WEST
	case WEST:
		return NORTH
	}

	panic("bruh")
}

func getNewPos(pos Pos, direction Direction) Pos {
	switch direction {
	case NORTH:
		return Pos{pos.X, pos.Y - 1}
	case EAST:
		return Pos{pos.X + 1, pos.Y}
	case SOUTH:
		return Pos{pos.X, pos.Y + 1}
	case WEST:
		return Pos{pos.X - 1, pos.Y}
	}

	panic("bruh")
}

func simulateGuardMovement(labmap Map, startpos Pos) VisitedMap {
	visited := make(VisitedMap)

	currentPos := startpos
	currentDirection := NORTH

	for {
		visited[currentPos] = true

		validmove, newPos, newDirection := move(labmap, currentPos, currentDirection)

		if !validmove {
			break
		}
		currentPos = newPos
		currentDirection = newDirection
	}

	return visited
}

func part1(labmap Map, startpos Pos) int {
	return len(simulateGuardMovement(labmap, startpos))
}

func isGuardStuckInLoop(labmap_og Map, startpos Pos, obstruction Pos) bool {
	labmap := shared.Copy2DArray(labmap_og)
	labmap[obstruction.Y][obstruction.X] = false
	visitedSequence := map[Pos]map[Direction]bool{}
	currentPos := startpos
	currentDirection := NORTH

	for {
		if _, exists := visitedSequence[currentPos]; !exists {
			visitedSequence[currentPos] = make(map[Direction]bool)
		}
		if visitedSequence[currentPos][currentDirection] {
			return true
		}

		visitedSequence[currentPos][currentDirection] = true

		validmove, newPos, newDirection := move(labmap, currentPos, currentDirection)

		if !validmove {
			return false
		}
		currentPos = newPos
		currentDirection = newDirection
	}
}

/* func part2(labmap Map, startpos Pos) int {
	visited := simulateGuardMovement(labmap, startpos)

	loopCount := 0
	for pos := range visited {
		if pos == startpos {
			continue
		}

		if isGuardStuckInLoop(labmap, startpos, pos) {
			loopCount++
		}
	}

	return loopCount
} */

// kinda cheating but if i pay for 24 threads i get to use 24 threads, so deal with it!!!!
func part2(labmap Map, startpos Pos) int {
	visited := simulateGuardMovement(labmap, startpos)

	loopCountChan := make(chan int, len(visited))

	var waitGroup sync.WaitGroup

	for pos := range visited {
		if pos == startpos {
			continue
		}

		waitGroup.Add(1)

		go func(checkPos Pos) {
			defer waitGroup.Done()

			if isGuardStuckInLoop(labmap, startpos, checkPos) {
				loopCountChan <- 1
			} else {
				loopCountChan <- 0
			}
		}(pos)
	}

	go func() {
		waitGroup.Wait()
		close(loopCountChan)
	}()

	totalLoops := 0
	for count := range loopCountChan {
		totalLoops += count
	}

	return totalLoops
}

func Run(part int, lines []string) int {
	labmap, startpos := parseInput(lines)

	if part == 1 {
		return part1(labmap, startpos)
	} else {
		return part2(labmap, startpos)
	}
}
