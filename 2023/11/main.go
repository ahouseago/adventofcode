package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var out []string
	var galaxies [][]int
	galaxyColumns := make(map[int]bool)
	emptyRows := make(map[int]bool)
	for row, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if line == strings.Repeat(".", len(line)) {
			emptyRows[row] = true
		} else {
			for col, char := range strings.Split(line, "") {
				if char != "." {
					galaxyColumns[col] = true
					galaxies = append(galaxies, []int{len(out), col})
				}
			}
		}
		out = append(out, line)
	}
	part1 := 0
	part2 := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			part1 += distance(g1, g2, galaxyColumns, emptyRows, 1)
			part2 += distance(g1, g2, galaxyColumns, emptyRows, 999_999)
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func distance(p1, p2 []int, galaxyColumns, emptyRows map[int]bool, emptyDist int) int {
	y1, y2 := p1[0], p2[0]
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	x1, x2 := p1[1], p2[1]
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	dist := y2 - y1 + x2 - x1
	for i := y1 + 1; i < y2; i++ {
		if _, isEmpty := emptyRows[i]; isEmpty {
			dist += emptyDist
		}
	}
	for i := x1 + 1; i < x2; i++ {
		if _, hasGalaxy := galaxyColumns[i]; !hasGalaxy {
			dist += emptyDist
		}
	}
	return dist
}
