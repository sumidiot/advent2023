package helpers

import (
	"day10/part1/helpers"
	"fmt"
	"log"
)

func Solve(lines []string) int {
	i := helpers.ParseLines(lines)
	return SolveInput(i, FindPathAssumeTwoStarts)
}

func SolveInput(i *helpers.Inputs, pf func(i *helpers.Inputs) []helpers.Coord) int {
	poly := pf(i)
	IsOnPath := ToMap(poly)
	count := 0
	for row := 0; row < len(i.Grid); row++ {
		for col := 0; col < len(i.Grid[row]); col++ {
			if !IsOnPath[helpers.Coord{Row: row, Col: col}] {
				if IsInside(helpers.Coord{Row: row, Col: col}, poly) {
					log.Printf("Found %v, %v inside", row, col)
					count++
				}
			}
		}
	}
	return count
}

func ToMap(poly []helpers.Coord) map[helpers.Coord]bool {
	ret := make(map[helpers.Coord]bool)
	for _, coord := range poly {
		ret[coord] = true
	}
	return ret
}

// Initial implementation from Part 1 was incorrect
/*
func IsInside(coord helpers.Coord, poly []helpers.Coord) bool {
	dSign := PointsUp(coord, poly[0], poly[len(poly)-1])
	for i := 1; i < len(poly); i++ {
		if PointsUp(coord, poly[i], poly[i-1]) != dSign {
			return false
		}
	}
	return true
}
*/

var (
	TestPoint  helpers.Coord = helpers.Coord{Row: -1, Col: -1}
	TestPoint2 helpers.Coord = helpers.Coord{Row: -8745, Col: -459}
	TestPoint3 helpers.Coord = helpers.Coord{Row: -1, Col: -12345}
	TestPoint4 helpers.Coord = helpers.Coord{Row: -988765, Col: -1}
	TestPoint5 helpers.Coord = helpers.Coord{Row: -13579, Col: -2468}
	TestPoint6 helpers.Coord = helpers.Coord{Row: 10, Col: 100000}
	TestPoint7 helpers.Coord = helpers.Coord{Row: 100000, Col: 10}
)

/** CoPilot chat gave me the next two functions
 * I tweaked them for the types in this file
 * Also they didn't seem to work correctly if the ray went through a vertex!?
 * Maybe a <= should be a < somewhere in the intersects method?
 * Changed the <= 1 to < 1 in intersects below, but that wasn't quite right either
 */
func IsInside(point helpers.Coord, polygon []helpers.Coord) bool {
	// Originally I had this outPoints array in the testPointIntersectGeneral,
	// but that's sending multiple rays out, which isn't right
	outPoints := []helpers.Coord{TestPoint, TestPoint2, TestPoint3, TestPoint4, TestPoint5, TestPoint6, TestPoint7}
	for idx, _ := range outPoints {
		count := 0
		errs := 0
		for i := 0; i < len(polygon); i++ {
			j := (i + 1) % len(polygon)
			hits, err := testPointIntersectGeneral(outPoints[idx:(idx+1)], point, polygon[i], polygon[j])
			if hits {
				log.Printf("Found %v, %v hitting %v -> %v", point.Row, point.Col, polygon[i], polygon[j])
				count++
			}
			if err {
				errs++
			}
		}
		if errs == 0 {
			return count%2 == 1
		}
	}
	panic(fmt.Sprintf("No suitable outpoint found for %v", point))
}
func IsInsideF(tpr, tpc float64, polygon []helpers.Coord) bool {
	count := 0
	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		if intersectF(tpr, tpc, polygon[i], polygon[j]) {
			count++
		}
	}
	return count%2 == 1
}

func testPointIntersectGeneral(outPoints []helpers.Coord, tp, p1, p2 helpers.Coord) (bool, bool) {
	count := 0
	for idx, outPoint := range outPoints {
		if idx > 0 {
			log.Printf("Testing test point idx %v at %v for %v hitting %v -> %v", idx, outPoint, tp, p1, p2)
		}
		hits, isVert := intersectGeneral(outPoint, tp, p1, p2)
		if hits {
			count++
			if !isVert {
				return true, false
			}
		} else {
			return false, false
		}
	}
	if count == len(outPoints) {
		return false, true
	}
	return false, false
}

/**
 * CoPilot put this in, but I haven't dug in to the math, so don't understand how to
 * check about endpoints. Haphazardly changing inequalities between strict or not.
 * Re-derived my own from scratch, below.
 */
