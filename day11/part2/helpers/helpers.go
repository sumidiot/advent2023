package helpers

import (
	"day11/part1/helpers"
)

func Solve(lines []string, scale int) int {
	n, s := SolveInput(helpers.ParseLines(lines), scale)
	return n + s * scale
}

func SolveInput(i *helpers.BaseInputs, scale int) (int, int) {
	norm, scaled := SolveEnriched(helpers.Enrich(i, scale))
	return norm, scaled
}

func SolveEnriched(i *helpers.RichInputs) (int, int) {
	norm, scaled := 0, 0
	for lidx := 0; lidx < len(i.Locs)-1; lidx++ {
		for ridx := lidx + 1; ridx < len(i.Locs); ridx++ {
			n, s := TypedPathLength(i, lidx, ridx)
			norm += n
			scaled += s
		}
	}
	return norm, scaled
}

func TypedPathLength(i *helpers.RichInputs, lidx int, ridx int) (int, int) {
	minR, maxR := helpers.MinMax(i.Locs[lidx].Row, i.Locs[ridx].Row)
	minC, maxC := helpers.MinMax(i.Locs[lidx].Col, i.Locs[ridx].Col)
	norm, scaled := 0, 0
	for row := minR; row < maxR; row++ {
		if i.RowIsEmpty[row] {
			scaled++
		} else {
			norm++
		}
	}
	for col := minC; col < maxC; col++ {
		if i.ColIsEmpty[col] {
			scaled++
		} else {
			norm++
		}
	}
	return norm, scaled
}