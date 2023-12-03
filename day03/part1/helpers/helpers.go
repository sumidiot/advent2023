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
			part, colEnd, empty := numAt(row, col, lines)
			if !empty {
				log.Printf("r: %d, c: %d, e: %d = v: %d", row, col, colEnd, part)
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

func numAt(row, col int, lines []string) (int, int, bool) {
	res, _ := regexp.MatchString("[0-9]", lines[row][col:(col+1)])
	if res {
		// substring of lines[row] from col to next non-digit
		colEnd := col
		for res && colEnd < len(lines[row]) {
			res, _ = regexp.MatchString("[0-9]", lines[row][colEnd:(colEnd+1)])
			if res {
				colEnd++
			}
		}
		part, _ := strconv.Atoi(lines[row][col:colEnd])
		return part, colEnd, false
	}
	return 0, 0, true
}

func touches(row, col, colEnd int, lines []string) bool {
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= colEnd+1; c++ {
			if r >= 0 && r < len(lines) && c >= 0 && c < len(lines[r]) {
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
