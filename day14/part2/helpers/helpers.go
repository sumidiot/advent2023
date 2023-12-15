package helpers

import (
	"day14/part1/helpers"
	"log"
)

func Solve(lines []string) int {
	return SolveInputs(helpers.ParseLines(lines))
}

type Recurrence struct {
	Begin  int
	Values []int
}

func SolveInputs(grid *helpers.Inputs) int {
	burnInBoundary := 100000
	endTime := 1000000000
	grid = BurnIn(grid, burnInBoundary)
	return SolveRecurrence(grid, burnInBoundary, endTime)
}

func SolveRecurrence(grid *helpers.Inputs, burnInBoundary int, endTime int) int {
	lastSeen := make(map[int]int) // from load => idx
	recurrences := make([]*Recurrence, 0)
	minRecLength := 20
	winRecIdx := -1
	for idx := burnInBoundary; idx < endTime && winRecIdx == -1; idx++ {
		if idx%20000 == 0 {
			log.Printf("Looking for recurrences, through %v, with %v options", idx, len(recurrences))
		}
		grid = Spin(grid)
		load := Load(grid)
		prevIdx, ok := lastSeen[load]
		if !ok {
			lastSeen[load] = idx
			recurrences = make([]*Recurrence, 0) // should try to re-use one object, maybe?
			recurrences = append(recurrences, &Recurrence{
				idx,
				[]int{load},
			})
		} else {
			for wri, rec := range recurrences {
				if rec.Begin == prevIdx && idx-prevIdx >= minRecLength {
					winRecIdx = wri
					log.Printf("Found winning recurrence ending at %v where value is %v, reccurrence length is %v", idx, load, len(recurrences[winRecIdx].Values))
					break
				} else {
					rec.Values = append(rec.Values, load)
				}
			}
		}
	}
	winRec := recurrences[winRecIdx]               // suppose endTime is the end of the recurrence interval (=begin)
	tdiff := endTime - winRec.Begin - 1            // suppose 3 values, 0, 1, 2, and so endTime = 4, beginTime = 0
	return winRec.Values[tdiff%len(winRec.Values)] // so tdiff = endTime - beginTime - 1
}

func Load(grid *helpers.Inputs) int {
	sum := 0
	for ridx, row := range *grid {
		for _, cell := range row {
			if cell == helpers.Rock {
				sum += (len((*grid)) - ridx)
			}
		}
	}
	return sum
}

func BurnIn(grid *helpers.Inputs, numToDo int) *helpers.Inputs {
	for idx := 0; idx < numToDo; idx++ {
		grid = Spin(grid)
	}
	return grid
}

func Spin(grid *helpers.Inputs) *helpers.Inputs {
	return RollEast(RollSouth(RollWest(RollNorth(grid))))
}

func RollNorth(grid *helpers.Inputs) *helpers.Inputs {
	for row := 1; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			m := MovesRow(grid, row, col, -1)
			if m != 0 {
				(*grid)[row][col] = helpers.Open
				(*grid)[row+m][col] = helpers.Rock
			}
		}
	}
	return grid
}

func RollSouth(grid *helpers.Inputs) *helpers.Inputs {
	for row := len(*grid) - 1; row >= 0; row-- {
		for col := 0; col < len((*grid)[row]); col++ {
			m := MovesRow(grid, row, col, 1)
			if m != 0 {
				(*grid)[row][col] = helpers.Open
				(*grid)[row+m][col] = helpers.Rock
			}
		}
	}
	return grid
}

func MovesRow(grid *helpers.Inputs, row int, col int, rowDelta int) int {
	if (*grid)[row][col] != helpers.Rock {
		return 0
	}
	trow := row + rowDelta
	for ; trow >= 0 && trow < len(*grid); trow += rowDelta {
		if (*grid)[trow][col] != helpers.Open {
			break
		}
	}
	return trow - row - rowDelta
}

func RollWest(grid *helpers.Inputs) *helpers.Inputs {
	for col := 1; col < len((*grid)[0]); col++ {
		for row := 0; row < len(*grid); row++ {
			m := MovesCol(grid, row, col, -1)
			if m != 0 {
				(*grid)[row][col] = helpers.Open
				(*grid)[row][col+m] = helpers.Rock
			}
		}
	}
	return grid
}

func RollEast(grid *helpers.Inputs) *helpers.Inputs {
	for col := len((*grid)[0]) - 1; col >= 0; col-- {
		for row := 0; row < len(*grid); row++ {
			m := MovesCol(grid, row, col, 1)
			if m != 0 {
				(*grid)[row][col] = helpers.Open
				(*grid)[row][col+m] = helpers.Rock
			}
		}
	}
	return grid
}

func MovesCol(grid *helpers.Inputs, row int, col int, colDelta int) int {
	if (*grid)[row][col] != helpers.Rock {
		return 0
	}
	tcol := col + colDelta
	for ; tcol >= 0 && tcol < len((*grid)[row]); tcol += colDelta {
		if (*grid)[row][tcol] != helpers.Open {
			break
		}
	}
	return tcol - col - colDelta
}
