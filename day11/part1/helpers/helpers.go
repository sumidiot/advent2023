package helpers

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *BaseInputs) int {
	return SolveEnriched(Enrich(i))
}

func SolveEnriched(i *RichInputs) int {
	sum := 0
	for lidx := 0; lidx < len(i.Locs)-1; lidx++ {
		for ridx := lidx + 1; ridx < len(i.Locs); ridx++ {
			sum += i.PathLength(lidx, ridx)
		}
	}
	return sum
}

func minMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (i *RichInputs) PathLength(lidx int, ridx int) int {
	minR, maxR := minMax(i.Locs[lidx].Row, i.Locs[ridx].Row)
	minC, maxC := minMax(i.Locs[lidx].Col, i.Locs[ridx].Col)
	rows := 0
	for row := minR; row < maxR; row++ {
		if i.RowIsEmpty[row] {
			rows += 2
		} else {
			rows++
		}
	}
	cols := 0
	for col := minC; col < maxC; col++ {
		if i.ColIsEmpty[col] {
			cols += 2
		} else {
			cols++
		}
	}
	return rows + cols
}

type BaseInputs struct {
	Grid [][]bool
}

type Coord struct {
	Row int
	Col int
}

type RichInputs struct {
	Base       *BaseInputs
	Locs       []Coord
	RowIsEmpty map[int]bool
	ColIsEmpty map[int]bool
}

func ParseLines(lines []string) *BaseInputs {
	grid := make([][]bool, len(lines))
	for row, line := range lines {
		grid[row] = make([]bool, len(line))
		for col, char := range line {
			grid[row][col] = char == '#'
		}
	}
	return &BaseInputs{grid}
}

// CoPilot wrote this, I renamed the 'filled' variable from 'char'
func Enrich(base *BaseInputs) *RichInputs {
	locations := make([]Coord, 0)
	rowIsEmpty := make(map[int]bool)
	colIsEmpty := make(map[int]bool)
	for row, line := range base.Grid {
		for col, filled := range line {
			if filled {
				locations = append(locations, Coord{row, col})
				rowIsEmpty[row] = false
				colIsEmpty[col] = false
			}
		}
	}
	for row := 0; row < len(base.Grid); row++ {
		if _, ok := rowIsEmpty[row]; !ok {
			rowIsEmpty[row] = true
		}
	}
	for col := 0; col < len(base.Grid[0]); col++ {
		if _, ok := colIsEmpty[col]; !ok {
			colIsEmpty[col] = true
		}
	}
	return &RichInputs{base, locations, rowIsEmpty, colIsEmpty}
}
