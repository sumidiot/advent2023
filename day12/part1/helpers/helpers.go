package helpers

import (
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	sum := 0
	for _, row := range *i {
		sum += SolveRow(row)
	}
	return sum
}

func SolveRow(r Row) int {
	return 0
}

type SpringState int

const (
	Operational SpringState = iota
	Damaged
	Unknown
)

type Row struct {
	States  []SpringState
	Contigs []int
}

type Inputs []Row

func ParseLines(lines []string) *Inputs {
	ret := make(Inputs, len(lines))
	for i, line := range lines {
		ret[i] = ParseLine(line)
	}
	return &ret
}

func ParseLine(line string) Row {
	ret := Row{}
	fields := strings.Fields(line) // two of them
	for _, c := range fields[0] {
		switch c {
		case '.':
			ret.States = append(ret.States, Operational)
		case '#':
			ret.States = append(ret.States, Damaged)
		case '?':
			ret.States = append(ret.States, Unknown)
		}
	}
	contigs := strings.Split(fields[1], ",")
	for _, c := range contigs {
		n, _ := strconv.Atoi(c)
		ret.Contigs = append(ret.Contigs, n)
	}
	return ret
}
