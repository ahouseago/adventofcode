package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1, part2 := 0, 0
	for _, line := range lines {
		var ns []int
		for _, str := range strings.Split(line, " ") {
			val, _ := strconv.Atoi(str)
			ns = append(ns, val)
		}
		part1 += nextVal(ns, false)
		part2 += nextVal(ns, true)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func nextVal(vals []int, prev bool) int {
	out := make([]int, len(vals)-1)
	set := make(map[int]bool)
	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]
		out[i-1] = diff
		set[diff] = true
	}
	next := out[0]
	if len(set) > 1 {
		next = nextVal(out, prev)
	}
	if prev {
		return vals[0] - next
	}
	return vals[len(vals)-1] + next
}
