package helpers

import (
	"day12/part1/helpers"
	"testing"
)

func TestExpandRow(t *testing.T) {
	row := helpers.Row{
		States: []helpers.SpringState{
			helpers.Operational,
			helpers.Damaged,
			helpers.Damaged,
			helpers.Unknown,
			helpers.Damaged,
		},
		Contigs: []int{
			2,
			1,
		},
	}
	expanded := ExpandRow(row, 3)
	if len(expanded.States) != 15 {
		t.Errorf("Expected 15 states, got %v", len(expanded.States))
	}
	if expanded.States[5] != helpers.Operational ||
		expanded.States[8] != helpers.Unknown ||
		expanded.States[9] != helpers.Damaged {
		t.Errorf("Expected test states to be Operational, Damaged, Damaged, Unknown, Damaged, got %v", expanded.States)
	}

	if len(expanded.Contigs) != 6 {
		t.Errorf("Expected 6 contigs, got %v", len(expanded.Contigs))
	}
	if expanded.Contigs[0] != expanded.Contigs[2] ||
		expanded.Contigs[1] != expanded.Contigs[3] ||
		expanded.Contigs[0] != 2 || expanded.Contigs[1] != 1 {
		t.Errorf("Expected contigs to be 2, 1, 2, 1, 2, 1, got %v", expanded.Contigs)
	}
}

func TestSplitRow(t *testing.T) {
	var states []helpers.SpringState
	var exp [][]helpers.SpringState
	states = []helpers.SpringState{
		helpers.Operational,
		helpers.Operational,
	}
	exp = SplitRow(states)
	if len(exp) > 0 {
		t.Errorf("Expected empty split, got %v", exp)
	}
	states = []helpers.SpringState{
		helpers.Unknown,
		helpers.Damaged,
		helpers.Unknown,
	}
	exp = SplitRow(states)
	if len(exp) != 1 || len(exp[0]) != 3 {
		t.Errorf("Expected 1 split of 3, got %v", exp)
	}
	states = []helpers.SpringState{
		helpers.Operational,
		helpers.Unknown,
		helpers.Operational,
		helpers.Damaged,
	}
	exp = SplitRow(states)
	if len(exp) != 2 || len(exp[0]) != 1 || len(exp[1]) != 1 ||
		exp[0][0] != helpers.Unknown || exp[1][0] != helpers.Damaged {
		t.Errorf("Expected 2 splits of 1, got %v", exp)
	}
	states = []helpers.SpringState{
		helpers.Unknown,
		helpers.Damaged,
		helpers.Operational,
		helpers.Operational,
		helpers.Damaged,
		helpers.Damaged,
	}
	exp = SplitRow(states)
	if len(exp) != 2 || len(exp[0]) != 2 || len(exp[1]) != 2 {
		t.Errorf("Expected 2 splits of 2, got %v", exp)
	}
}
