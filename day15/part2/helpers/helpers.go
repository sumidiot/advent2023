package helpers

import (
	"strconv"
	"strings"
	"day15/part1/helpers"
)


func Solve(lines []string) int {
	return SolveInstructions(ParseLines(lines))
}

func SolveInstructions(instructions []Instruction) int {
	boxes := MakeBoxes(instructions)
	return SumLengths(boxes)
}

func MakeBoxes(instructions []Instruction) [][]MakeLength {
	boxes := make([][]MakeLength, 256)
	for _, ins := range instructions {
		code := ins.GetCode()
		boxIdx := helpers.Hash(ins.GetCode())
		switch ins.(type) {
		case Remove:
			delIdx := -1
			for i := 0; i < len(boxes[boxIdx]); i++ {
				if (boxes[boxIdx])[i].Code == code {
					delIdx = i
					break
				}
			}
			if delIdx >= 0 {
				boxes[boxIdx] = append((boxes[boxIdx])[:delIdx], (boxes[boxIdx])[delIdx + 1:]...)
			}
		case MakeLength:
			ml, _ := ins.(MakeLength)
			replaceIdx := -1
			for i := 0; i < len(boxes[boxIdx]); i++ {
				if (boxes[boxIdx])[i].Code == code {
					replaceIdx = i
					break
				}
			}
			if replaceIdx >= 0 {
				boxes[boxIdx][replaceIdx] = ml
			} else {
				boxes[boxIdx] = append(boxes[boxIdx], ml)
			}
		}
	
	}
	return boxes
}

func SumLengths(boxes [][]MakeLength) int {
	sum := 0
	for bidx, box := range boxes {
		for lidx, ml := range box {
			sum += (1 + bidx) * (1 + lidx) * ml.Length
		}
	}
	return sum
}

type Instruction interface {
	GetCode() string
}

type Remove struct {
	Code string
}

func (r Remove) GetCode() string {
	return r.Code
}

type MakeLength struct {
	Code string
	Length int
}

func (m MakeLength) GetCode() string {
	return m.Code
}

func ParseLines(lines []string) []Instruction {
	fields := strings.Split(lines[0], ",")
	instructions := make([]Instruction, len(fields))
	for idx, field := range fields {
		if strings.Contains(field, "-") {
			instructions[idx] = Remove{ field[0:(len(field) - 1)] }
		} else {
			eqFields := strings.Split(field, "=")
			l, _ := strconv.Atoi(eqFields[1])
			instructions[idx] = MakeLength{ eqFields[0], l }
		}
	}
	return instructions
}