package helpers

import (
	"testing"
)

func testLines() []string {
	return []string {
		"seeds: 79 14 55 13",
        "",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
}

func TestParse(t *testing.T) {
	lines := testLines()
	sm := ParseSeedMaps(lines)
	if len(sm.Seeds) != 4 {
		t.Errorf("Should have 4 seeds, got %v", len(sm.Seeds))
	}
	if len(sm.seedToSoil) != 2 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.seedToSoil))
	}
	if len(sm.soilToFert) != 3 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.soilToFert))
	}
	if len(sm.fertToWater) != 4 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.fertToWater))
	}
	if len(sm.waterToLight) != 2 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.waterToLight))
	}
	if len(sm.lightToTemp) != 3 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.lightToTemp))
	}
	if len(sm.tempToHum) != 2 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.tempToHum))
	}
	if len(sm.humToLoc) != 2 {
		t.Errorf("Should have 2 seedtoSoil, got %v", len(sm.humToLoc))
	}
}

func TestTransformOne(t *testing.T) {
	m := IdMap { 50, 98, 2 }
	act, exp := TransformOne(1, m), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformOne(50, m), 50
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformOne(98, m), 50
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformOne(99, m), 51
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformOne(100, m), 100
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestTransformMany(t *testing.T) {
	ims := []IdMap { IdMap { 50, 98, 2}, IdMap {52, 50, 48 }}
	var act, exp int
	act, exp = TransformMany(1, ims), 1
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(48, ims), 48
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(50, ims), 52
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(51, ims), 53
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(97, ims), 99
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(98, ims), 50
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(99, ims), 51
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = TransformMany(100, ims), 100
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func TestSeedLocation(t *testing.T) {
	var act, exp int
	in := ParseSeedMaps(testLines())
	act, exp = in.SeedLocation(79), 82
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = in.SeedLocation(14), 43
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = in.SeedLocation(55), 86
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	act, exp = in.SeedLocation(13), 35
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}
func TestSolve(t *testing.T) {
	var act, exp int
	act, exp = SolveInput(ParseSeedMaps(testLines())), 35
	if act != exp {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}