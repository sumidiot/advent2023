package helpers

import (
	"testing"
)

func testLines() []string {
	return []string {
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
}

func TestParse(t *testing.T) {
	var r Row
	r = ParseLine("???.### 1,1,3")
	if len(r.States) != 7 || r.States[0] != Unknown || r.States[1] != Unknown || r.States[2] != Unknown || r.States[3] != Operational || r.States[4] != Damaged || r.States[5] != Damaged || r.States[6] != Damaged {
		t.Errorf("Expected 7 states, got %v in %v", len(r.States), r.States)
	}
	if len(r.Contigs) != 3 || r.Contigs[0] != 1 || r.Contigs[1] != 1 || r.Contigs[2] != 3 {
		t.Errorf("Expected 3 contigs, got %v in %v", len(r.Contigs), r.Contigs)
	}
}

func TestSolveRow(t *testing.T) {
	var act, exp int
	act, exp = SolveRow(ParseLine("???.### 1,1,3")), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveRow(ParseLine(".??..??...?##. 1,1,3")), 4
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveRow(ParseLine("?#?#?#?#?#?#?#? 1,3,1,6")), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveRow(ParseLine("????.#...#... 4,1,1")), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveRow(ParseLine("????.######..#####. 1,6,5")), 4
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveRow(ParseLine("?###???????? 3,2,1")), 10
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}