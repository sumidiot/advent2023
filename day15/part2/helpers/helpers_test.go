package helpers

import (
	"reflect"
	"testing"
)

// "day15/part1/helpers"

func TestParse(t *testing.T) {
	act := ParseLines([]string{ "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7" })
	exp := &([]Instruction{
			&MakeLength{ "rn", 1 },
			&Remove{ "cm" },
			&MakeLength{ "qp", 3 },
			&MakeLength{ "cm", 2 },
			&Remove{ "qp" },
			&MakeLength{ "pc", 4 },
			&MakeLength{ "ot", 9 },
			&MakeLength{ "ab", 5 },
			&Remove{ "pc" },
			&MakeLength{ "pc", 6 },
			&MakeLength{ "ot", 7 },
		})
	if !reflect.DeepEqual(act, exp) {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestMakeBoxes(t *testing.T) {
	var inss []Instruction
	var act [][]MakeLength
	inss = ParseLines([]string{ "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7" })
	act = MakeBoxes(inss)
	exp0 := []MakeLength{ MakeLength{ "rn", 1 }, MakeLength { "cm", 2}, }
	exp3 := []MakeLength{ MakeLength{ "ot", 7 }, MakeLength { "ab", 5}, MakeLength { "pc", 6 } }
	if !reflect.DeepEqual(act[0], exp0) {
		t.Errorf("Expected %v, got %v", exp0, act[0])
	}
	if !reflect.DeepEqual(act[3], exp3) {
		t.Errorf("Expected %v, got %v", exp3, act[3])
	}
}

func TestSumLengths(t *testing.T) {
	boxes := [][]MakeLength{
		[]MakeLength{ MakeLength{ "rn", 1 }, MakeLength { "cm", 2}, },
		make([]MakeLength, 0),
		make([]MakeLength, 0),
		[]MakeLength{ MakeLength{ "ot", 7 }, MakeLength { "ab", 5}, MakeLength { "pc", 6 } },
	}
	act, exp := SumLengths(boxes), 145
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}