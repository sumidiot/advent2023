package helpers

import (
	"testing"
)

var TestLines = []string {
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

func TestParseInput(t *testing.T) {
	inputs := ParseLines(TestLines)
	if len(inputs.Instructions) != 2 {
		t.Errorf("Expected 2 instructions, got %d", len(inputs.Instructions))
	}
	if len(inputs.Nodes) != 7 {
		t.Errorf("Expected 8 nodes, got %d", len(inputs.Nodes))
	}
}