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

func Solve(lines []string) int {
	numCards := make(map[int]int)
	for _, line := range lines {
		card := parse(line)
		ws := cardWinners(card)
		incrementBy(numCards, card.Id, 1)
		for i := 0; i < ws; i++ {
			incrementBy(numCards, card.Id + 1 + i, numCards[card.Id])
		}
	}
	return sum(numCards)
}

func sum(m map[int]int) int {
	s := 0
	for _, n := range m {
		s += n
	}
	return s
}

func incrementBy(m map[int]int, n int, b int) {
	if _, ok := m[n]; ok {
		m[n] += b
	} else {
		m[n] = b
	}
}

func parse(line string) Card {
	fields := strings.Fields(line)
	id, _ := strconv.Atoi(strings.TrimSuffix(fields[1], ":"))
	w := []int{}
	m := []int{}
	idx := 2
	for ; idx < len(fields) && fields[idx] != "|"; idx++ {
		n, _ := strconv.Atoi(fields[idx])
		w = append(w, n)
	}
	idx++
	for ; idx < len(fields); idx++ {
		n, _ := strconv.Atoi(fields[idx])
		m = append(m, n)
	}
	return Card{id, w, m}
}

func cardWinners(card Card) int {
	winners := 0
	for _, n := range card.Mine {
		if contains(card.Winners, n) {
			winners++
		}
	}
	return winners
}

func contains(arr []int, n int) bool {
	for _, i := range arr {
		if i == n {
			return true
		}
	}
	return false
}
