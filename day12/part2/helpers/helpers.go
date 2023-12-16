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
	retStates := make([]helpers.SpringState, len(r.States)*num)
	retContig := make([]int, len(r.Contigs)*num)
	for idx, state := range r.States {
		for i := 0; i < num; i++ {
			retStates[i*len(r.States)+idx] = state
		}
	}
	for idx, contig := range r.Contigs {
		for i := 0; i < num; i++ {
			retContig[i*len(r.Contigs)+idx] = contig
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
 *
 * Generally the solution here is recursive, finding some solutions and leaving a smaller problem to iterate on
 */
func SolutionsInSplit(states []helpers.SpringState, contigs []int) []SplitSolution {
	ret := make([]SplitSolution, 0)
	if CanTakeNoContigs(states) {
		ret = append(ret, SplitSolution{
			TakenContigs: 0,
			NumSolutions: 1,
		})
	}
	if len(contigs) > 0 {
		contig := contigs[0]
		posFirstDamaged := -1
		for idx, state := range states {
			if state == helpers.Damaged {
				posFirstDamaged = idx
				break
			}
		}

		// if we take a single contig, where within states will we have the first empty, and how many solutions does each possibility give?
		splitIdxSolutions := make(map[int]int)
		if posFirstDamaged == -1 {
			unkSolutions := AllUnknownSolutions(len(states), contig)
			for _, unkSolution := range unkSolutions {
				splitIdxSolutions[unkSolution] = 1
			}
		} else {
			// suppose contig = 2
			// 0123456
			// ???#... can be achieved
			// ##.^ 3 is the first posFirstDamaged that'd work to sneek a solution in first
			if posFirstDamaged >= contig+1 {
				// -1 because we need to leave a gap after the contig
				beforeSolutions := AllUnknownSolutions(posFirstDamaged-1, contig)
				for _, soln := range beforeSolutions {
					prev, ok := splitIdxSolutions[soln]
					if !ok {
						prev = 0
					}
					splitIdxSolutions[soln] = prev + 1
				}
			} else {
				solns := SolutionsIncluding(states, contig, posFirstDamaged)
				for _, soln := range solns {
					prev, ok := splitIdxSolutions[soln]
					if !ok {
						prev = 0
					}
					splitIdxSolutions[soln] = prev + 1
				}
			}
		}

		// reminder, splitIdxSolutions now maps from places we'd put a . to the number of solutions that would give, to take one contig
		// for each solution that takes one contig, and has an endpoint at splitIdx,
		// we can see if we can additionally take more solutions from splitIdx + 1, as long as that's not the end
		// if we can make more solutions from there, we add one to their "takes", and the num solutions is
		// the number of solutions we make in the first part times the number in the second part
		// if we find no more solutions, we just return our own takes
		for splitIdx, numSolns := range splitIdxSolutions {
			// recursively see if we can make more
			// suppose we have 5 states, 01234
			// SolutionsInSplit assumes states is non-empty, so we can only go up to idx 3
			if splitIdx < len(states)-1 {
				moreSolns := SolutionsInSplit(states[:splitIdx], contigs[1:])
				if len(moreSolns) > 0 {
					for _, moreSoln := range moreSolns {
						ret = append(ret, SplitSolution{
							TakenContigs: 1 + moreSoln.TakenContigs,
							NumSolutions: numSolns * moreSoln.NumSolutions,
						})
					}
				} else {
					ret = append(ret, SplitSolution{
						TakenContigs: 1,
						NumSolutions: numSolns,
					})
				}
			} else {
				ret = append(ret, SplitSolution{
					TakenContigs: 1,
					NumSolutions: numSolns,
				})
			}
		}
	}
	return ret
}

func CanTakeNoContigs(states []helpers.SpringState) bool {
	for _, state := range states {
		if state == helpers.Damaged {
			return false
		}
	}
	return true
}

// if we have numStates ?s, and are trying to make a block of size contig,
// return the array of possible endpoints
func AllUnknownSolutions(numStates int, contig int) []int {
	if numStates < contig {
		return []int{}
	}
	// if we have 5 states and are trying to make a contig of size 2:
	// 01234
	// ##...
	// .##..
	// ..##.
	// ...##
	ret := make([]int, numStates-contig+1)
	for idx := contig; idx <= numStates; idx++ {
		ret[idx-contig] = idx
	}
	return ret
}

func SolutionsIncluding(states []helpers.SpringState, contig int, posFirstDamaged int) []int {
	// ???#... suppose contig = 3. posFirstDamaged = 3. first possible start idx is 1
	lidx := posFirstDamaged - contig + 1
	if lidx < 0 {
		lidx = 0
	}

	ret := make([]int, 0)
	// config = 3 => lidx, lidx + 1, lidx + 2 => ridx = lidx + contig - 1
	for ; lidx <= posFirstDamaged; lidx++ {
		ridx := lidx + contig - 1
		if ridx < len(states) || states[ridx] == helpers.Unknown {
			ret = append(ret, ridx)
		}
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
	if numUnknown+1 < len(contigs) {
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
	// and enough states to make the contigs and gaps (in summary anyway, not necessarily each contig block)

	// how many ways are there to do it?
	// one way to think about it:
	//    since there are no Operational
	//    and we are trying to make len(contigs) blocks
	// we have len(states) - numUnknown filled in. we need sum(contigs) filled in
	// so we are going to make sum(contigs) - (len(states) - numUnknown) be empty

	minLenIfTaking := make([]int, len(contigs)+1)
	minLenIfTaking[0] = 0
	for idx := 0; idx < len(contigs); idx++ {
		minLenIfTaking[idx+1] = minLenIfTaking[idx] + contigs[idx] + 1
		if idx == 0 {
			minLenIfTaking[1] = minLenIfTaking[1] - 1 // don't need initial gap
		}
	}
	if minLenIfTaking[len(minLenIfTaking)-1] >= len(states) {
		return 0
	}
	return 0
}
