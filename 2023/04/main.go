package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

var re = regexp.MustCompile(`\d+`)

func main() {
	part1 := 0
	cards := getInputLines()
	winningCardCounts := make(map[int]int)
	for cardIndex, line := range cards {
		scratchcard := strings.Split(line, " | ")
		winningNumbers := re.FindAllString(strings.Split(scratchcard[0], ":")[1], -1)
		numbers := re.FindAllString(scratchcard[1], -1)
		winning := make(map[string]bool, len(winningNumbers))
		score := 0
		for _, number := range winningNumbers {
			winning[number] = true
		}
		for _, option := range numbers {
			if winning[option] {
				score++
			}
		}
		winningCardCounts[cardIndex] = score
		part1 += int(math.Pow(2, float64(score-1)))
	}

	fmt.Println("Part 1:", part1)

	part2 := 0
	numCards := make(map[int]int)
	for cardIndex := 0; cardIndex < len(winningCardCounts); cardIndex++ {
		winCount := winningCardCounts[cardIndex]
		for i := 0; i <= numCards[cardIndex]; i++ {
			part2++
			for j := cardIndex + 1; j <= cardIndex+winCount; j++ {
				numCards[j]++
			}
		}
	}

	fmt.Println("Part 2:", part2)
}