func intersectGeneral(p1, p2, p3, p4 helpers.Coord) (bool, bool) {
	d := (p2.Col-p1.Col)*(p4.Row-p3.Row) - (p2.Row-p1.Row)*(p4.Col-p3.Col)
	if d == 0 {
		return false, false
	}
	// r := ((p1.Row-p3.Row)*(p4.Col-p3.Col) - (p1.Col-p3.Col)*(p4.Row-p3.Row)) / d
	// s := ((p1.Row-p3.Row)*(p2.Col-p1.Col) - (p1.Col-p3.Col)*(p2.Row-p1.Row)) / d
	r := (float64((p3.Col-p1.Col)*(p2.Row-p1.Row)) - float64((p3.Row-p1.Row)*(p2.Col-p1.Col))) / float64(d)
	s := (float64(p3.Col-p1.Col) + float64(p4.Col-p3.Col)*r) / float64(p2.Col-p1.Col)
	ans := r >= 0 && r <= 1 && s >= 0 && s <= 1
	isVert := r == 0 || r == 1 || s == 0 || s == 1
	return ans, isVert
}

// assumes a ray to -1, -1. checking if (-1, -1) -> tp intersects p1 -> p2 not at the end
/**
 * Solved by hand on pencil and paper, for funsies!
 */
func intersect(tp, p1, p2 helpers.Coord) bool {
	return intersectF(float64(tp.Row), float64(tp.Col), p1, p2)
}

func intersectF(tpr, tpc float64, p1, p2 helpers.Coord) bool {
	numL := (float64(p1.Row) + 1.0) * (tpc + 1.0)
	numR := (float64(p1.Col) + 1.0) * (tpr + 1.0)
	denL := (tpr + 1.0) * float64(p2.Col-p1.Col)
	denR := (tpc + 1.0) * float64(p2.Row-p1.Row)
	num := numL - numR
	den := denL - denR
	r := float64(num) / float64(den)
	r2 := float64(float64(p2.Row-p1.Row)*r+float64(p1.Row+1)) / float64(tpr+1)
	// if r >= 0 && r < 1 && r2 > 0 && r2 < 1 && (r == 0 || r == 1) {
	// log.Printf("To %v, %v intersects %v -> %v at %v, %v", tpr, tpc, p1, p2, r, r2)
	// }
	return r >= 0 && r < 1 && r2 > 0 && r2 < 1
}

// we know (-1, -1) -> tp intersects p1 -> p2 at a vertex
// r2 is in (0, 1) along (-1, -1) -> at the point of intersection
// test points r2 +- epsilon, see if they are both in or out
// point is (-1 + (tp.Row + 1) * r2, -1 + (tp.Col + 1) * r2
func Crosses(tpr, tpc float64, p1, p2 helpers.Coord, r, r2 float64) bool {
	t1r := -1.0 + (tpr+1.0)*r2 - 0.000001
	t1c := -1.0 + (tpc+1.0)*r2 - 0.000001
	t2r := -1.0 + (tpr+1.0)*r2 + 0.000001
	t2c := -1.0 + (tpr+1.0)*r2 + 0.000001
	in1 := IsInsideF(t1r, t1c, []helpers.Coord{p1, p2})
	in2 := IsInsideF(t2r, t2c, []helpers.Coord{p1, p2})
	return in1 != in2
}

func PointsUp(a, b, c helpers.Coord) bool {
	// CoPilot put this in, I double-checked it, seems right
	return (b.Col-a.Col)*(c.Row-a.Row) > (b.Row-a.Row)*(c.Col-a.Col)
}

func FindPathAssumeTwoStarts(i *helpers.Inputs) []helpers.Coord {
	ret := make([]helpers.Coord, 0)
	ret = append(ret, i.Start)
	loc := helpers.Step{From: i.Start, To: helpers.FindStartSteps(i)[0]}
	for !loc.To.Equals(i.Start) {
		ret = append(ret, loc.To)
		loc = helpers.Step{From: loc.To, To: helpers.FindNextCoord(i, loc.To, loc.From)}
	}
	ret = append(ret, i.Start)
	return ret
}

func PathFinder(coord helpers.Coord) func(*helpers.Inputs) []helpers.Coord {
	return func(i *helpers.Inputs) []helpers.Coord {
		return FindPathGivenStart(i, coord)
	}
}

func FindPathGivenStart(i *helpers.Inputs, coord helpers.Coord) []helpers.Coord {
	ret := make([]helpers.Coord, 0)
	ret = append(ret, i.Start)
	loc := helpers.Step{From: i.Start, To: coord}
	for !loc.To.Equals(i.Start) {
		ret = append(ret, loc.To)
		loc = helpers.Step{From: loc.To, To: helpers.FindNextCoord(i, loc.To, loc.From)}
	}
	ret = append(ret, i.Start)
	return ret
}
