package helpers

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	return 0
}

type Inputs struct {
}

func ParseLines(lines []string) *Inputs {
	return &Inputs{}
}
