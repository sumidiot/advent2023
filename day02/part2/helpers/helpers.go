package helpers

import (
	"day02/part1/helpers"
)

func FindSupport(gs helpers.GameSamples) map[string]int {
	ret := make(map[string]int)
	for _, sample := range gs.Samples {
		for col, num := range sample {
			prev, ok := ret[col]
			if !ok || prev < num {
				ret[col] = num
			}
		}
	}
	return ret
}

func Power(cubes map[string]int) int {
	red, okr := cubes["red"]
	if !okr {
		red = 0
	}
	green, okg := cubes["green"]
	if !okg {
		green = 0
	}
	blue, okb := cubes["blue"]
	if !okb {
		blue = 0
	}
	return red * green * blue
}
