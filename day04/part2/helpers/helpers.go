package helpers

import (
	"regexp"
	"strconv"
)

func Solve(lines []string) int {
	ret := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == '*' {
				parts := allAdjacentParts(row, col, lines)
				if len(parts) == 2 {
					ret += (parts[0] * parts[1])
				}
			}
		}
	}
	return ret
}

/**
 * return integers adjacent to location row/col in lines
 */
func allAdjacentParts(row int, col int, lines []string) []int {
	ret := make([]int, 0)
	partLeft, isLeft := findPartLeft(col, lines[row])
	if isLeft {
		ret = append(ret, partLeft)
	}
	partRight, isRight := findPartRight(col, lines[row])
	if isRight {
		ret = append(ret, partRight)
	}
	if row > 0 {
		partsUp := findPartsAdjacent(col, lines[row-1])
		ret = append(ret, partsUp...)
	}
	if row < len(lines)-1 {
		partsDown := findPartsAdjacent(col, lines[row+1])
		ret = append(ret, partsDown...)
	}
	return ret
}

func findPartsAdjacent(col int, line string) []int {
	lidx := col - 1
	if lidx < 0 {
		lidx = 0
	}
	for ; lidx >= 0; lidx-- {
		reg, _ := regexp.MatchString("[0-9]", line[lidx:(lidx+1)])
		if !reg {
			break
		}
	}
	if lidx < 0 {
		lidx = 0
	}
	ridx := col + 1
	if ridx > len(line) {
		ridx = len(line)
	}
	for ; ridx < len(line); ridx++ {
		reg, _ := regexp.MatchString("[0-9]", line[ridx:(ridx+1)])
		if !reg {
			break
		}
	}
	reg := regexp.MustCompile("[0-9]+")
	parts := reg.FindAllString(line[lidx:ridx], 2) // shouldn't need to say 2, but just in case :)
	ret := make([]int, 0)
	for _, part := range parts {
		p, _ := strconv.Atoi(part)
		ret = append(ret, p)
	}
	return ret
}

func findPartLeft(col int, line string) (int, bool) {
	str := ""
	for i := col - 1; i >= 0; i-- {
		reg, _ := regexp.MatchString("[0-9]", line[i:(i+1)])
		if !reg {
			break
		}
		str = line[i:(i+1)] + str
	}
	if len(str) > 0 {
		ret, _ := strconv.Atoi(str)
		return ret, true
	} else {
		return 0, false
	}
}

func findPartRight(col int, line string) (int, bool) {
	num, _, isNum := startsAsNum(line[col+1:])
	return num, isNum
}

// copied from part1
func startsAsNum(line string) (int, int, bool) {
	reg := regexp.MustCompile("^[0-9]+")
	if reg.MatchString(line) {
		str := reg.FindString(line)
		part, _ := strconv.Atoi(str)
		return part, len(str), true
	} else {
		return 0, 0, false
	}
}
