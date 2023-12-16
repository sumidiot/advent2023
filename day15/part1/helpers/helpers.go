package helpers

import "strings"

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	sum := 0
	for _, s := range i.Seqs {
		sum += Hash(s)
	}
	return sum
}

func Hash(str string) int {
	cur := 0
	for _, c := range str {
		cur += int(c)
		cur *= 17
		cur = cur % 256
	}
	return cur
}

type Inputs struct {
	Seqs []string
}

func ParseLines(lines []string) *Inputs {
	return &Inputs{strings.Split(lines[0], ",")}
}
