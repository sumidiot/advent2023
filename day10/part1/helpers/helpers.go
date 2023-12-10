package helpers

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	return SolveAssumeTwoStarts(i)
}

type GridSpace int

const (
	Empty GridSpace = iota
	Vert
	Horiz
	BotLeft
	BotRight
	TopRight
	TopLeft
	Start
)

func ToGridSpace(r rune) GridSpace {
	switch r {
	case '.':
		return Empty
	case '|':
		return Vert
	case '-':
		return Horiz
	case 'L':
		return BotLeft
	case 'J':
		return BotRight
	case '7':
		return TopRight
	case 'F':
		return TopLeft
	case 'S':
		return Start
	default:
		panic("invalid grid space")
	}
}

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Coord struct {
	Row int
	Col int
}

func (c Coord) Equals(o Coord) bool {
	return c.Row == o.Row && c.Col == o.Col
}

type Inputs struct {
	Grid [][]GridSpace
	Start Coord
}

func (i *Inputs) At(c Coord) GridSpace {
	return i.Grid[c.Row][c.Col]
}

func ParseLines(lines []string) *Inputs {
	grid := make([][]GridSpace, len(lines))
	start := Coord{-1, -1}
	for i, line := range lines {
		grid[i] = make([]GridSpace, len(line))
		for j, r := range line {
			grid[i][j] = ToGridSpace(r)
			if grid[i][j] == Start {
				start = Coord{i, j}
			}
		}
	}
	return &Inputs{grid, start}
}

func FindStartDirectionOptions(i *Inputs) []Dir {
	s := i.Start
	ret := make([]Dir, 0)
	if s.Row > 0 && (i.Grid[s.Row-1][s.Col] == Vert || i.Grid[s.Row-1][s.Col] == TopLeft || i.Grid[s.Row-1][s.Col] == TopRight) {
		ret = append(ret, Up)
	}
	if s.Row < len(i.Grid) - 1 && (i.Grid[s.Row+1][s.Col] == Vert || i.Grid[s.Row+1][s.Col] == BotLeft || i.Grid[s.Row+1][s.Col] == BotRight) {
		ret = append(ret, Down)
	}
	if s.Col > 0 && (i.Grid[s.Row][s.Col-1] == Horiz || i.Grid[s.Row][s.Col-1] == BotLeft || i.Grid[s.Row][s.Col-1] == TopLeft) {
		ret = append(ret, Left)
	}
	if s.Col < len(i.Grid[0]) - 1 && (i.Grid[s.Row][s.Col+1] == Horiz || i.Grid[s.Row][s.Col+1] == BotRight || i.Grid[s.Row][s.Col+1] == TopRight) {
		ret = append(ret, Right)
	}
	return ret
}

func FindStartSteps(i *Inputs) []Coord {
	s := i.Start
	ret := make([]Coord, 0)
	if s.Row > 0 && (i.Grid[s.Row-1][s.Col] == Vert || i.Grid[s.Row-1][s.Col] == TopLeft || i.Grid[s.Row-1][s.Col] == TopRight) {
		ret = append(ret, Coord{s.Row-1, s.Col})
	}
	if s.Row < len(i.Grid) - 1 && (i.Grid[s.Row+1][s.Col] == Vert || i.Grid[s.Row+1][s.Col] == BotLeft || i.Grid[s.Row+1][s.Col] == BotRight) {
		ret = append(ret, Coord{s.Row+1, s.Col})
	}
	if s.Col > 0 && (i.Grid[s.Row][s.Col-1] == Horiz || i.Grid[s.Row][s.Col-1] == BotLeft || i.Grid[s.Row][s.Col-1] == TopLeft) {
		ret = append(ret, Coord{s.Row, s.Col-1})
	}
	if s.Col < len(i.Grid[0]) - 1 && (i.Grid[s.Row][s.Col+1] == Horiz || i.Grid[s.Row][s.Col+1] == BotRight || i.Grid[s.Row][s.Col+1] == TopRight) {
		ret = append(ret, Coord{s.Row, s.Col+1})
	}
	return ret
}

func FindNextCoord(i *Inputs, cur Coord, prev Coord) Coord {
	switch i.At(cur) {
	case Empty: panic("empty?")
	case Vert:
		if prev.Row < cur.Row {
			return Coord{cur.Row+1, cur.Col}
		} else if prev.Row > cur.Row{
			return Coord{cur.Row-1, cur.Col}
		} else {
			panic("vert")
		}
	case Horiz:
		if prev.Col < cur.Col {
			return Coord{cur.Row, cur.Col+1}
		} else if prev.Col > cur.Col{
			return Coord{cur.Row, cur.Col-1}
		} else {
			panic("horiz")
		}
	case TopLeft:
		if prev.Row == cur.Row && prev.Col == cur.Col + 1 {
			return Coord{cur.Row+1, cur.Col}
		} else if prev.Col == cur.Col && prev.Row == cur.Row + 1 {
			return Coord{cur.Row, cur.Col+1}
		} else {
			panic("top left")
		}
	case TopRight:
		if prev.Row == cur.Row && prev.Col == cur.Col - 1 {
			return Coord{cur.Row+1, cur.Col}
		} else if prev.Col == cur.Col && prev.Row == cur.Row + 1 {
			return Coord{cur.Row, cur.Col-1}
		} else {
			panic("top right")
		}
	case BotLeft:
		if prev.Row == cur.Row && prev.Col == cur.Col + 1 {
			return Coord{cur.Row-1, cur.Col}
		} else if prev.Col == cur.Col && prev.Row == cur.Row - 1 {
			return Coord{cur.Row, cur.Col+1}
		} else {
			panic("bot left")
		}
	case BotRight:
		if prev.Row == cur.Row && prev.Col == cur.Col - 1 {
			return Coord{cur.Row-1, cur.Col}
		} else if prev.Col == cur.Col && prev.Row == cur.Row - 1 {
			return Coord{cur.Row, cur.Col-1}
		} else {
			panic("bot right")
		}
	default: panic("panic")
	}
}

type Step struct {
	From Coord
	To Coord
}

func SolveAssumeTwoStarts(i *Inputs) int {
	locs := FindStartSteps(i)
	loc1 := Step{i.Start, locs[0]}
	loc2 := Step{i.Start, locs[1]}
	stepsTaken := 1
	for ; !HaveConnected(loc1, loc2); stepsTaken++ {
		loc1 = Step{loc1.To, FindNextCoord(i, loc1.To, loc1.From)}
		loc2 = Step{loc2.To, FindNextCoord(i, loc2.To, loc2.From)}
	}
	return stepsTaken
}

func HaveConnected(s1 Step, s2 Step) bool {
	// CoPilot wrote this. I probably would have only done the first two cases
	// and, actually, the final case was wrong (the start location!)
	return s1.To.Equals(s2.To) || s1.To.Equals(s2.From) // || s1.From.Equals(s2.To) || s1.From.Equals(s2.From)
}
