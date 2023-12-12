package helpers

import (
	"testing"
)

func testLines() []string {
	return []string {
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
}

func TestPathLength(t *testing.T) {
	i := Enrich(ParseLines(testLines()))
	var act, exp int
	act, exp = i.PathLength(4, 8), 9
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = i.PathLength(0, 6), 15
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = i.PathLength(2, 5), 17
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = i.PathLength(7, 8), 5
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}