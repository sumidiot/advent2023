package helpers

import (
	"slices"
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	return SolveInput(ParseLines(lines))
}

func SolveInput(i *Inputs) int {
	slices.SortFunc(i.HandBids, HandOrder)
	winnings := 0
	for idx, hb := range i.HandBids {
		winnings += hb.Bid * (idx + 1)
	}
	return winnings
}

/**
 * Return negative if hb1 < hb2, 0 if equal, positive if hb1 > hb2
 */
func HandOrder(hb1, hb2 HandBid) int {
	ht1 := GetHandType(hb1.Hand)
	ht2 := GetHandType(hb2.Hand)
	if ht1 != ht2 {
		// If hb1 < hb2 is determined by HandType, hb2 has stronger hand, so lower int value. thus ht1 - ht2 > 0
		return -(int(ht1) - int(ht2))
	} else {
		// same type, compare individual cards
		// if hb1 < hb2 is determined by a card's value, hb2 has a stronger card, so lower int value
		// CoPilot mostly wrote this, I changed the negatives (hopefully correctly!)
		if hb1.Hand[0] == hb2.Hand[0] {
			if hb1.Hand[1] == hb2.Hand[1] {
				if hb1.Hand[2] == hb2.Hand[2] {
					if hb1.Hand[3] == hb2.Hand[3] {
						if hb1.Hand[4] == hb2.Hand[4] {
							return 0
						} else {
							return -(int(hb1.Hand[4]) - int(hb2.Hand[4]))
						}
					} else {
						return -(int(hb1.Hand[3]) - int(hb2.Hand[3]))
					}
				} else {
					return -(int(hb1.Hand[2]) - int(hb2.Hand[2]))
				}
			} else {
				return -(int(hb1.Hand[1]) - int(hb2.Hand[1]))
			}
		} else {
			return -(int(hb1.Hand[0]) - int(hb2.Hand[0]))
		}
	}
}

/**
 * This enum captures the standard twelve playing cards
 */
type Card int

const (
	Ace Card = iota
	King
	Queen
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	Jack
)

func (c Card) String() string {
	return [...]string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}[c]
}

func (c Card) Value() int {
	return [...]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}[c]
}

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (ht HandType) String() string {
	return [...]string{"FiveOfAKind", "FourOfAKind", "FullHouse", "ThreeOfAKind", "TwoPair", "OnePair", "HighCard"}[ht]
}

func GetHandType(h []Card) HandType {
	cc := CardCounts(h)
	if cc[Jack] > 0 {
		jacks := cc[Jack]
		delete(cc, Jack)
		maxCount := 0
		maxCard := Ace
		for c, v := range cc {
			if v > maxCount {
				maxCount = v
				maxCard = c
			}
		}
		cc[maxCard] += jacks
	}
	switch len(cc) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, v := range cc {
			if v == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, v := range cc {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	default:
		panic("WTF")
	}
}

func CardCounts(h []Card) map[Card]int {
	cc := make(map[Card]int)
	for _, c := range h {
		cc[c]++
	}
	return cc
}

func ToCard(c byte) Card {
	switch c {
	case 'A':
		return Ace
	case 'K':
		return King
	case 'Q':
		return Queen
	case 'T':
		return Ten
	case '9':
		return Nine
	case '8':
		return Eight
	case '7':
		return Seven
	case '6':
		return Six
	case '5':
		return Five
	case '4':
		return Four
	case '3':
		return Three
	case '2':
		return Two
	case 'J':
		return Jack
	default:
		panic("WTF")
	}
}

type HandBid struct {
	Hand []Card
	Bid  int
}

type Inputs struct {
	HandBids []HandBid
}

func ParseLines(lines []string) *Inputs {
	hbs := make([]HandBid, len(lines))
	for idx, line := range lines {
		hbs[idx] = ParseLine(line)
	}
	return &Inputs{hbs}
}

func ParseLine(line string) HandBid {
	b, _ := strconv.Atoi(strings.Fields(line)[1])
	h := make([]Card, 5)
	for idx := 0; idx < 5; idx++ {
		h[idx] = ToCard(line[idx])
	}
	return HandBid{h, b}
}
