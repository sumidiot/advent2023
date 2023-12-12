package helpers

import (
	"testing"
)

func testLines() []string {
	return []string{
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

func TestSolve(t *testing.T) {
	var act, exp int
	act, exp = Solve(testLines(), 10), 1030
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = Solve(testLines(), 100), 8410
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}
