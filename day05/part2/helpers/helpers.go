package helpers

import (
	"day05/part1/helpers"
	"fmt"
	"log"
	"sort"
)

func Solve(lines []string) int {
	return SolveWithEndpoints(helpers.ParseSeedMaps(lines))
}

func BruteForceSolveInput(i *helpers.Inputs) int {
	ret := i.SeedLocation(i.Seeds[0])
	for pidx := 0; pidx < len(i.Seeds); pidx += 2 {
		log.Println("Up to pair starting at", pidx)
		start := i.Seeds[pidx]
		num := i.Seeds[pidx+1]
		for seed := start; seed < start+num; {
			l := i.SeedLocation(seed)
			if l < ret {
				ret = l
			}
			seed++
		}
	}
	return ret
}

type Interval struct {
	left int
	right int
}

func SolveWithEndpoints(i *helpers.Inputs) int {
	intervals := seedIntervals(i)
	intervals = propogateEndpoints(i.GetMapAt(0), intervals)
	fmt.Println("After level 0", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(1), intervals)
	fmt.Println("After level 1", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(2), intervals)
	fmt.Println("After level 2", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(3), intervals)
	fmt.Println("After level 3", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(4), intervals)
	fmt.Println("After level 4", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(5), intervals)
	fmt.Println("After level 5", fmt.Sprintf("%v", len(intervals)))
	intervals = propogateEndpoints(i.GetMapAt(6), intervals)
	fmt.Println("After level 6", fmt.Sprintf("%v", len(intervals)))
	return minEnd(intervals)
}

func seedIntervals(i *helpers.Inputs) []Interval {
	intervals := make([]Interval, 0)
	for pidx := 0; pidx < len(i.Seeds); pidx += 2 {
		start := i.Seeds[pidx]
		num := i.Seeds[pidx+1]
		intervals = append(intervals, Interval{start, start+num-1})
	}
	return intervals
}

func propogateEndpoints(m *[]helpers.IdMap, intervals []Interval) []Interval {
	ends := sourceMapEndpoints(*m)
	newIntervals := make([]Interval, 0)
	for _, interval := range(intervals) {
		split := splitInterval(interval, ends)
		sumCheck(split, interval)
		for _, i := range(split) {
			newIntervals = append(newIntervals, Interval{helpers.TransformMany(i.left, *m), helpers.TransformMany(i.right, *m)})
		}
	}
	return newIntervals
}

func sumCheck(intervals []Interval, orig Interval) {
	sum := 0
	for _, interval := range(intervals) {
		sum += interval.right - interval.left + 1
	}
	if sum != orig.right - orig.left + 1 {
		panic(fmt.Sprintf("Sum check failed: %v %v", sum, orig.right - orig.left + 1))
	}
}

func minEnd(intervals []Interval) int {
	min := intervals[0].left
	for _, interval := range(intervals) {
		if interval.left < min {
			min = interval.left
		}
	}
	return min
}

func splitInterval(interval Interval, ends []int) []Interval {
	ret := make([]Interval, 0)
	remToSplit := &interval
	for _, end := range(ends) {
		if end > remToSplit.left && end <= remToSplit.right {
			if (remToSplit.left != remToSplit.right) {
				ret = append(ret, Interval{remToSplit.left, end - 1})
				remToSplit = &Interval{end, remToSplit.right}
			} else {
				break
			}
		} else if end > remToSplit.right {
			break
		}
	}
	ret = append(ret, *remToSplit)
	return ret
}

func sourceMapEndpoints(ims []helpers.IdMap) []int {
	ret := map[int]bool{}
	ret[0] = true
	for _, im := range(ims) {
		ret[im.SourceStart] = true
		ret[im.SourceStart + im.RangeLength] = true
	}
	keys := make([]int, 0)
	for k, _ := range(ret) {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}