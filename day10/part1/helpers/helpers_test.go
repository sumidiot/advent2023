package helpers

import (
	"testing"
)

func GetInput1() []string {
	return []string {
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
}

func GetInput2() []string {
	return []string {
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}
}

func TestParse(t *testing.T) {
	ParseLines(GetInput1())
	ParseLines(GetInput2())
}

func TestFindStarts(t *testing.T) {
	var act []Dir
	act = FindStartDirectionOptions(ParseLines(GetInput1()))
	if len(act) != 2 || act[0] != Down || act[1] != Right {
		t.Errorf("Expected [Down, Right], got %v", act)
	}
	act = FindStartDirectionOptions(ParseLines(GetInput2()))
	if len(act) != 2 || act[0] != Down || act[1] != Right {
		t.Errorf("Expected [Down, Right], got %v", act)
	}
}

func TestFindNextCoord(t *testing.T) {
	var act, exp Coord
	in2 := ParseLines(GetInput2())
	testPair := func (cur Coord, one Coord, two Coord) {
		act, exp = FindNextCoord(in2, cur, one), two
		if act != exp {
			t.Errorf("Expected %v, got %v in %v", exp, act, cur)
		}
		act, exp = FindNextCoord(in2, cur, two), one
		if act != exp {
			t.Errorf("Expected %v, got %v in %v", exp, act, cur)
		}
	}
	testPair(Coord{1, 1}, Coord{1, 2}, Coord{2, 1}) // F
	testPair(Coord{1, 2}, Coord{1, 1}, Coord{0, 2}) // J
	testPair(Coord{1, 3}, Coord{0, 3}, Coord{2, 3}) // |
	testPair(Coord{3, 2}, Coord{3, 1}, Coord{3, 3}) // -
	testPair(Coord{2, 3}, Coord{1, 3}, Coord{2, 4}) // L
	testPair(Coord{2, 4}, Coord{2, 3}, Coord{3, 4}) // 7
}