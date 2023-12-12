package helpers

import (
	"day10/part1/helpers"
	"testing"
)

func GetInput1() []string {
	return []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
}

func GetInput2() []string {
	return []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}
}

func GetInput3() []string {
	return []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
}

func GetInput4() []string {
	return []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
}

func GetInput5() []string {
	return []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
}

func TestInside1(t *testing.T) {
	var exp, act int
	act, exp = SolveInput(helpers.ParseLines(GetInput1()), PathFinder(helpers.Coord{1, 2})), 1
	if act != exp {
	t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveInput(helpers.ParseLines(GetInput2()), PathFinder(helpers.Coord{2, 1})), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveInput(helpers.ParseLines(GetInput3()), PathFinder(helpers.Coord{1, 2})), 4
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveInput(helpers.ParseLines(GetInput4()), PathFinder(helpers.Coord{4, 13})), 8
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = SolveInput(helpers.ParseLines(GetInput5()), PathFinder(helpers.Coord{1, 4})), 10
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestCoordInside(t *testing.T) {
	var i *helpers.Inputs
	var poly []helpers.Coord
	i = helpers.ParseLines(GetInput1())
	poly = FindPathGivenStart(i, helpers.Coord{1, 2})
	if !IsInside(helpers.Coord{Row: 2, Col: 2}, poly) {
		t.Errorf("Expected true, got false")
	}
	if IsInside(helpers.Coord{Row: 1, Col: 4}, poly) {
		t.Errorf("Expected false, got true")
	}
	if IsInside(helpers.Coord{Row: 2, Col: 4}, poly) {
		t.Errorf("Expected false, got true")
	}
	if IsInside(helpers.Coord{Row: 3, Col: 4}, poly) {
		t.Errorf("Expected false, got true")
	}
	i = helpers.ParseLines(GetInput2())
	poly = FindPathGivenStart(i, helpers.Coord{2, 1})
	if !IsInside(helpers.Coord{Row: 2, Col: 2}, poly) {
		t.Errorf("Expected true, got false")
	}
	i = helpers.ParseLines(GetInput3())
	poly = FindPathGivenStart(i, helpers.Coord{1, 2})
	if !IsInside(helpers.Coord{Row: 6, Col: 2}, poly) {
		t.Errorf("Expected true, got false")
	}
	i = helpers.ParseLines(GetInput4())
	poly = FindPathGivenStart(i, helpers.Coord{4, 13})
	if IsInside(helpers.Coord{Row: 1, Col: 16}, poly) {
		t.Errorf("Expected false, got true")
	}
	poly = PathFinder(helpers.Coord{4, 13})(i)
	if IsInside(helpers.Coord{Row: 1, Col: 16}, poly) {
		t.Errorf("Expected false, got true")
	}
	poly = FindPathGivenStart(i, helpers.Coord{4, 13})
	if IsInside(helpers.Coord{Row: 1, Col: 17}, poly) {
		t.Errorf("Expected false, got true")
	}
	poly = FindPathGivenStart(i, helpers.Coord{4, 13})
	if IsInside(helpers.Coord{Row: 1, Col: 19}, poly) {
		t.Errorf("Expected false, got true")
	}
}

func TestIntersect(t *testing.T) {
	if intersect(helpers.Coord{1, 4}, helpers.Coord{1, 3}, helpers.Coord{2, 3}) {
		t.Errorf("Expected false, got true")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{2, 3}, helpers.Coord{3, 3}) {
		t.Errorf("Expected false, got true")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{3, 3}, helpers.Coord{2, 3}) {
		t.Errorf("Expected false, got true")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{3, 3}, helpers.Coord{3, 2}) {
		t.Errorf("Expected false, got true")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{3, 2}, helpers.Coord{3, 2}) {
		t.Errorf("Expected false, got true")
	}

	if !intersect(helpers.Coord{2, 2}, helpers.Coord{1, 1}, helpers.Coord{1, 2}) {
		t.Errorf("Expected true, got false")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{1, 2}, helpers.Coord{1, 1}) {
		t.Errorf("Expected false, got true")
	}
	if !intersect(helpers.Coord{2, 2}, helpers.Coord{1, 1}, helpers.Coord{2, 1}) {
		t.Errorf("Expected true, got false")
	}
	if intersect(helpers.Coord{2, 2}, helpers.Coord{2, 1}, helpers.Coord{1, 1}) {
		t.Errorf("Expected false, got true")
	}
}

func TestIntersectGeneral(t *testing.T) {
	var hits bool
	hits, _ = intersectGeneral(
		helpers.Coord{-1, -1},
		helpers.Coord{0, 0},
		helpers.Coord{2, 14},
		helpers.Coord{1, 14},
	)
	if hits {
		t.Errorf("Expected false, got true")
	}
	hits, _ = intersectGeneral(
		helpers.Coord{-1, -1},
		helpers.Coord{0, 0},
		helpers.Coord{2, 14},
		helpers.Coord{1, 14},
	)
	if hits {
		t.Errorf("Expected false, got true")
	}
	hits, _ = intersectGeneral(
		helpers.Coord{-1, -1},
		helpers.Coord{1, 16},
		helpers.Coord{1, 7},
		helpers.Coord{0, 7},
	)
	if hits {
		t.Errorf("Expected false, got true")
	}
}
