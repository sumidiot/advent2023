package helpers

import (
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseSeedMaps(lines))
}

func SolveInput(i *Inputs) int {
	ret := i.SeedLocation(i.Seeds[0])
	for _, s := range(i.Seeds[1:]) {
		l := i.SeedLocation(s)
		if l < ret {
			ret = l
		}
	}
	return ret
}

type Inputs struct {
	Seeds []int
	seedToSoil []IdMap
	soilToFert []IdMap
	fertToWater []IdMap
	waterToLight []IdMap
	lightToTemp []IdMap
	tempToHum []IdMap
	humToLoc []IdMap
}

func (i *Inputs) GetMapAt(idx int) *[]IdMap {
	switch idx {
	case 0: return &i.seedToSoil
	case 1: return &i.soilToFert
	case 2: return &i.fertToWater
	case 3: return &i.waterToLight
	case 4: return &i.lightToTemp
	case 5: return &i.tempToHum
	case 6: return &i.humToLoc
	default: panic("wtf")
	}
}

func emptyInputs() *Inputs {
	return &Inputs {
		make([]int, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
		make([]IdMap, 0),
	}
}

type IdMap struct {
	DestStart int
	SourceStart int
	RangeLength int
}

func ParseSeedMaps(lines []string) *Inputs {
	ret := emptyInputs()
	ret.Seeds = parseSeeds(lines[0])
	lidx := 3
	midx := 0
	for ; lidx < len(lines); lidx++ {
		if len(lines[lidx]) == 0 {
			midx++
			lidx++
		} else {
			fs := strings.Fields(lines[lidx])
			ds, _ := strconv.Atoi(fs[0])
			ss, _ := strconv.Atoi(fs[1])
			rl, _ := strconv.Atoi(fs[2])
			im := IdMap { ds, ss, rl }
			m := ret.GetMapAt(midx)
			*m = append(*m, im)
		}
	}
	return ret
}

func parseSeeds(line string) []int {
	fs := strings.Fields(line)
	ret := make([]int, 0)
	for _, f := range(fs[1:]) {
		n, _ := strconv.Atoi(f)
		ret = append(ret, n)
	}
	return ret
}

func TransformOne(src int, m IdMap) int {
	if src >= m.SourceStart && (src - m.SourceStart) < m.RangeLength {
		return m.DestStart + (src - m.SourceStart)
	} else {
		return src
	}
}

func TransformMany(src int, ms []IdMap) int {
	for _, im := range(ms) {
		next := TransformOne(src, im)
		if next != src {
			return next
		}
	}
	return src
}

func (i *Inputs) SeedLocation(seed int) int {
	soil := TransformMany(seed, i.seedToSoil)
	fert := TransformMany(soil, i.soilToFert)
	water := TransformMany(fert, i.fertToWater)
	light := TransformMany(water, i.waterToLight)
	temp := TransformMany(light, i.lightToTemp)
	hum := TransformMany(temp, i.tempToHum)
	return TransformMany(hum, i.humToLoc)
}