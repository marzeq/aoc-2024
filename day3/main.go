package day3

import (
	"strconv"
	"strings"
)

type tokeniser struct {
	contents  []rune
	currIndex int
}

func newTokeniser(contents string) tokeniser {
	return tokeniser{
		contents:  []rune(contents),
		currIndex: 0,
	}
}

func (t *tokeniser) PeekAhead(i int) rune {
	if t.currIndex+i == len(t.contents) {
		return 0
	}

	return t.contents[t.currIndex+i]
}

func (t *tokeniser) Peek() rune {
	return t.PeekAhead(0)
}

func (t *tokeniser) Consume() rune {
	c := t.Peek()

	t.currIndex++

	if t.currIndex > len(t.contents) {
		t.currIndex = len(t.contents)
	}

	return c
}

func (t *tokeniser) GoBack() {
	t.currIndex--
	if t.currIndex < 0 {
		t.currIndex = 0
	}
}

func (t *tokeniser) CheckNumber() bool {
	return t.Peek() >= '0' && t.Peek() <= '9'
}

func (t *tokeniser) Check(ch rune) bool {
	return t.Consume() == ch
}

func (t *tokeniser) ParseMultiply() int {
	if !t.Check('m') || !t.Check('u') || !t.Check('l') || !t.Check('(') || !t.CheckNumber() {
		return -1
	}
	a := t.ParseNumber()

	if !t.Check(',') {
		return -1
	}

	if !t.CheckNumber() {
		return -1
	}
	b := t.ParseNumber()

	if !t.Check(')') {
		return -1
	}

	return a * b
}

const (
	DO = iota
	DONT
	NOT_DO_DONT
)

func (t *tokeniser) ParseDoDont() int {
	if !t.Check('d') || !t.Check('o') {
		return NOT_DO_DONT
	}

	got := t.Consume()

	if got == '(' && t.Check(')') {
		return DO
	}

	if got == 'n' && t.Check('\'') && t.Check('t') && t.Check('(') && t.Check(')') {
		return DONT
	}

	return NOT_DO_DONT
}

func (t *tokeniser) ParseNumber() int {
	snum := ""

	for t.Peek() != 0 && t.CheckNumber() {
		snum += string(t.Consume())
	}

	num, _ := strconv.Atoi(snum)

	return num
}

func part1(prog string) int {
	res := 0

	t := newTokeniser(prog)

	for t.Peek() != 0 { // while not end of file
		mulres := t.ParseMultiply()
		if mulres != -1 {
			res += mulres
			continue
		}

		if t.currIndex != 0 && t.PeekAhead(-1) == 'm' {
			t.GoBack()
		}
	}

	return res
}

func part2(prog string) int {
	res := 0
	enabled := true

	t := newTokeniser(prog)

	for t.Peek() != 0 { // while not end of file
		mulres := t.ParseMultiply()
		if mulres != -1 && enabled {
			res += mulres
			continue
		}

		t.GoBack()

		dodontres := t.ParseDoDont()

		if dodontres == DO {
			enabled = true
		} else if dodontres == DONT {
			enabled = false
		}

		if t.currIndex != 0 && (t.PeekAhead(-1) == 'm' || t.PeekAhead(-1) == 'd') {
			t.GoBack()
		}
	}

	return res
}

func Run(part int, lines []string) int {
	prog := strings.Join(lines, "\n")

	if part == 1 {
		return part1(prog)
	} else {
		return part2(prog)
	}
}
