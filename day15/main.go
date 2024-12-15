package day15

import (
	"github.com/marzeq/aoc-2024/shared"
)

type (
	Point2 = shared.Point2[int]
	Vec2   = shared.Vec2[int]
)

var (
	VecUp    = shared.VecUpOnce
	VecDown  = shared.VecDownOnce
	VecLeft  = shared.VecLeftOnce
	VecRight = shared.VecRightOnce
)

func charToDirection(c rune) Vec2 {
	switch c {
	case '^':
		return VecUp
	case 'v':
		return VecDown
	case '<':
		return VecLeft
	case '>':
		return VecRight
	default:
		panic("unreachable")
	}
}

func directionToString(dir Vec2) string {
	if dir.Equals(VecUp) {
		return "^"
	} else if dir.Equals(VecDown) {
		return "v"
	} else if dir.Equals(VecRight) {
		return ">"
	} else {
		return "<"
	}
}

type Map struct {
	Walls         map[Point2]bool
	Boxes         map[Point2]bool
	Robot         Point2
	Width, Height int
}

func (mp Map) String() (res string) {
	for y := 0; y < mp.Height; y++ {
		for x := 0; x < mp.Width; x++ {
			p := Point2{X: x, Y: y}
			if mp.Boxes[p] {
				res += "O"
			} else if mp.Walls[p] {
				res += "#"
			} else if p == mp.Robot {
				res += "@"
			} else {
				res += "."
			}
		}

		if y < mp.Height-1 {
			res += "\n"
		}
	}

	return
}

func parseInput(lines []string) (mp Map, moves []Vec2) {
	parsingMap := true

	mp.Width = len(lines[0])

	mp.Walls = make(map[Point2]bool)
	mp.Boxes = make(map[Point2]bool)

	for y, line := range lines {
		if line == "" {
			parsingMap = false
			mp.Height = y
			continue
		}

		if parsingMap {
			for x, c := range line {
				switch c {
				case '@':
					mp.Robot = Point2{X: x, Y: y}
				case '#':
					mp.Walls[Point2{X: x, Y: y}] = true
				case 'O':
					mp.Boxes[Point2{X: x, Y: y}] = true
				}
			}
		} else {
			for _, c := range line {
				moves = append(moves, charToDirection(c))
			}
		}
	}

	return
}

func (mp *Map) getContinuousBoxesAmt(pos Point2, dir Vec2) (amount int, lastBox Point2) {
	for ; ; amount++ {
		if !mp.Boxes[pos] {
			return amount, pos.Transform(dir.Negate())
		}
		pos = pos.Transform(dir)
	}
}

func (mp *Map) moveBoxes(firstBox, lastBox Point2, dir Vec2) bool {
	afterLast := lastBox.Transform(dir)
	if mp.Walls[afterLast] {
		return false
	}

	mp.Boxes[firstBox] = false
	mp.Boxes[afterLast] = true
	return true
}

func (mp *Map) processMove(move Vec2) {
	newPos := mp.Robot.Transform(move)

	if mp.Walls[newPos] {
		return
	}

	boxesAheadAmt, lastBox := mp.getContinuousBoxesAmt(newPos, move)

	if boxesAheadAmt == 0 {
		mp.Robot = newPos
		return
	}

	firstBox := newPos

	if mp.moveBoxes(firstBox, lastBox, move) {
		mp.Robot = newPos
	}
}

func calculateCoord(pos Point2) int {
	return pos.X + pos.Y*100
}

func part1(mp Map, moves []Vec2) int {
	for _, move := range moves {
		mp.processMove(move)
	}

	s := 0

	for box, isTrue := range mp.Boxes {
		if !isTrue {
			continue
		}
		s += calculateCoord(box)
	}

	return s
}

func part2() int {
	return 0
}

func Run(part int, lines []string) int {
	mp, moves := parseInput(lines)

	if part == 1 {
		return part1(mp, moves)
	} else {
		return part2()
	}
}
