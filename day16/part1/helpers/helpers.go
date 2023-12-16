package helpers

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(grid [][]Cell) int {
	isEnergized := make(map[int]map[int]bool, 0)
	traversed := map[Dir]map[int]map[int]bool{
		Up:    make(map[int]map[int]bool, 0),
		Down:  make(map[int]map[int]bool, 0),
		Left:  make(map[int]map[int]bool, 0),
		Right: make(map[int]map[int]bool, 0),
	}
	beams := []Beam{Beam{Row: 0, Col: 0, Dir: Right}}
	MakeTrue(isEnergized, 0, 0)
	MakeTrue(traversed[Right], 0, 0)
	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]
		nexts := NextBeams(grid, beam)
		for _, next := range nexts {
			MakeTrue(isEnergized, next.Row, next.Col)
			if !IsTrue(traversed[next.Dir], next.Row, next.Col) {
				MakeTrue(traversed[next.Dir], next.Row, next.Col)
				beams = append(beams, next)
			} // else: been here before
		}
	}
	energized := 0
	for _, m := range isEnergized {
		for _, t := range m {
			if t {
				energized++
			}
		}
	}
	return energized
}

func NextBeams(grid [][]Cell, beam Beam) []Beam {
	switch grid[beam.Row][beam.Col] {
	case Open:
		{
			next, ok := beam.Forward(len(grid), len(grid[0]))
			if ok {
				return []Beam{next}
			} else {
				return []Beam{}
			}
		}
	case VertSplit:
		{
			if beam.Dir == Up || beam.Dir == Down {
				next, ok := beam.Forward(len(grid), len(grid[0]))
				if ok {
					return []Beam{next}
				} else {
					return []Beam{}
				}
			} else {
				return []Beam{
					Beam{Row: beam.Row, Col: beam.Col, Dir: Up},
					Beam{Row: beam.Row, Col: beam.Col, Dir: Down},
				}
			}
		}
	case HorizSplit:
		{
			if beam.Dir == Left || beam.Dir == Right {
				next, ok := beam.Forward(len(grid), len(grid[0]))
				if ok {
					return []Beam{next}
				} else {
					return []Beam{}
				}
			} else {
				return []Beam{
					Beam{Row: beam.Row, Col: beam.Col, Dir: Left},
					Beam{Row: beam.Row, Col: beam.Col, Dir: Right},
				}
			}
		}
	case ForwardSlash:
		{
			var next Beam
			switch beam.Dir {
			case Up:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Right}
			case Down:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Left}
			case Left:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Down}
			case Right:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Up}
			default:
				panic("Unknown direction")
			}
			fwd, ok := next.Forward(len(grid), len(grid[0]))
			if ok {
				return []Beam{fwd}
			} else {
				return []Beam{}
			}
		}
	case BackSlash:
		{
			var next Beam
			switch beam.Dir {
			case Up:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Left}
			case Down:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Right}
			case Left:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Up}
			case Right:
				next = Beam{Row: beam.Row, Col: beam.Col, Dir: Down}
			default:
				panic("Unknown direction")
			}
			fwd, ok := next.Forward(len(grid), len(grid[0]))
			if ok {
				return []Beam{fwd}
			} else {
				return []Beam{}
			}
		}
	default:
		panic("Unknown cell type")
	}
}

func MakeTrue(m map[int]map[int]bool, row int, col int) {
	if m[row] == nil {
		m[row] = make(map[int]bool, 0)
	}
	m[row][col] = true
}

func IsTrue(m map[int]map[int]bool, row int, col int) bool {
	if m[row] == nil {
		return false
	}
	return m[row][col]
}

type Cell int

const (
	Open Cell = iota
	VertSplit
	HorizSplit
	ForwardSlash
	BackSlash
)

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Beam struct {
	Row int
	Col int
	Dir Dir
}

func (b Beam) Forward(numRows int, numCols int) (Beam, bool) {
	switch b.Dir {
	case Up:
		{
			if b.Row >= 1 {
				return Beam{Row: b.Row - 1, Col: b.Col, Dir: Up}, true
			} else {
				return b, false
			}
		}
	case Down:
		{
			if b.Row < numRows-1 {
				return Beam{Row: b.Row + 1, Col: b.Col, Dir: Down}, true
			} else {
				return b, false
			}
		}
	case Left:
		{
			if b.Col >= 1 {
				return Beam{Row: b.Row, Col: b.Col - 1, Dir: Left}, true
			} else {
				return b, false
			}
		}
	case Right:
		{
			if b.Col < numCols-1 {
				return Beam{Row: b.Row, Col: b.Col + 1, Dir: Right}, true
			} else {
				return b, false
			}
		}
	default:
		panic("Unknown direction")
	}
}

func ParseLines(lines []string) [][]Cell {
	cells := make([][]Cell, len(lines))
	for idx, line := range lines {
		cells[idx] = ParseLine(line)
	}
	return cells
}

func ParseLine(line string) []Cell {
	cells := make([]Cell, len(line))
	for idx, char := range line {
		switch char {
		case '.':
			cells[idx] = Open
		case '|':
			cells[idx] = VertSplit
		case '-':
			cells[idx] = HorizSplit
		case '/':
			cells[idx] = ForwardSlash
		case '\\':
			cells[idx] = BackSlash
		default:
			panic("Unknown character")
		}

	}
	return cells
}
