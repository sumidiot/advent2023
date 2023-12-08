package helpers

import (
	"testing"
)

// "day08/part1/helpers"
// "testing"

// my numbers worked out to be the following
// 2023/12/07 22:28:28 [[22199] [14893] [16579] [17141] [18827] [13207]]
// 22199 = 79 * 281
// 14893 = 53 * 281
// 16579 = 59 * 281
// 17141 = 61 * 281
// 18827 = 67 * 281
// 13207 = 47 * 281

func TestPairwiseGCDs(t *testing.T) {
	num := 281*79*53*59*61*67*47
	if num != 13334102464297 {
		t.Errorf("Expected 0, got %v", num)
	}
}