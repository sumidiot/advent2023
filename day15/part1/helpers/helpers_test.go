package helpers

import (
	"testing"
)

func TestHash(t *testing.T) {
	ins := []string{
		"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7",
	}
	exps := []int{
		30, 253, 97, 47, 14, 180, 9, 197, 48, 214, 231,
	}
	for idx, in := range ins {
		exp, act := Hash(in), exps[idx]
		if exp != act {
			t.Errorf("Expected %d to equal %d", act, exp)
		}
	}
}
