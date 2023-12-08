package helpers

import "strings"

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	instIdx := 0
	numSteps := 0
	curNode := "AAA"
	for curNode != "ZZZ" {
		// Thanks CoPilot!
		curNode = AdvanceNode(i, instIdx, curNode)
		numSteps++
		instIdx = (instIdx + 1) % len(i.Instructions)
	}
	return numSteps
}

func AdvanceNode(i *Inputs, instIdx int, curNode string) string {
	inst := i.Instructions[instIdx]
	if inst == Left {
		curNode = i.Nodes[curNode].Left
	} else if inst == Right {
		curNode = i.Nodes[curNode].Right
	}
	return curNode
}

type Instruction int

const (
	Left Instruction = iota
	Right
)

// Thanks CoPilot!
func InstructionFromString(s string) Instruction {
	switch s {
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic("Invalid instruction: " + s)
	}
}

type Node struct {
	Id    string
	Left  string
	Right string
}

type Inputs struct {
	Instructions []Instruction
	Nodes        map[string]*Node
}

func ParseLines(lines []string) *Inputs {
	inss := ParseInstructions(lines[0])
	nodes := make(map[string]*Node)
	for _, line := range lines[2:] {
		node := ParseNode(line)
		nodes[node.Id] = node
	}
	return &Inputs{inss, nodes}
}

// Thanks CoPilot!
func ParseInstructions(line string) []Instruction {
	inss := make([]Instruction, len(line))
	for idx, char := range line {
		inss[idx] = InstructionFromString(string(char))
	}
	return inss
}

func ParseNode(line string) *Node {
	parts := strings.Fields(line)
	left := parts[2]
	right := parts[3]
	return &Node{parts[0], left[1:(len(left)-1)], right[0:(len(right)-1)]}
}
