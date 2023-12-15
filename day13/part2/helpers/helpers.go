package helpers

import (
	"day13/part1/helpers"
	"log"
)

func Solve(lines []string) int {
	return SolveInput(helpers.ParseLines(lines))
}

func SolveInput(i *helpers.Inputs) int {
	sum := 0
	for idx, grid := range *i {
		log.Printf("Solving grid %v", idx)
		sum += SolveGrid(grid)
	}
	return sum
}

func SolveGrid(grid helpers.Grid) int {
	ov, oiv := FindVertical(grid, -1, -1, -1)
	oh, oih := FindHorizontal(grid, -1, -1, -1)
	erow, ecol := -1, -1
	if oiv {
		ecol = ov
	}
	if oih {
		erow = oh
	
	}
	for srow := 0; srow < len(grid); srow++ {
		for scol := 0; scol < len(grid[srow]); scol++ {
			sv, siv := FindVertical(grid, srow, scol, ecol)
			if siv && (!oiv || ov != sv) {
				return sv + 1
			} else {
				sh, sih := FindHorizontal(grid, srow, scol, erow)
				if sih && (!oih || oh != sh) {
					return 100 * (sh + 1)
				}
			}
		}
	}
	panic("No solution")
}

func AgreeAt(grid helpers.Grid, row1, col1, row2, col2, srow, scol int) bool {
	left := grid[row1][col1]
	if (srow == row1 && scol == col1) {
		left = !left
	}
	right := grid[row2][col2]
	if (srow == row2 && scol == col2) {
		right = !right
	}
	return left == right
}

func FindVertical(grid helpers.Grid, srow, scol int, ecol int) (int, bool) {
	for col := 0; col < len(grid[0])-1; col++ {
		if col == ecol {
			continue
		}
		isMirror := true
		for left := col; isMirror && left >= 0 && col+1+(col-left) < len(grid[0]); left-- {
			right := col + 1 + (col - left)
			for row := 0; isMirror && row < len(grid); row++ {
				if !AgreeAt(grid, row, left, row, right, srow, scol) {
					isMirror = false
				}
			}
		}
		if isMirror {
			return col, true
		}
	}
	return -1, false
}

func FindHorizontal(grid helpers.Grid, srow, scol int, erow int) (int, bool) {
	for row := 0; row < len(grid)-1; row++ {
		if row == erow {
			continue
		}
		isMirror := true
		for top := row; isMirror && top >= 0 && row+1+(row-top) < len(grid); top-- {
			bottom := row + 1 + (row - top)
			for col := 0; isMirror && col < len(grid[row]); col++ {
				if !AgreeAt(grid, top, col, bottom, col, srow, scol) {
					isMirror = false
				}
			}
		}
		if isMirror {
			return row, true
		}
	}
	return -1, false
}