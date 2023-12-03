package helpers

import (
	"log"
	"regexp"
	"strconv"
)

func Solve(lines []string) int {
	row := 0
	partSum := 0
	for row < len(lines) {
		col := 0
		for col < len(lines[row]) {
			part, partLen, isNum := numAt(row, col, lines)
			colEnd := col + partLen
			if isNum {
				log.Printf("r: %d, c: %d, e: %d = v: %d", row, col, partLen, part)
				if touches(row, col, colEnd, lines) {
					partSum += part
				}
				col = colEnd
			} else {
				col++
			}
		}
		row++
		log.Printf("Up to row %d", row)
	}
	return partSum
}

/**
 * If line starts with a positive or negative integer, return the integer and the
 * length of the matched string, and true. Otherwise, return 0, 0, false.
 */
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

func numAt(row, col int, lines []string) (int, int, bool) {
	return startsAsNum(lines[row][col:])
}

func touches(row, col, colEnd int, lines []string) bool {
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= colEnd; c++ {
			if r >= 0 && r < len(lines) && c >= 0 && c < len(lines[r]) && (r != row || c < col || c >= colEnd) {
				res, _ := regexp.MatchString("[^.0-9]", lines[r][c:(c+1)])
				if res {
					log.Printf("   Match at %d %d", r, c)
					return true
				}
			}
		}
	}
	return false
}
