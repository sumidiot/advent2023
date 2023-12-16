package helpers

import (
	"day16/part1/helpers"
)

func Solve(lines []string) int {
	grid := helpers.ParseLines(lines)
	max := 0
	for row := 0; row < len(grid); row++ {
		var s int
		s = helpers.SolveInput(grid, helpers.Beam{Row: row, Col: 0, Dir: helpers.Right})
		if s > max {
			max = s
		}
		s = helpers.SolveInput(grid, helpers.Beam{Row: row, Col: len(grid[0]) - 1, Dir: helpers.Left})
		if s > max {
			max = s
		}
	}
	for col := 0; col < len(grid[0]); col++ {
		var s int
		s = helpers.SolveInput(grid, helpers.Beam{Row: 0, Col: col, Dir: helpers.Down})
		if s > max {
			max = s
		}
		s = helpers.SolveInput(grid, helpers.Beam{Row: len(grid) - 1, Col: col, Dir: helpers.Up})
		if s > max {
			max = s
		}
	}
	return max
}
