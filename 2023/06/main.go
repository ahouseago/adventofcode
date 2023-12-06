package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var re = regexp.MustCompile(`\d+`)

func main() {
	var parsed [][]int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var ns []int
		for _, str := range re.FindAllString(line, -1) {
			val, _ := strconv.ParseInt(str, 10, 0)
			ns = append(ns, int(val))
		}
		parsed = append(parsed, ns)
	}
	part1 := 1
	for i := 0; i < len(parsed[0]); i++ {
		time, dist := parsed[0][i], parsed[1][i]
		part1 *= numRecordBeating(time, dist)
	}
	fmt.Println("Part 1:", part1)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	time, _ := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", ""), 10, 0)
	dist, _ := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", ""), 10, 0)
	fmt.Println("Part 2:", numRecordBeating(int(time), int(dist)))
}

func numRecordBeating(time, dist int) int {
	for speed := 0; speed < time; speed++ {
		if (time-speed)*speed > dist {
			return time - speed - speed + 1
		}
	}
	return 0
}
