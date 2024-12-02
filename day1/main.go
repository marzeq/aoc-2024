package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/marzeq/aoc-2024/shared"
)

func parseInput(lines []string) ([]int, []int) {
	nums1 := []int{}
	nums2 := []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		linespl := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(linespl[0])
		nums1 = append(nums1, num1)

		num2, _ := strconv.Atoi(linespl[1])
		nums2 = append(nums2, num2)
	}

	return nums1, nums2
}

func part1(nums1 []int, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	res := 0

	for i, n1 := range nums1 {
		n2 := nums2[i]

		res += shared.Abs(n1 - n2)
	}

	return res
}

func part2(nums1 []int, nums2 []int) int {
	occurmap := make(map[int]int)

	for _, n2 := range nums2 {
		occurmap[n2] = occurmap[n2] + 1
	}

	res := 0

	for _, n1 := range nums1 {
		res += occurmap[n1] * n1
	}

	return res
}

func Run(part int, lines []string) {
	nums1, nums2 := parseInput(lines)
	var res int

	if part == 1 {
		res = part1(nums1, nums2)
	} else {
		res = part2(nums1, nums2)
	}

	fmt.Println(res)
}
