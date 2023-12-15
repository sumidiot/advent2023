package helpers

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	loadRow := make([]int, len((*i)[0]))
	for j := 0; j < len((*i)[0]); j++ {
		loadRow[j] = len(*i)
	}
	sum := 0
	for ridx, row := range *i {
		for cidx, cell := range row {
			switch cell {
			case Rock:
				{
					sum += loadRow[cidx]
					loadRow[cidx] = loadRow[cidx] - 1
				}
			case Pillar:
				{
					loadRow[cidx] = len(*i) - ridx - 1
				}
			}
		}
	}
	return sum
}

type CellType int

const (
	Open CellType = iota
	Pillar
	Rock
)

type Inputs [][]CellType

func ParseLines(lines []string) *Inputs {
	ret := make(Inputs, len(lines))
	for i, line := range lines {
		ret[i] = make([]CellType, len(line))
		for j, c := range line {
			switch c {
			case '.':
				ret[i][j] = Open
			case '#':
				ret[i][j] = Pillar
			case 'O':
				ret[i][j] = Rock
			default:
				panic("unexpected input")
			}
		}
	}
	return &ret
}
