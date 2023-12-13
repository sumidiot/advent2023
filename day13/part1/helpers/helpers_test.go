package helpers

import (
	"testing"
)

func testLines() *Inputs {
	return ParseLines([]string {
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	})
}

func TestParse(t *testing.T) {
	i := testLines()
	if len(*i) != 2 {
		t.Errorf("Expected 2 grids, got %v", len(*i))
	}
	if len((*i)[0]) != 7 {
		t.Errorf("Expected grid 0 to have 7 rows, got %v", len((*i)[0]))
	}
	if len((*i)[1]) != 7 {
		t.Errorf("Expected grid 1 to have 7 rows, got %v", len((*i)[1]))
	}
}

func TestFindVertical(t *testing.T) {
	i := testLines()
	vert, isVert := FindVertical((*i)[0])
	if !isVert {
		t.Errorf("Expected vertical mirror, got %v", isVert)
	}
	if vert != 4 {
		t.Errorf("Expected vertical mirror at col 4, got %v", vert)
	}
	horiz, isHoriz := FindHorizontal((*i)[0])
	if isHoriz {
		t.Errorf("Expected no horizontal mirror, got %v", horiz)
	}
}

func TestFindHorizontal(t *testing.T) {
	i := testLines()
	horiz, isHoriz := FindHorizontal((*i)[1])
	if !isHoriz {
		t.Errorf("Expected horizontal mirror, got %v", isHoriz)
	}
	if horiz != 3 {
		t.Errorf("Expected horizontal mirror at row 3, got %v", horiz)
	}
	vert, isVert := FindVertical((*i)[1])
	if isVert {
		t.Errorf("Expected no vertical mirror, got %v", vert)
	}
}

func TestSolve(t *testing.T) {
	i := testLines()
	ans, exp := SolveInput(i), 405
	if ans != exp {
		t.Errorf("Expected %v, got %v", exp, ans)
	}
}