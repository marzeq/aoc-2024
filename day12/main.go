package day12

type index struct {
	r, c int
}

var directions = []index{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

var cornerToOrtho = map[index][]index{
	{1, 1}:   {{1, 0}, {0, 1}},
	{-1, 1}:  {{-1, 0}, {0, 1}},
	{1, -1}:  {{1, 0}, {0, -1}},
	{-1, -1}: {{-1, 0}, {0, -1}},
}

func Run(part int, lines []string) int {
	input := parseInput(lines)

	if part == 1 {
		return part1(input)
	} else {
		return part2(input)
	}
}

func parseInput(lines []string) [][]rune {
	puzzleInput := [][]rune{}

	for _, line := range lines {
		temp := make([]rune, 0, len(line))
		for _, r := range line {
			temp = append(temp, r)
		}
		puzzleInput = append(puzzleInput, temp)
	}
	return puzzleInput
}

func part1(grid [][]rune) int {
	cost := 0

	visited := make(map[index]bool)
	for r := range grid {
		for c := range grid[r] {
			if !visited[index{r: r, c: c}] {
				var area, perimeter int
				explore(index{r: r, c: c}, grid, visited, grid[r][c], &area, &perimeter)
				cost += (area + 1) * perimeter
			}
		}
	}
	return cost
}

func explore(curr index, grid [][]rune, visited map[index]bool, currPlant rune, area, perimeter *int) {
	maxRow, maxCol := len(grid), len(grid[0])
	neighbourCount := 0

	visited[curr] = true

	for _, dir := range directions {
		nextIndex := index{r: curr.r + dir.r, c: curr.c + dir.c}
		if isWithinBounds(nextIndex, maxRow, maxCol) && grid[nextIndex.r][nextIndex.c] == currPlant {
			neighbourCount++
			if !visited[nextIndex] {
				(*area)++
				explore(nextIndex, grid, visited, currPlant, area, perimeter)
			}
		}
	}
	(*perimeter) += 4 - neighbourCount
}

func part2(grid [][]rune) (cost int) {
	visited := make(map[index]bool)
	for r := range grid {
		for c := range grid[r] {
			if !visited[index{r: r, c: c}] {
				var area, corners int
				exploreV2(index{r: r, c: c}, grid, visited, grid[r][c], &area, &corners)
				cost += (area + 1) * corners
			}
		}
	}
	return
}

func exploreV2(curr index, grid [][]rune, visited map[index]bool, currPlant rune, area, corners *int) {
	maxRow, maxCol := len(grid), len(grid[0])
	visited[curr] = true

	for _, dir := range directions {
		nextIndex := index{r: curr.r + dir.r, c: curr.c + dir.c}
		if isWithinBounds(nextIndex, maxRow, maxCol) && grid[nextIndex.r][nextIndex.c] == currPlant {
			if !visited[nextIndex] {
				(*area)++
				exploreV2(nextIndex, grid, visited, currPlant, area, corners)
			}
		}
	}

	for corner, pair := range cornerToOrtho {
		c := index{r: curr.r + corner.r, c: curr.c + corner.c}
		i1 := index{r: curr.r + pair[0].r, c: curr.c + pair[0].c}
		i2 := index{r: curr.r + pair[1].r, c: curr.c + pair[1].c}

		if !match(i1, curr, grid) && !match(i2, curr, grid) {
			(*corners)++
		}
		if match(i1, curr, grid) && match(i2, curr, grid) && !match(curr, c, grid) {
			(*corners)++
		}
	}
}

func match(i1, i2 index, grid [][]rune) bool {
	maxRow, maxCol := len(grid), len(grid[0])

	if !isWithinBounds(i1, maxRow, maxCol) && !isWithinBounds(i2, maxRow, maxCol) {
		return true
	} else if isWithinBounds(i1, maxRow, maxCol) && isWithinBounds(i2, maxRow, maxCol) {
		p1, p2 := grid[i1.r][i1.c], grid[i2.r][i2.c]
		return p1 == p2
	} else {
		return false
	}
}

func isWithinBounds(i index, maxRow, maxCol int) bool {
	return i.r >= 0 && i.c >= 0 && i.r < maxRow && i.c < maxCol
}
