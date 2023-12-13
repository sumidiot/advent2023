package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	sum := 0
	for idx, row := range *i {
		if idx%20 == 0 {
			fmt.Println(idx)
		}
		sum += SolveRow(row)
	}
	return sum
}

func SolveRow(r Row) int {
	return SolveRowBruteForce(r, 0, 0, make([]SpringState, r.NumUnknown()))
}

func SolveRowBruteForce(r Row, fromIdx int, subIdx int, subs []SpringState) int {
	if fromIdx >= len(r.States) {
		if ValidRow(r, subs) {
			// fmt.Println(r, subs)
			return 1
		}
		return 0
	} else if r.States[fromIdx] != Unknown {
		return SolveRowBruteForce(r, fromIdx+1, subIdx, subs)
	} else {
		sum := 0
		for _, state := range []SpringState{Operational, Damaged} {
			subs[subIdx] = state
			sum += SolveRowBruteForce(r, fromIdx+1, subIdx+1, subs)
		}
		return sum
	}
}

func ValidRow(r Row, subs []SpringState) bool {
	curContig := 0
	inBroken := false
	contigIdx := -1
	subIdx := 0
	for _, state := range r.States {
		if state == Operational || (state == Unknown && subs[subIdx] == Operational) {
			if inBroken {
				if contigIdx >= len(r.Contigs) || r.Contigs[contigIdx] != curContig {
					return false
				}
				inBroken = false
				curContig = 0
			}
			if state == Unknown {
				subIdx++
			}
		} else if state == Damaged || (state == Unknown && subs[subIdx] == Damaged) {
			if !inBroken {
				contigIdx++
				if contigIdx >= len(r.Contigs) {
					return false
				}
				curContig = 0
				inBroken = true
			}
			curContig++
			if state == Unknown {
				subIdx++
			}
		} else {
			panic("WTH")
		}
	}
	if inBroken && r.Contigs[contigIdx] != curContig {
		return false
	}
	return contigIdx == len(r.Contigs)-1
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

func (r *Row) NumUnknown() int {
	sum := 0
	for _, state := range r.States {
		if state == Unknown {
			sum++
		}
	}
	return sum
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
