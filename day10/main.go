package day10

import "strconv"

var directions = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func parseInput(lines []string) [][]int {
	topographicMap := make([][]int, len(lines))
	for i, line := range lines {
		topographicMap[i] = make([]int, len(line))
		for j, ch := range line {

			value, _ := strconv.Atoi(string(ch))
			topographicMap[i][j] = value
		}
	}
	return topographicMap
}

func findTrailheadScore(mapData [][]int, x, y int) int {
	reached := map[[2]int]bool{}
	var dfs func(cx, cy, height int)

	dfs = func(cx, cy, height int) {
		reached[[2]int{cx, cy}] = true

		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]

			if nx >= 0 && ny >= 0 && nx < len(mapData) && ny < len(mapData[0]) && !reached[[2]int{nx, ny}] && mapData[nx][ny] == height+1 {
				dfs(nx, ny, mapData[nx][ny])
			}
		}
	}

	dfs(x, y, 0)

	score := 0
	for pos := range reached {
		if mapData[pos[0]][pos[1]] == 9 {
			score++
		}
	}

	return score
}

func part1(mapData [][]int) int {
	totalScore := 0

	for i := 0; i < len(mapData); i++ {
		for j := 0; j < len(mapData[0]); j++ {
			if mapData[i][j] == 0 {
				totalScore += findTrailheadScore(mapData, i, j)
			}
		}
	}

	return totalScore
}

func findTrailheadRating(mapData [][]int, x, y int) int {
	trailCount := 0
	var dfs func(cx, cy, height int, path map[[2]int]bool)

	dfs = func(cx, cy, height int, path map[[2]int]bool) {
		path[[2]int{cx, cy}] = true

		if height == 9 {
			trailCount++
			return
		}

		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]

			if nx >= 0 && ny >= 0 && nx < len(mapData) && ny < len(mapData[0]) && !path[[2]int{nx, ny}] && mapData[nx][ny] == height+1 {
				newPath := make(map[[2]int]bool)
				for k := range path {
					newPath[k] = true
				}
				dfs(nx, ny, mapData[nx][ny], newPath)
			}
		}
	}

	dfs(x, y, 0, map[[2]int]bool{})

	return trailCount
}

func part2(mapData [][]int) int {
	totalRating := 0

	for i := 0; i < len(mapData); i++ {
		for j := 0; j < len(mapData[0]); j++ {
			if mapData[i][j] == 0 {
				totalRating += findTrailheadRating(mapData, i, j)
			}
		}
	}

	return totalRating
}

func Run(part int, lines []string) int {
	mapData := parseInput(lines)

	if part == 1 {
		return part1(mapData)
	} else {
		return part2(mapData)
	}
}
