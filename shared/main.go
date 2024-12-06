package shared

import (
	"fmt"
	"os"
	"strings"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Max[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
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
