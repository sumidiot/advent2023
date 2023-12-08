package helpers

import (
	"day08/part1/helpers"
	"fmt"
	"math"
)

func Solve(lines []string) int {
	return LCMSolve(helpers.ParseLines(lines))
}

/**
 * We have (740-3+1) = 738 nodes
 * Of those, 6 are start nodes (end with 'A')
 * Our instruction length is 281
 * The LCM of nodes and instructions is 207378 (=product, relatively prime)
 * If we find the times that each node ends at a Z, we can find the joint LCM combination
 */
func LCMSolve(i *helpers.Inputs) int {
	sns := FindStartNodes(i)
	periods := make([][]int, 0)
	for _, sn := range sns {
		periods = append(periods, FindPeriods(i, sn))
	}
	return SolvePeriods(periods, -1, 1)
}

func minMaybe(a, b int) int {
	if (a == -1) {
		return b
	} else if a < b {
		return a
	}
	return b
}

func SolvePeriods(periods [][]int, bestKnown int, accum int) int {
	if len(periods) == 1 {
		return minMaybe(bestKnown, lcm(accum, periods[0][0]))
	} else {
		for _, p := range(periods[0]) {
			// CoPilot guess here was reasonable but not quite right
			bestKnown = minMaybe(bestKnown, SolvePeriods(periods[1:], bestKnown, lcm(accum, p)))
		}
		return bestKnown
	}
}

func FindPeriods(i *helpers.Inputs, startNode string) []int {
	periods := make([]int, 0)
	curNode := startNode
	stopIdx := lcm(len(i.Instructions), len(i.Nodes))
	instIdx := 0
	numSteps := 0
	for numSteps < stopIdx {
		curNode = helpers.AdvanceNode(i, instIdx, curNode)
		numSteps++
		if curNode[len(curNode)-1:] == "Z" {
			// observation is that the periods for a given node all work out to be multiples of a base period
			// this is a bit of an assumption, it should a + bn for some a, b more generally,
			// and in fact it could be that there are several, a_i + b_i n,
			// but i guess we get lucky. having too many periods leaves more work for later, but shouldn't lead to a wrong answer
			if len(periods) == 0 || numSteps % periods[len(periods)-1] != 0 {
				periods = append(periods, numSteps)
			}
		}
		instIdx = (instIdx + 1) % len(i.Instructions)
	}
	return periods
}

func BruteForceSolveInput(i *helpers.Inputs) int {
	instIdx := 0
	curNodes := FindStartNodes(i)
	numSteps := 0
	for !IsEnd(curNodes) {
		// Thanks CoPilot!
		inst := i.Instructions[instIdx]
		numSteps++
		if inst == helpers.Left {
			curNodes = MoveLeft(i, curNodes)
		} else if inst == helpers.Right {
			curNodes = MoveRight(i, curNodes)
		}
		instIdx = (instIdx + 1) % len(i.Instructions)
		if numSteps > len(i.Instructions)*len(i.Nodes)*len(curNodes) {
			panic("Aborting")
		}
		if numSteps%500000 == 0 {
			fmt.Println(numSteps)
		}
	}
	return numSteps
}

func FindStartNodes(i *helpers.Inputs) []string {
	startNodes := make([]string, 0)
	for id, _ := range i.Nodes {
		if id[len(id)-1:] == "A" {
			startNodes = append(startNodes, id)
		}
	}
	return startNodes
}

func IsEnd(curNodes []string) bool {
	for _, node := range curNodes {
		if node[len(node)-1:] != "Z" {
			return false
		}
	}
	return true
}

// Thanks CoPilot!
func MoveLeft(i *helpers.Inputs, curNodes []string) []string {
	newNodes := make([]string, 0)
	for _, node := range curNodes {
		newNodes = append(newNodes, i.Nodes[node].Left)
	}
	return newNodes
}

// Thanks CoPilot!
func MoveRight(i *helpers.Inputs, curNodes []string) []string {
	newNodes := make([]string, 0)
	for _, node := range curNodes {
		newNodes = append(newNodes, i.Nodes[node].Right)
	}
	return newNodes
}

// Thanks CoPilot!
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Thanks CoPilot!
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
