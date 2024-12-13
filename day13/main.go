package day13

import (
	"math"
	"strconv"
	"strings"

	"github.com/marzeq/aoc-2024/shared"
)

type (
	Point2 = shared.Point2[float64]
	Vec2   = shared.Vec2[float64]
	Matrix = shared.Matrix2x2[float64]
)

type Button struct {
	MoveVec Vec2
}

type Machine struct {
	A, B  Button
	Prize Point2
}

func parseInput(lines []string) []Machine {
	machines := []Machine{}

	for i := 0; i < len(lines); i += 4 {
		buttonA := parseButton(lines[i])
		buttonB := parseButton(lines[i+1])
		prize := parsePrize(lines[i+2])

		machines = append(machines, Machine{
			A:     buttonA,
			B:     buttonB,
			Prize: Point2{X: prize[0], Y: prize[1]},
		})
	}
	return machines
}

func parseButton(line string) Button {
	fields := strings.Fields(line)
	x, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(fields[2], "X+"), ","))
	y, _ := strconv.Atoi(strings.TrimPrefix(fields[3], "Y+"))
	return Button{MoveVec: Vec2{X: float64(x), Y: float64(y)}}
}

func parsePrize(line string) [2]float64 {
	fields := strings.Split(line, ",")
	x, _ := strconv.Atoi(strings.TrimPrefix(fields[0], "Prize: X="))
	y, _ := strconv.Atoi(strings.TrimPrefix(fields[1], " Y="))
	return [2]float64{float64(x), float64(y)}
}

// see file problem_formal.tex or it's corresponding pdf

func isWhole(x float64) bool {
	epsilon := 1e-4 // arbitrary value chosen by trial and error
	return shared.Abs(x-math.Round(x)) < epsilon
}

func findLowest3nPlusM(a, b Vec2, p Point2) float64 {
	matrix := Matrix{
		A: a.X, B: b.X,
		C: a.Y, D: b.Y,
	}
	det := matrix.Det()

	if det != 0 {
		inv := matrix.Invert()
		nmVec := inv.MulVec2(p.ToVec2())
		n := nmVec.X
		m := nmVec.Y

		if n <= 0 || m <= 0 || !isWhole(n) || !isWhole(m) {
			return -1
		}

		return 3*math.Round(n) + math.Round(m)
	}

	// sadly in every machine there is only one pair n,m that satisfies,
	// which means i wasted loads of time on the second part in the pdf, but whatever...

	panic("unreachable")
}

func getBestValueIfAny(machine Machine) (int, bool) {
	got := findLowest3nPlusM(machine.A.MoveVec, machine.B.MoveVec, machine.Prize)
	if got == -1 {
		return -1, false
	}
	return int(got), isWhole(got)
}

func part1(machines []Machine) int {
	res := 0

	for _, machine := range machines {
		got, valid := getBestValueIfAny(machine)

		if valid {
			res += got
		}
	}

	return res
}

func part2(machines []Machine) int {
	res := 0

	for _, machine := range machines {
		machine.Prize = machine.Prize.Transform(Vec2{X: 10000000000000, Y: 10000000000000})
		got, valid := getBestValueIfAny(machine)

		if valid {
			res += got
		}
	}

	return res
}

func Run(part int, lines []string) int {
	machines := parseInput(lines)

	if part == 1 {
		return part1(machines)
	} else {
		return part2(machines)
	}
}
