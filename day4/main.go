package day4

func getColumnsFromLines(lines []string) []string {
	cols := []string{}

	for _, line := range lines {
		for i, ch := range line {
			if i >= len(cols) {
				cols = append(cols, "")
			}

			cols[i] += string(ch)
		}
	}

	return cols
}

func getDiagonalsFromLines(lines []string) ([]string, []string) {
	numRows := len(lines)
	numCols := len(lines[0])

	leftDiagonals := []string{}
	rightDiagonals := []string{}

	for startCol := 0; startCol < numCols; startCol++ {
		diagonal := []rune{}
		row, col := 0, startCol
		for row < numRows && col < numCols {
			diagonal = append(diagonal, rune(lines[row][col]))
			row++
			col++
		}
		leftDiagonals = append(leftDiagonals, string(diagonal))
	}

	for startRow := 1; startRow < numRows; startRow++ {
		diagonal := []rune{}
		row, col := startRow, 0
		for row < numRows && col < numCols {
			diagonal = append(diagonal, rune(lines[row][col]))
			row++
			col++
		}
		leftDiagonals = append(leftDiagonals, string(diagonal))
	}

	for startCol := 0; startCol < numCols; startCol++ {
		diagonal := []rune{}
		row, col := 0, startCol
		for row < numRows && col >= 0 {
			diagonal = append(diagonal, rune(lines[row][col]))
			row++
			col--
		}
		rightDiagonals = append(rightDiagonals, string(diagonal))
	}

	for startRow := 1; startRow < numRows; startRow++ {
		diagonal := []rune{}
		row, col := startRow, numCols-1
		for row < numRows && col >= 0 {
			diagonal = append(diagonal, rune(lines[row][col]))
			row++
			col--
		}
		rightDiagonals = append(rightDiagonals, string(diagonal))
	}

	return leftDiagonals, rightDiagonals
}

func countOuccurences(s string, word string) int {
	res := 0

	letterofword := 0
	for _, ch := range s {
		if ch != rune(word[letterofword]) {
			if ch == rune(word[0]) {
				letterofword = 1
			} else {
				letterofword = 0
			}
			continue
		}

		letterofword++

		if letterofword == len(word) {
			res++
			letterofword = 0
		}
	}

	return res
}

func countOccurrencesWithRev(s string, word string) int {
	return countOuccurences(s, word) + countOuccurences(s, reverseString(word))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func part1(lines []string, cols []string, diags1 []string, diags2 []string) int {
	res := 0

	for _, line := range lines {
		res += countOccurrencesWithRev(line, "XMAS")
	}

	for _, col := range cols {
		res += countOccurrencesWithRev(col, "XMAS")
	}

	for _, diag := range diags1 {
		res += countOccurrencesWithRev(diag, "XMAS")
	}

	for _, diag := range diags2 {
		res += countOccurrencesWithRev(diag, "XMAS")
	}

	return res
}

func getSquaresFromLines(lines []string, size int) [][]string {
	squares := [][]string{}

	if size > len(lines[0]) {
		panic("size of square bigger than input")
	}

	for linei := 0; linei <= len(lines)-size; linei++ {
		lineswecareabout := lines[linei : size+linei]

		for coli := 0; coli <= len(lineswecareabout[0])-size; coli++ {
			square := []string{}

			for _, line := range lineswecareabout {
				square = append(square, line[coli:size+coli])
			}

			squares = append(squares, square)
		}
	}

	return squares
}

func isSquareValid(square []string) bool {
	return square[1][1] == 'A' && ((square[0][0] == 'M' && square[2][2] == 'S') || (square[0][0] == 'S' && square[2][2] == 'M')) && ((square[0][2] == 'M' && square[2][0] == 'S') || (square[0][2] == 'S' && square[2][0] == 'M'))
}

func part2(lines []string) int {
	res := 0

	squares := getSquaresFromLines(lines, 3)

	for _, square := range squares {
		if isSquareValid(square) {
			res++
		}
	}

	return res
}

func Run(part int, lines []string) int {
	if part == 1 {
		cols := getColumnsFromLines(lines)
		diags1, diags2 := getDiagonalsFromLines(lines)
		return part1(lines, cols, diags1, diags2)
	} else {
		return part2(lines)
	}
}
