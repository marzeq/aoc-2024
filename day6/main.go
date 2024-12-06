package day6

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

		newPos := getNewPos(currentPos, currentDirection)

		if !isPosInsideMap(newPos, labmap) {
			break
		}

		if labmap[newPos.Y][newPos.X] {
			currentPos = newPos
		} else {
			currentDirection = getNextDirection(currentDirection)
			currentPos = getNewPos(currentPos, currentDirection)
		}
	}

	return visited
}

func part1(labmap Map, startpos Pos) int {
	return len(simulateGuardMovement(labmap, startpos))
}

func part2(labmap Map, startpos Pos) int {
	return 0
}

func Run(part int, lines []string) int {
	labmap, startpos := parseInput(lines)

	if part == 1 {
		return part1(labmap, startpos)
	} else {
		return part2(labmap, startpos)
	}
}
