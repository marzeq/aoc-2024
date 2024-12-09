package day9

import "fmt"

type SegmentType int

const (
	FILE SegmentType = iota
	FREE
)

type Segment struct {
	stype    SegmentType
	id       int // unused if stype == FREE
	size     int
	unusable bool
	next     *Segment
	prev     *Segment
}

func (s Segment) String() string {
	st := ""

	if s.next == nil {
		st += "last"
	}

	if s.prev == nil {
		st += "first"
	}

	if s.stype == FREE {
		return fmt.Sprintf("FreeSpace(%d%s)", s.size, st)
	} else {
		return fmt.Sprintf("File<%d>(%d%s)", s.id, s.size, st)
	}
}

func newFile(id int, size int) Segment {
	return Segment{
		stype: FILE,
		id:    id,
		size:  size,
	}
}

func newFreeSpace(size int) Segment {
	return Segment{
		stype: FREE,
		size:  size,
	}
}

func parseInput(lines []string) (*Segment, *Segment) {
	l := lines[0]
	var first *Segment
	var last *Segment

	for i, ch := range l {
		n := int(ch - 48)

		if n == 0 {
			continue
		}

		var sp Segment
		if i%2 == 0 {
			sp = newFile(i/2, n)
		} else {
			sp = newFreeSpace(n)
		}

		if i == 0 {
			first = &sp
			last = &sp
		} else {
			last.next = &sp
			sp.prev = last
			last = &sp
		}
	}

	return first, last
}

func printLinkedList(first *Segment) {
	for curr := first; curr != nil; curr = curr.next {
		var ch rune

		if curr.unusable {
			ch = '-'
		} else if curr.stype == FREE {
			ch = '.'
		} else {
			ch = rune(curr.id + 48)
		}

		for i := 0; i < curr.size; i++ {
			fmt.Print(string(ch))
		}
		fmt.Print(" ")
	}

	fmt.Println()
}

func findFirstFreeSpace(first, last *Segment, atLeastSize int) *Segment {
	if atLeastSize == 0 && first.stype == FREE && !first.unusable {
		return first
	}

	for curr := first; curr != nil && curr != last; curr = curr.next {
		if curr.stype == FREE && curr.size >= atLeastSize && !curr.unusable {
			return curr
		}
	}
	return nil
}

func findFileById(last *Segment, id int) *Segment { // id = -1 means we dont care about id
	if (id == -1 || last.id == id) && last.stype == FILE {
		return last
	}

	for curr := last; curr != nil; curr = curr.prev {
		if curr.stype == FILE && (id == -1 || curr.id == id) {
			return curr
		}
	}

	return nil
}

func removeLast(last *Segment) *Segment {
	last.prev.next = nil

	newlast := last.prev

	return newlast
}

func insertBefore(curr, seg *Segment) {
	newSegment := &Segment{
		stype: seg.stype,
		size:  seg.size,
		id:    seg.id,
	}

	beforeCurr := curr.prev

	if beforeCurr != nil {
		beforeCurr.next = newSegment
	}

	curr.prev = newSegment
	newSegment.next = curr
	newSegment.prev = beforeCurr
}

func calculateChecksum(first *Segment) int {
	checksum := 0
	position := 0
	for curr := first; curr != nil; curr = curr.next {
		if curr.stype == FILE {
			for i := 0; i < curr.size; i++ {
				checksum += curr.id * position
				position++
			}
		} else {
			position += curr.size
		}
	}
	return checksum
}

func part1(first, last *Segment) int {
	for {
		frspace := findFirstFreeSpace(first, last, 0)

		if frspace == nil {
			break
		}

		for {

			file := findFileById(last, -1)
			if file.prev == nil && file.next == nil {
				break
			}
			sizediff := frspace.size - file.size

			if sizediff == 0 {
				frspace.stype = FILE
				frspace.id = file.id

				last = removeLast(last)
				if last.stype == FREE {
					last = removeLast(last)
				}

				break
			} else if sizediff < 0 {
				frspace.stype = FILE
				frspace.id = file.id

				file.size = (-sizediff)
				break
			} else {
				nfile := newFile(file.id, file.size)
				insertBefore(frspace, &nfile)

				last = removeLast(last)
				if last.stype == FREE {
					last = removeLast(last)
				}

				frspace.size = sizediff
				continue
			}
		}
	}

	if last.id == last.prev.id {
		last.prev.size += last.size
		last = removeLast(last)
	}
	return calculateChecksum(first)
}

func part2(first, last *Segment) int {
	for id := last.id; id >= 0; id-- {
		file := findFileById(last, id)

		if file == nil {
			panic("bruh")
		}

		fspace := findFirstFreeSpace(first, file, file.size)

		if fspace == nil {
			continue
		}

		fspace.size -= file.size

		insertBefore(fspace, file)

		file.stype = FREE
		file.unusable = true

		if fspace.size == 0 {
			first, last = removeSegment(first, last, fspace)
		}

	}

	return calculateChecksum(first)
}

func removeSegment(first, last, seg *Segment) (*Segment, *Segment) {
	if last == seg {
		return first, removeLast(last)
	}

	if first == seg {
		return removeFirst(first), last
	}

	next := seg.next
	prev := seg.prev

	prev.next = next
	next.prev = prev

	seg = nil

	return first, last
}

func removeFirst(first *Segment) *Segment {
	first.next.prev = nil

	newfirst := first.next

	return newfirst
}

func Run(part int, lines []string) int {
	first, last := parseInput(lines)

	if part == 1 {
		return part1(first, last)
	} else {
		return part2(first, last)
	}
}
