package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type hand struct {
	cards         map[string]int
	withJokers    map[string]int
	bid           int
	originalOrder []string
}

func main() {
	var hands []hand
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, " ")
		cards := strings.Split(split[0], "")
		bid, _ := strconv.Atoi(split[1])
		cardMap, withJokers := make(map[string]int, 5), make(map[string]int, 5)
		jokers := 0
		for _, card := range cards {
			cardMap[card]++
			if card == "J" {
				jokers++
			} else {
				withJokers[card]++
			}
		}
		maxWithoutJokers, maxValue := "", 0
		for card, count := range withJokers {
			if count > maxValue {
				maxValue = count
				maxWithoutJokers = card
			}
		}
		withJokers[maxWithoutJokers] += jokers
		hands = append(hands, hand{cardMap, withJokers, bid, cards})
	}

	sort.Slice(hands, func(i, j int) bool {
		return sortHandTypes(hands[i].cards, hands[j].cards, hands[i].originalOrder, hands[j].originalOrder)
	})

	part1 := 0
	for i, h := range hands {
		part1 += (len(hands) - i) * h.bid
	}

	fmt.Println("Part 1:", part1)

	handTypes["J"] = 1
	sort.Slice(hands, func(i, j int) bool {
		return sortHandTypes(hands[i].withJokers, hands[j].withJokers, hands[i].originalOrder, hands[j].originalOrder)
	})

	part2 := 0
	for i, h := range hands {
		part2 += (len(hands) - i) * h.bid
	}

	fmt.Println("Part 2:", part2)

}
func sortHandTypes(a, b map[string]int, orderA, orderB []string) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}

	if len(a) == 2 {
		aFourOfKind := false
		for _, count := range a {
			if count == 4 {
				aFourOfKind = true
			}
		}
		bFourOfKind := false
		for _, count := range b {
			if count == 4 {
				bFourOfKind = true
			}
		}
		if aFourOfKind != bFourOfKind {
			return aFourOfKind
		}
		return fallback(orderA, orderB)
	}
	if len(a) == 3 {
		aThreeOfKind := false
		for _, count := range a {
			if count == 3 {
				aThreeOfKind = true
			}
		}
		bThreeOfKind := false
		for _, count := range b {
			if count == 3 {
				bThreeOfKind = true
			}
		}
		if aThreeOfKind != bThreeOfKind {
			return aThreeOfKind
		}
		return fallback(orderA, orderB)
	}
	return fallback(orderA, orderB)
}

var handTypes = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func fallback(as, bs []string) bool {
	for i := 0; i < len(as); i++ {
		a, b := as[i], bs[i]
		if a == b {
			continue
		}
		return handTypes[a] > handTypes[b]
	}
	return false
}
