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

func TestSolve(t *testing.T) {
	var n, s int
	n, s = Solve(testLines(), 10)
	if n + s * 10 != 1030 {
		t.Errorf("Expected 1030 for 10, got %v", n + s * 10)
	}
	n, s = Solve(testLines(), 100)
	if n + s * 100 != 8410 {
		t.Errorf("Expected 8410 for 100, got %v", n + s * 100)
	}
}