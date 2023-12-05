package helpers

import (
	"day04/part1/helpers"
)

func Solve(lines []string) int {
	copies := make(map[int]int)
	for _, line := range(lines) {
		card := helpers.ParseLine(line)
		getOrIncrement(copies, card.Id, 1)
		numWin := NumWinners(card)
		for i := 1; i <= numWin; i++ {
			getOrIncrement(copies, card.Id + i, copies[card.Id])
		}
	}
	ret := 0
	for _, v := range(copies) {
		ret += v
	}
	return ret
}

func getOrIncrement(copies map[int]int, id int, v int) {
	cop, ok := copies[id]
	if !ok {
		cop = 0
	}
	copies[id] = cop + v
}

func NumWinners(card helpers.Card) int {
	ret := 0
	for _, mine := range(card.Mine) {
		for _, winner := range(card.Winners) {
			if mine == winner {
				ret++
				break
			}
		}
	}
	return ret
}