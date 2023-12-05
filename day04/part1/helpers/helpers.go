package helpers

import (
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	Winners []int
	Mine    []int
}

func Solve(lines []string) int64 {
	var ret int64
	ret = 0
	for _, line := range lines {
		ret += SolveCard(ParseLine(line))
	}
	return ret
}

func ParseLine(line string) Card {
	fields := strings.Fields(line)
	id, _ := strconv.Atoi(strings.TrimSuffix(fields[1], ":"))
	splitterIdx := 2
	for ; splitterIdx < len(fields) && fields[splitterIdx] != "|"; splitterIdx++ {
	}
	if len(fields) == splitterIdx {
		panic("Couldn't find splitter")
	}
	winners := make([]int, 0)
	for i := 2; i < splitterIdx; i++ {
		v, _ := strconv.Atoi(fields[i])
		winners = append(winners, v)
	}
	mine := make([]int, 0)
	for i := splitterIdx + 1; i < len(fields); i++ {
		v, _ := strconv.Atoi(fields[i])
		mine = append(mine, v)
	}
	return Card {
		id,
		winners,
		mine,
	}
}

func SolveCard(card Card) int64 {
	var ret int64
	ret = 0
	for _, mine := range card.Mine {
		matches := false
		for _, winner := range card.Winners {
			if mine == winner {
				matches = true
				break
			}
		}
		if matches {
			if ret == 0 {
				ret = 1
			} else {
				ret *= 2
			}
		}
	}
	return ret
}
