package helpers

import "log"

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	sum := 0
	for idx, grid := range *i {
		log.Printf("Solving grid %v", idx)
		sum += SolveGrid(grid)
	}
	return sum
}

func SolveGrid(grid Grid) int {
	vert, isVert := FindVertical(grid)
	if isVert {
		return vert + 1
	} else {
		horiz, isHoriz := FindHorizontal(grid)
		if !isHoriz {
			panic("No reflection")
		}
		return 100 * (horiz + 1)
	}
}

func FindVertical(grid Grid) (int, bool) {
	for col := 0; col < len(grid[0])-1; col++ {
		isMirror := true
		for left := col; isMirror && left >= 0 && col+1+(col-left) < len(grid[0]); left-- {
			right := col + 1 + (col - left)
			for row := 0; isMirror && row < len(grid); row++ {
				if grid[row][left] != grid[row][right] {
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

func FindHorizontal(grid Grid) (int, bool) {
	for row := 0; row < len(grid)-1; row++ {
		isMirror := true
		for top := row; isMirror && top >= 0 && row+1+(row-top) < len(grid); top-- {
			bottom := row + 1 + (row - top)
			for col := 0; isMirror && col < len(grid[row]); col++ {
				if grid[top][col] != grid[bottom][col] {
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

type Grid [][]bool

type Inputs []Grid

func ParseLines(lines []string) *Inputs {
	grids := make(Inputs, 0)
	grid := make(Grid, 0)
	for _, line := range lines {
		if len(line) == 0 {
			grids = append(grids, grid)
			grid = make(Grid, 0)
		} else {
			row := make([]bool, 0)
			for _, char := range line {
				row = append(row, char == '#')
			}
			grid = append(grid, row)
		}
	}
	grids = append(grids, grid)
	return &grids
}
