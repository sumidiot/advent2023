package helpers

import (
	"testing"
)

func TestSolveThatPanicedOnce(t *testing.T) {
	lines := []string {
		"##..###..#.####",
		".........#.##.#",
		"##...#...##..##",
		"##..##.##..##..",
		"##.###.##.####.",
		"......##..####.",
		"###.###.#......",
		"...#..###..##..",
		"....#..########",
		"..##.###.#.##.#",
		"###.###..#.##.#",
		"##.##.###.#..#.",
		"...#...##......",
		"##.......#.##.#",
		"..#####...####.",
		"...##.#...#..#.",
		"##.####.#######",
	}
	act, exp := Solve(lines), 12
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}