package helpers

import (
	"math"
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseInput(lines))
}

func SolveInput(races []Race) int {
	prod := 1
	for _, r := range(races) {
		prod *= SolvedWaysToWin(r)
	}
	return prod
}

type Race struct {
	Time int
	Distance int
}

func ParseInput(lines []string) []Race {
	times := strings.Fields(lines[0])[1:]
	dists := strings.Fields(lines[1])[1:]
	ret := make([]Race, 0)
	for idx, tStr := range(times) {
		t, _ := strconv.Atoi(tStr)
		d, _ := strconv.Atoi(dists[idx])
		ret = append(ret, Race { t, d })
	}
	return ret
}

/**
 * Solve (t - n)*n > d where d and t are constant
 * tn - n^2 - d > 0
 * n^2 + - tn + d < 0
 * n = (t +- sqrt(t^2 - 4d)) / 2
 */
func SolvedWaysToWin(r Race) int {
	mid := float64(r.Time) / 2
	delta := math.Sqrt(math.Pow(float64(r.Time), 2) - float64(4 * r.Distance))/2
	left := math.Ceil(mid - delta)
	right := math.Floor(mid + delta)
	ret := int(right - left + 1)
	if (left < right && ScoreIfHold(int(left), r) <= r.Distance) {
		ret--
	}
	if (left < right && ScoreIfHold(int(right), r) <= r.Distance) {
		ret--
	}
	return ret
}

func BruteForceWaysToWin(r Race) int {
	ret := 0
	for i := 1; i < r.Time; i++ {
		if ScoreIfHold(i, r) > r.Distance {
			ret++
		}
	}
	return ret
}

func ScoreIfHold(n int, r Race) int {
	return (r.Time - n) * n
}