package helpers

import (
	"testing"
)

func TestFindPartsAdjacent(t *testing.T) {
	var ans []int;
	ans = findPartsAdjacent(0, "...")
	if len(ans) != 0 {
		t.Errorf("Expected empty slice, got %v", ans)
	}
	ans = findPartsAdjacent(1, "...")
	if len(ans) != 0 {
		t.Errorf("Expected empty slice, got %v", ans)
	}
	ans = findPartsAdjacent(2, "...")
	if len(ans) != 0 {
		t.Errorf("Expected empty slice, got %v", ans)
	}
	ans = findPartsAdjacent(0, "1..")
	if len(ans) != 1 {
		t.Errorf("Expected empty slice, got %v", ans)
	}
}