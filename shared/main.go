package shared

import (
	"fmt"
	"os"
	"strings"
)

type Vec2[T Number] struct {
	X, Y T
}

func (vec Vec2[T]) Negate() Vec2[T] {
	return Vec2[T]{-vec.X, -vec.Y}
}

func (vec Vec2[T]) Equals(v2 Vec2[T]) bool {
	return vec.X == v2.X && vec.Y == v2.Y
}

var (
	VecUpOnce    = Vec2[int]{X: 0, Y: -1}
	VecDownOnce  = Vec2[int]{X: 0, Y: 1}
	VecLeftOnce  = Vec2[int]{X: -1, Y: 0}
	VecRightOnce = Vec2[int]{X: 1, Y: 0}
)

type Point2[T Number] struct {
	X, Y T
}

func (point Point2[T]) Transform(vec Vec2[T]) Point2[T] {
	return point.TransformTimes(vec, 1)
}

func (point Point2[T]) TransformTimes(vec Vec2[T], t T) Point2[T] {
	return Point2[T]{point.X + t*vec.X, point.Y + t*vec.Y}
}

func (point Point2[T]) DeltaVector(point2 Point2[T]) Vec2[T] {
	return Vec2[T]{point.X - point2.X, point.Y - point2.Y}
}

func (point Point2[T]) ToVec2() Vec2[T] {
	return Vec2[T]{point.X, point.Y}
}

func (point Point2[T]) Equals(p2 Point2[T]) bool {
	return point.X == p2.X && point.Y == p2.Y
}

type Matrix2x2[T Number] struct {
	A, B, C, D T
}

func (m Matrix2x2[T]) Det() T {
	return m.A*m.D - m.B*m.C
}

func (m Matrix2x2[T]) MulScalar(s T) Matrix2x2[T] {
	return Matrix2x2[T]{
		A: m.A * s, B: m.B * s, C: m.C * s, D: m.D * s,
	}
}

func (m Matrix2x2[T]) MulVec2(v Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: m.A*v.X + m.B*v.Y,
		Y: m.C*v.X + m.D*v.Y,
	}
}

func (m Matrix2x2[T]) Invert() Matrix2x2[T] {
	return Matrix2x2[T]{
		A: m.D, B: -m.B,
		C: -m.C, D: m.A,
	}.MulScalar(1 / m.Det())
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func UniqueCombinations[T any](input []T, size int) [][]T {
	result := [][]T{}
	comb := make([]T, size)

	var helper func(start, depth int)
	helper = func(start, depth int) {
		if depth == size {
			temp := make([]T, size)
			copy(temp, comb)
			result = append(result, temp)
			return
		}

		for i := start; i < len(input); i++ {
			comb[depth] = input[i]
			helper(i+1, depth+1)
		}
	}

	helper(0, 0)
	return result
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Sign[T Number](a T) T {
	if a > 0 {
		return 1
	} else if a == 0 {
		return 0
	}

	return -1
}

func GetLines(day int) []string {
	fname := fmt.Sprintf("inputs/%d.txt", day)
	f, err := os.ReadFile(fname)
	if err != nil {
		panic(fmt.Sprintf("Create file %s to run", fname))
	}

	ret := strings.Split(string(f), "\n")
	return ret[:len(ret)-1]
}

func Copy2DArray[T any](original [][]T) [][]T {
	cpy := make([][]T, len(original))
	for i := range original {
		cpy[i] = make([]T, len(original[i]))
		copy(cpy[i], original[i])
	}
	return cpy
}
