package helpers

import (
	"testing"
)

func TestSolveOneCardNoWinner(t *testing.T) {
	ans := Solve([]string{"Card 1: 1 3 | 2"})
	if ans != 1 {
		t.Errorf("Expected 1, got %v", ans)
	}
}

func TestSolveTwoCardsNoWinners(t *testing.T) {
	ans := Solve([]string{"Card 1: 1 3 | 2", "Card 2: 1 3 | 2"})
	if ans != 2 {
		t.Errorf("Expected 2, got %v", ans)
	}
}

func TestSolveTwoCardsOneWinners(t *testing.T) {
	ans := Solve([]string{"Card 1: 1 3 | 3", "Card 2: 1 3 | 2"})
	if ans != 3 {
		t.Errorf("Expected 3, got %v", ans)
	}
}

func TestSolveStackingMultipleWinners(t *testing.T) {
	ans := Solve([]string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	})
	if ans != 30 {
		t.Errorf("Expected 30, got %v", ans)
	}
}

func TestParse(t *testing.T) {
	ans := parse("Card 10:  1 2 |  3")
	if ans.Id != 10 ||
		len(ans.Mine) != 1 ||
		ans.Mine[0] != 3 ||
		len(ans.Winners) != 2 ||
		ans.Winners[0] != 1 ||
		ans.Winners[1] != 2 {

		t.Errorf("Expected 10, (1, 2), 3, got %v", ans)
	}
}

func TestCardNoWinners(t *testing.T) {
	ans := cardWinners(Card{1, []int{1, 2}, []int{3}})
	if ans != 0 {
		t.Errorf("Expected 0, got %v", ans)
	}
}
func TestCardWithOneWinner(t *testing.T) {
	ans := cardWinners(Card{1, []int{1, 2}, []int{2}})
	if ans != 1 {
		t.Errorf("Expected 1, got %v", ans)
	}
}
func TestCardMultipleWinners(t *testing.T) {
	ans := cardWinners(Card{1, []int{1, 2, 3}, []int{1, 3}})
	if ans != 2 {
		t.Errorf("Expected 2, got %v", ans)
	}
}

func TestSum(t *testing.T) {
	ans := sum(map[int]int{1: 1, 2: 2})
	if ans != 3 {
		t.Errorf("Expected 3, got %v", ans)
	}
}

func TestOrIncrementBy(t *testing.T) {
	m := make(map[int]int)
	n, e := m[0]
	if n != 0 || e {
		t.Errorf("Expected 0, false, got %v, %v", n, e)
	}
	incrementBy(m, 0, 1)
	n, e = m[0]
	if n != 1 || !e {
		t.Errorf("Expected 1, true, got %v, %v", n, e)
	}
	incrementBy(m, 0, 2)
	n, e = m[0]
	if n != 3 || !e {
		t.Errorf("Expected 3, true, got %v, %v", n, e)
	}
}
