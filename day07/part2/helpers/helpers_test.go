package helpers

import (
	"testing"
)

func TestHandType(t *testing.T) {
	var act, exp HandType
	act, exp = GetHandType([]Card{Ten, Five, Five, Jack, Five}), FourOfAKind
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = GetHandType([]Card{King, Ten, Jack, Jack, Ten}), FourOfAKind
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = GetHandType([]Card{Queen, Queen, Queen, Jack, Ace}), FourOfAKind
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}