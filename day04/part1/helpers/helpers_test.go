package helpers

import (
	"testing"
)

func TestSolveCard(t *testing.T) {
	c1 := Card { 1, []int{ 41, 48, 83, 86, 17,}, []int{ 83, 86,  6, 31, 17,  9, 48, 53, }}
	c2 := Card { 2, []int{ 13, 32, 20, 16, 61,}, []int{ 61, 30, 68, 82, 17, 32, 24, 19, }}
	c3 := Card { 3, []int{  1, 21, 53, 59, 44,}, []int{ 69, 82, 63, 72, 16, 21, 14,  1, }}
	c4 := Card { 4, []int{ 41, 92, 73, 84, 69,}, []int{ 59, 84, 76, 51, 58,  5, 54, 83, }}
	c5 := Card { 5, []int{ 87, 83, 26, 28, 32,}, []int{ 88, 30, 70, 12, 93, 22, 82, 36, }}
	c6 := Card { 6, []int{ 31, 18, 13, 56, 72,}, []int{ 74, 77, 10, 23, 35, 67, 36, 11, }}
	s1 := SolveCard(c1)
	s2 := SolveCard(c2)
	s3 := SolveCard(c3)
	s4 := SolveCard(c4)
	s5 := SolveCard(c5)
	s6 := SolveCard(c6)
	if s1 != 8 {
		t.Errorf("C1 should be 8, got %d", s1)
	}
	if s2 != 2 {
		t.Errorf("C2 should be 2, got %d", s2)
	}
	if s3 != 2 {
		t.Errorf("C3 should be 2, got %d", s3)
	}
	if s4 != 1 {
		t.Errorf("C4 should be 1, got %d", s4)
	}
	if s5 != 0 {
		t.Errorf("C5 should be 0, got %d", s5)
	}
	if s6 != 0 {
		t.Errorf("C6 should be 0, got %d", s6)
	}
}