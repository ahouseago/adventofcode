package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInputMaps() ([]int, [][][3]int) {
	var seeds []int
	maps := make([][][3]int, 7)
	for i, section := range strings.Split(strings.TrimSpace(input), "\n\n") {
		if i == 0 {
			for _, str := range strings.Split(strings.TrimPrefix(section, "seeds: "), " ") {
				parsed, _ := strconv.ParseInt(str, 10, 0)
				seeds = append(seeds, int(parsed))
			}
			continue
		}
		var mappings [][3]int
		for _, line := range strings.Split(section, "\n")[1:] {
			var lineNumbers [3]int
			for i, str := range strings.Split(line, " ") {
				parsed, _ := strconv.ParseInt(str, 10, 0)
				lineNumbers[i] = int(parsed)
			}
			mappings = append(mappings, lineNumbers)
		}
		maps[i-1] = mappings
	}
	return seeds, maps
}

func main() {
	seeds, maps := getInputMaps()
	var part1 int
	for _, seed := range seeds {
		value := getSeedValue(seed, maps)
		if part1 == 0 || part1 > value {
			part1 = value
		}
	}
	fmt.Println("Part 1", part1)

	// Time to brute-force the shit out of this...
	var part2 int
	for i := 0; i < len(seeds)-1; i += 2 {
		start, rng := seeds[i], seeds[i+1]
		for j := start; j < start+rng; j++ {
			value := getSeedValue(j, maps)
			if part2 == 0 || part2 > value {
				part2 = value
			}
		}
	}
	fmt.Println("Part 2", part2)
}

func getSeedValue(seed int, maps [][][3]int) int {
	value := seed
	for _, mapping := range maps {
		for _, destSrcRange := range mapping {
			dst, src, rng := destSrcRange[0], destSrcRange[1], destSrcRange[2]
			if value >= src && value < src+rng {
				value = value - src + dst
				break
			}
		}
	}
	return value
}
