package helpers

import (
	"day14/part1/helpers"
	"reflect"
	"testing"
)

// "testing"

func testLines() []string {
	return []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
}

func TestSpin2(t *testing.T) {
	grid := helpers.ParseLines(testLines())
	var actLoad, expLoad int
	grid = RollNorth(grid)
	expected := helpers.ParseLines([]string{
		"OOOO.#.O..", // 10 * 5
		"OO..#....#", //  9 * 2
		"OO..O##..O", //  8 * 4
		"O..#.OO...", //  7 * 3
		"........#.", //  6 * 0
		"..#....#.#", //  5 * 0
		"..O..#.O.O", //  4 * 3
		"..O.......", //  3 * 1
		"#....###..", //  2 * 0
		"#....#....", //  1 * 0
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	actLoad, expLoad = Load(grid), 10*5 + 9*2 + 8*4 + 7*3 + 6*0 + 5*0 + 4*3 + 3*1 + 2*0 + 1*0
	if actLoad != expLoad {
		t.Errorf("load is not as expected, got %v expected %v", actLoad, expLoad)
	}
	grid = RollWest(grid)
	expected = helpers.ParseLines([]string{
		"OOOO.#O...",
		"OO..#....#",
		"OOO..##O..",
		"O..#OO....",
		"........#.",
		"..#....#.#",
		"O....#OO..",
		"O.........",
		"#....###..",
		"#....#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollSouth(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#.O..#",
		"O..O.##...",
		"O.O#......",
		"O.O....O#.",
		"O.#..O.#.#",
		"O....#....",
		"OO....OO..",
		"#O...###..",
		"#O..O#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollEast(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		"...OO##...",
		".OO#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#....",
		"......OOOO",
		"#...O###..",
		"#..OO#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = Spin(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
}

func TestSpinA(t *testing.T) {
	grid := helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		"...OO##...",
		".OO#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#....",
		"......OOOO",
		"#...O###..",
		"#..OO#....",
	})
	grid = RollNorth(grid)
	expected := helpers.ParseLines([]string{
		".OOO.#.OO.",
		".O..#....#",
		"....O##...",
		"...#OOO...",
		"...OO.O.#.",
		"..#.O.O#O#",
		".....#.O.O",
		"..........",
		"#....###..",
		"#....#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollWest(grid)
	expected = helpers.ParseLines([]string{
		"OOO..#OO..",
		"O...#....#",
		"O....##...",
		"...#OOO...",
		"OOO.....#.",
		"..#OO..#O#",
		".....#OO..",
		"..........",
		"#....###..",
		"#....#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollSouth(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#.O..#",
		".....##...",
		"..O#......",
		"O.O....O#.",
		"O.#..O.#.#",
		"O....#O...",
		"O.....OO..",
		"#O..O###..",
		"#O.OO#..O.",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollEast(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
}



func TestSpinB(t *testing.T) {
	grid := helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	})
	grid = RollNorth(grid)
	expected := helpers.ParseLines([]string{
		".OO..#.OO.",
		"....#....#",
		"....O##...",
		"...#OOO...",
		"...OO.O.#.",
		"..#O...#O#",
		"..O..#.O.O",
		".........O",
		"#....###.O",
		"#....#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollWest(grid)
	expected = helpers.ParseLines([]string{
		"OO...#OO..",
		"....#....#",
		"O....##...",
		"...#OOO...",
		"OOO.....#.",
		"..#O...#O#",
		"O....#OO..",
		"O.........",
		"#....###O.",
		"#....#....",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollSouth(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#.O..#",
		".....##...",
		"O..#......",
		"O.O....O#.",
		"O.#..O.#.#",
		"O....#O...",
		"O.....OO..",
		"#O...###O.",
		"#O.OO#..O.",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
	grid = RollEast(grid)
	expected = helpers.ParseLines([]string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#...O###.O",
		"#.OOO#...O",
	})
	if !reflect.DeepEqual(grid, expected) {
		t.Errorf("grid is not as expected: %v", grid)
	}
}


func TestMoves(t *testing.T) {
	input := helpers.ParseLines([]string{
		"OOOO.#.O..",
		"OO..#....#",
		"OO..O##..O",
		"O..#.OO...",
		"........#.",
		"..#....#.#",
		"..O..#.O.O",
		"..O.......",
		"#....###..",
		"#....#....",
	})
	expected := -2
	actual := MovesCol(input, 2, 4, -1)
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}