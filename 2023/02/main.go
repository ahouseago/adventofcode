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

var re = regexp.MustCompile("\\d+")
var colourRe = regexp.MustCompile("(\\d+) (red|green|blue)")

var validCubes = map[string]int64{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	sum := 0
	sumPowers := 0
	for _, line := range getInputLines() {
		game := re.FindString(line)
		gameNumber, err := strconv.ParseInt(game, 10, 0)
		if err != nil {
			log.Fatal(err, game, line)
		}
		cubes := colourRe.FindAllString(line, -1)
		valid := true
		maxes := make(map[string]int64)
		for _, count := range cubes {
			number, err := strconv.ParseInt(re.FindString(count), 10, 0)
			if err != nil {
				log.Fatal(err, line, count)
			}
			for colour, maxNumber := range validCubes {
				if strings.HasSuffix(count, colour) && number > maxNumber {
					valid = false
				}
				if strings.HasSuffix(count, colour) && maxes[colour] < number {
					maxes[colour] = number
				}
			}
		}
		if valid {
			sum += int(gameNumber)
		}
		power := 1
		for _, val := range maxes {
			power *= int(val)
		}
		sumPowers += power
	}
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sumPowers)
}
