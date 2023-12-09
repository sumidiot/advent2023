package helpers

import (
	"day09/part1/helpers"
)

func Solve(lines []string) int {
	return SolveInput(helpers.ParseLines(lines))
}

func SolveInput(i *helpers.Inputs) int {
	ret := 0
	for _, hist := range i.Histories {
		ret += SolveHistory(hist)
	}
	return ret
}

func SolveHistory(hist helpers.History) int {
	c, isc := helpers.CheckConst(hist)
	if isc {
		return c
	} else {
		d := helpers.Diffs(hist)
		nd := SolveHistory(d)
		// hist[0] - ans = nd
		return hist[0] - nd
	}
}
