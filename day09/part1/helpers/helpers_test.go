package helpers

import (
	"testing"
)

func testLines() []string {
	return []string {
	}
}

func TestDiff(t *testing.T) {
	var h History
	var act History
	h = History { 0, 3, 6, 9, 12, 15 }
	act = Diffs(h)
	if len(act) != 5 {
		t.Errorf("Expected %d, got %d", 5, len(act))
	}
	if act[0] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[0])
	}
	if act[1] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[1])
	}
	if act[2] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[2])
	}
	if act[3] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[3])
	}
	if act[4] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[4])
	}

	h = History { 10, 13, 16, 21, 30, 45, 68 }
	act = Diffs(h)
	if len(act) != 6 {
		t.Errorf("Expected %d, got %d", 6, len(act))
	}
	if act[0] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[0])
	}
	if act[1] != 3 {
		t.Errorf("Expected %d, got %d", 3, act[1])
	}
	if act[2] != 5 {
		t.Errorf("Expected %d, got %d", 5, act[2])
	}
	if act[3] != 9 {
		t.Errorf("Expected %d, got %d", 9, act[3])
	}
	if act[4] != 15 {
		t.Errorf("Expected %d, got %d", 15, act[4])
	}
	if act[5] != 23 {
		t.Errorf("Expected %d, got %d", 23, act[5])
	}
}

func TestSolveHistory(t *testing.T) {
	var h History
	var act int
	h = History { 0, 3, 6, 9, 12, 15 }
	act = SolveHistory(h)
	if act != 18 {
		t.Errorf("Expected %d, got %d", 18, act)
	}

	h = History { 10, 13, 16, 21, 30, 45 }
	act = SolveHistory(h)
	if act != 68 {
		t.Errorf("Expected %d, got %d", 68, act)
	}
}