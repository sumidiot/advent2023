package helpers

import (
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	ret := 0
	for _, hist := range i.Histories {
		ret += SolveHistory(hist)
	}
	return ret
}

func SolveHistory(hist History) int {
	c, isc := CheckConst(hist)
	if isc {
		return c
	} else {
		d := Diffs(hist)
		nd := SolveHistory(d)
		return hist[len(hist)-1] + nd
	}
}

func Diffs(hist History) History {
	d := make(History, len(hist)-1)
	for i := 1; i < len(hist); i++ {
		d[i-1] = hist[i] - hist[i-1]
	}
	return d
}

func CheckConst(hist History) (int, bool) {
	// CoPilot got the initial case here but then wanted to just return 0, on the first try
	// when I added the else, it helped out for that block
	if len(hist) == 1 {
		return hist[0], true
	} else {
		c := hist[0]
		for _, n := range hist {
			if n != c {
				return 0, false
			}
		}
		return c, true
	}
}

type History []int

type Inputs struct {
	Histories []History
}

func ParseLines(lines []string) *Inputs {
	hists := make([]History, len(lines))
	for i, line := range lines {
		hists[i] = ParseLine(line)
	}
	return &Inputs{Histories: hists}
}

func ParseLine(line string) History {
	fields := strings.Fields(line)
	hist := make(History, len(fields))
	for i, field := range fields {
		n, _ := strconv.Atoi(field)
		hist[i] = n
	}
	return hist
}