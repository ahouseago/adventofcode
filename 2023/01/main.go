package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

var re = regexp.MustCompile(`\d`)

var digits = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func main() {
	unparsed := []int{}
	for _, line := range getInputLines() {
		digitPositions := re.FindAllStringIndex(line, -1)
		if len(digitPositions) == 0 {
			continue
		}
		first, last := line[digitPositions[0][0]], line[digitPositions[len(digitPositions)-1][0]]
		val, err := strconv.ParseInt(string(first)+string(last), 10, 0)
		if err != nil {
			log.Fatal(err, line)
		}
		unparsed = append(unparsed, int(val))
	}
	sum := 0
	for _, val := range unparsed {
		sum += val
	}
	fmt.Println("Part 1:", sum)

	firstReString := "\\d"
	lastReString := "\\d"
	for key := range digits {
		firstReString += fmt.Sprintf("|(%s)", key)
		lastReString += fmt.Sprintf("|(%s)", Reverse(key))
	}

	firstRe := regexp.MustCompile(firstReString)
	lastRe := regexp.MustCompile(lastReString)

	sum2 := 0
	for _, line := range getInputLines() {
		first := firstRe.FindString(line)
		last := lastRe.FindString(Reverse(line))
		if unparsed, ok := digits[first]; ok {
			first = unparsed
		}
		if unparsed, ok := digits[Reverse(last)]; ok {
			last = unparsed
		}
		val, err := strconv.ParseInt(first+last, 10, 0)
		if err != nil {
			log.Fatal(err, line)
		}
		sum2 += int(val)
	}

	fmt.Println("Part 2:", sum2)
}

func Reverse(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
