package helpers

import (
	"day09/part1/helpers"
	"testing"
)

func TestSolveHistory(t *testing.T) {
	var h helpers.History
	var act int
	h = helpers.History { 10, 13, 16, 21, 30, 45 }
	act = SolveHistory(h)
	if act != 5 {
		t.Errorf("Expected %d, got %d", 5, act)
	}
}
