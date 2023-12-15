package helpers

import (
	"day12/part1/helpers"
	"fmt"
)

func Solve(lines []string) int {
	return SolveInput(helpers.ParseLines(lines))
}

func SolveInput(i *helpers.Inputs) int {
	sum := 0
	for idx, row := range *i {
		if idx%20 == 0 {
			fmt.Println(idx)
		}
		sum += SolveRow(ExpandRow(row, 5))
	}
	return sum
}

func SolveRow(r helpers.Row) int {
	return 0
}

func ExpandRow(r helpers.Row, num int) helpers.Row {
	retStates := make([]helpers.SpringState, len(r.States) * num)
	retContig := make([]int, len(r.Contigs) * num)
	for idx, state := range r.States {
		for i := 0; i < num; i++ {
			retStates[i * len(r.States) + idx] = state
		}
	}
	for idx, contig := range r.Contigs {
		for i := 0; i < num; i++ {
			retContig[i * len(r.Contigs) + idx] = contig
		}
	}
	return helpers.Row{States: retStates, Contigs: retContig}
}

func SplitRow(states []helpers.SpringState) [][]helpers.SpringState {
	ret := make([][]helpers.SpringState, 0)
	cur := make([]helpers.SpringState, 0)
	for _, state := range states {
		switch state {
		case helpers.Operational:
			if len(cur) > 0 {
				ret = append(ret, cur)
				cur = make([]helpers.SpringState, 0)
			}
		default:
			cur = append(cur, state)
		}
	}
	if len(cur) > 0 {
		ret = append(ret, cur)
	}
	return ret
}

type SplitSolution struct {
	TakenContigs int
	NumSolutions int
}

/**
 * The point of a "split" here is that all of the states are either ? or . # (damaged).
 * Also that the neighbors are not Damaged or Unknown.
 * We identify which combinations of number of contigs we can take for this split,
 * and for each choice, how many solutions that would give.
 *
 * We assume states is non-empty, but contigs may be empty.
 *
 * We return an empty array if there are no solutions.
 */
func SolutionsInSplit(states []helpers.SpringState, contigs []int) []SplitSolution {
	// taking cumsum of contigs, we in theory could probably take any number of contigs
	// up until the cumsum is greater than len(states).
	// actually you have to add an extra one in between to account for making the gap
	ret := make([]SplitSolution, 0)
	cumsum := 0
	taking := 0
	for ; cumsum < len(states) && taking <= len(contigs); taking++ {
		solutions := BruteForceSolveSimplifiedRow(states, contigs[0:taking])
		ret = append(ret, SplitSolution{
			TakenContigs: taking,
			NumSolutions: solutions,
		})
	}
	return ret
}

/**
 * The goal here is to explicitly use all the contigs
 */
func BruteForceSolveSimplifiedRow(states []helpers.SpringState, contigs []int) int {
	// if no contigs
	if len(contigs) == 0 {
		// the only valid solution is if all states are Unknown
		for _, state := range states {
			if state == helpers.Damaged {
				return 0
			}
		}
		return 1
	}
	// since we have ? and #, we can only make as many gaps as we have ?
	// suppose we have 0 ?, then we can only make 1 contig
	// if we have one ?, then we can only two contigs
	numUnknown := 0
	for _, state := range states {
		if state == helpers.Unknown {
			numUnknown++
		}
	}
	if numUnknown + 1 < len(contigs) {
		return 0
	}
	// we have enough ? to make all the contigs

	// do we have enough states?
	numRequiredStates := contigs[0]
	for _, contig := range contigs[1:] {
		numRequiredStates += contig + 1
	}
	if numRequiredStates > len(states) {
		return 0
	}

	// ok, so we have enough ? to make the number of contigs,
	// and enough states to make the contigs and gaps

	// how many ways are there to do it?
	// one way to think about it:
	//    since there are no Operational
	//    and we are trying to make len(contigs) blocks

	minLenIfTaking := make([]int, len(contigs) + 1)
	minLenIfTaking[0] = 0
	for idx := 0; idx < len(contigs); idx++ {
		minLenIfTaking[idx + 1] = minLenIfTaking[idx] + contigs[idx] + 1
		if idx == 0 {
			minLenIfTaking[1] = minLenIfTaking[1] - 1 // don't need initial gap
		}
	}
	if minLenIfTaking[len(minLenIfTaking) - 1] >= len(states) {
		return 0
	}
	return 0
}