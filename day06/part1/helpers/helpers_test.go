package helpers

import "testing"

func testLines() []string {
	return []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
}

func TestBruteForceWaysToWin(t *testing.T) {
	var act, exp int
	act, exp = BruteForceWaysToWin(Race{7, 9}), 4
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = BruteForceWaysToWin(Race{15, 40}), 8
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = BruteForceWaysToWin(Race{30, 200}), 9
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestSolvedWaysToWin(t *testing.T) {
	var act, exp int
	act, exp = SolvedWaysToWin(Race{7, 9}), 4
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolvedWaysToWin(Race{15, 40}), 8
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolvedWaysToWin(Race{30, 200}), 9
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestSolveInput(t *testing.T) {
	act, exp := SolveInput(ParseInput(testLines())), 288
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}