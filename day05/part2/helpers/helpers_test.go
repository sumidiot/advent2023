package helpers

import (
	"day05/part1/helpers"
	"fmt"
	"log"
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

func TestShowEnds(t *testing.T) {
	i := helpers.ParseSeedMaps(testLines())
	for layer := 0; layer < 7; layer++ {
		log.Println("Layer", layer)
		m := i.GetMapAt(layer)
		for _, im := range(*m) {
			log.Println("  ", im.SourceStart, im.SourceStart + im.RangeLength)
		}
	}
}

func TestPropogateEndpoints(t *testing.T) {
	intervals := []Interval {
		Interval { 57, 69 },
		Interval { 81, 94 },
	}
	maps := []helpers.IdMap {
		helpers.IdMap { 49, 53, 8 },
		helpers.IdMap { 0, 11, 42 },
		helpers.IdMap { 42, 0, 7 },
		helpers.IdMap { 57, 7, 4 },
	}
	act := propogateEndpoints(&maps, intervals)
	log.Println(fmt.Sprintf("%v", act))
	if len(act) != 3 {
		t.Errorf("Expect 3 got %v in %v", len(act), act)
	}
}

func TestSplitInterval(t *testing.T) {
	interval := Interval { 57, 69 }
	ends := []int { 0, 7, 11, 53, 61 }
	act := splitInterval(interval, ends)
	if len(act) != 2 {
		t.Errorf("Expect 2 got %v in %v", len(act), act)
	}
	if act[0].left != 57 {
		t.Errorf("Expect 57 got %v in %v", act[0].left, act)
	}
	if act[0].right != 60 {
		t.Errorf("Expect 60 got %v in %v", act[0].right, act)
	}
	if act[1].left != 61 {
		t.Errorf("Expect 61 got %v in %v", act[1].left, act)
	}
	if act[1].right != 69 {
		t.Errorf("Expect 69 got %v in %v", act[1].right, act)
	}
	interval2 := Interval { 3736161691, 3908507816 }
	ends2 := []int { 0,180112079,377309295,550227227,620489350,761650773,977633281,1134944893,1310995831,1320824602,1410745074,2073610212,2268770144,2340196738,2371714487,2511534001,2626319859,2810146607,3055921442,3111697940,3233532062,3243594344,3354932755,3399998845,3407181792,3477485810,3751501334,3759669547,3786530179,3826838977,3864598346,3901395270,3922686886,3930358410,4024199858,4040990485,4163445628,4164003330,4294967296}
	splits := splitInterval(interval2, ends2)
	log.Println(fmt.Sprintf("%v", splits))
	sumCheck(splits, interval2)
}

func TestSourceMapEndpoints(t *testing.T) {
	ims := []helpers.IdMap {
		helpers.IdMap { 49, 53, 8 },
		helpers.IdMap { 0, 11, 42 },
		helpers.IdMap { 42, 0, 7 },
		helpers.IdMap { 57, 7, 4 },
	}
	act := sourceMapEndpoints(ims)
	if len(act) != 5 {
		t.Errorf("Expect 5 got %v in %v", len(act), act)
	}
	if act[0] != 0 {
		t.Errorf("Expect 0 got %v in %v", act[0], act)
	}
	if act[1] != 7 {
		t.Errorf("Expect 7 got %v in %v", act[1], act)
	}
	if act[2] != 11 {
		t.Errorf("Expect 11 got %v in %v", act[2], act)
	}
	if act[3] != 53 {
		t.Errorf("Expect 53 got %v in %v", act[3], act)
	}
	if act[4] != 61 {
		t.Errorf("Expect 61 got %v in %v", act[4], act)
	}
}