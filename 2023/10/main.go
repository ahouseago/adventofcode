package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type direction string

var (
	L direction = "L"
	D direction = "D"
	U direction = "U"
	R direction = "R"
)

var dirToVec = map[direction][]int{
	U: {-1, 0},
	D: {1, 0},
	L: {0, -1},
	R: {0, 1},
}

var lookup = map[string]map[direction]direction{
	"|": map[direction]direction{D: D, U: U},
	"-": map[direction]direction{L: L, R: R},
	"7": map[direction]direction{R: D, U: L},
	"F": map[direction]direction{L: D, U: R},
	"J": map[direction]direction{R: U, D: L},
	"L": map[direction]direction{L: U, D: R},
}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var start []int
	var grid [][]string
	for i, line := range lines {
		grid = append(grid, make([]string, len(strings.TrimSpace(line))))
		for j, char := range strings.Split(strings.TrimSpace(line), "") {
			grid[i][j] = char
			if char == "S" {
				start = append(start, i, j)
			}
		}
	}

	// var left, up, right, down bool
	var lastDir direction
	var next []int
	var vector []int
	up := grid[start[0]-1][start[1]]
	switch up {
	case "|", "F", "7":
		lastDir = U
		vector = dirToVec[lookup[up][lastDir]]
	}
	down := grid[start[0]+1][start[1]]
	switch down {
	case "|", "L", "J":
		lastDir = D
		vector = dirToVec[lookup[down][lastDir]]
	}
	left := grid[start[0]][start[1]-1]
	switch left {
	case "-", "L", "F":
		lastDir = L
		vector = dirToVec[lookup[left][lastDir]]
	}
	right := grid[start[0]][start[1]+1]
	switch right {
	case "-", "J", "7":
		lastDir = R
		vector = dirToVec[lookup[right][lastDir]]
	}
	next = []int{start[0] + vector[0], start[1] + vector[1]}
	steps := 0
	loop := make(map[int]map[int]string)
	if _, ok := loop[start[0]]; !ok {
		loop[start[0]] = make(map[int]string)
	}
	loop[start[0]][start[1]] = "S"
	if _, ok := loop[next[0]]; !ok {
		loop[next[0]] = make(map[int]string)
	}
	loop[next[0]][next[1]] = grid[next[0]][next[1]]
	for next[0] != start[0] || next[1] != start[1] {
		steps++
		dir := lookup[grid[next[0]][next[1]]][lastDir]
		vector = dirToVec[dir]
		next = []int{next[0] + vector[0], next[1] + vector[1]}
		lastDir = dir
		if _, ok := loop[next[0]]; !ok {
			loop[next[0]] = make(map[int]string)
		}
		loop[next[0]][next[1]] = grid[next[0]][next[1]]
	}

	fmt.Println("Part 1:", (steps+1)/2)

	for i, line := range grid {
		out := ""
		for j := range line {
			if str, inLoop := loop[i][j]; inLoop {
				if str == "S" {
					out += "┘" // I checked this manually
				} else {
					out += toBox[str]
				}
			} else {
				out += "."
			}
		}
		fmt.Println(out) // I used a paint fill on this output...
	}
}

var toBox = map[string]string{
	"|": "│",
	"-": "─",
	"7": "┐",
	"J": "┘",
	"F": "┌",
	"L": "└",
}
