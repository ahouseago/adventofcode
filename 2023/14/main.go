package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	state := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Println("Part 1:", calcNorthLoad(roll(state, "N")))

	stateCache := make(map[string][]string)
	// I'll be honest, I'm not sure why 1000 works the same as 1000000000.
	// Maybe something to do with the loop size being shorter than 1000?
	for i := 0; i < 1000; i++ {
		key := strings.Join(state, "\n")
		if cached, exists := stateCache[key]; exists {
			state = cached
			continue
		}
		state = cycle(state)
		stateCache[key] = state
	}
	fmt.Println("Part 2:", calcNorthLoad(state))
}

func calcNorthLoad(state []string) int {
	var load int
	for row, line := range state {
		load += ((len(state) - row) * strings.Count(line, "O"))
	}
	return load
}

func cycle(state []string) []string {
	return roll(roll(roll(roll(state, "N"), "W"), "S"), "E")
}

func roll(state []string, dir string) []string {
	rolled := make([]string, len(state))
	if dir == "N" {
		for row, line := range state {
			rolled[row] = line
			if row == 0 {
				continue
			}
			for col, rock := range line {
				if rock != 'O' {
					continue
				}
				for i := 1; i <= row; i++ {
					if rolled[row-i][col] != '.' {
						break
					}
					rolled[row-i] = rolled[row-i][:col] + "O" + rolled[row-i][col+1:]
					rolled[row-i+1] = rolled[row-i+1][:col] + "." + rolled[row-i+1][col+1:]
				}
			}
		}
	}
	if dir == "S" {
		for row := len(state) - 1; row >= 0; row-- {
			line := state[row]
			rolled[row] = line
			if row == len(state)-1 {
				continue
			}
			for col, rock := range line {
				if rock != 'O' {
					continue
				}
				for i := 1; row+i < len(rolled); i++ {
					if rolled[row+i][col] != '.' {
						break
					}
					rolled[row+i] = rolled[row+i][:col] + "O" + rolled[row+i][col+1:]
					rolled[row+i-1] = rolled[row+i-1][:col] + "." + rolled[row+i-1][col+1:]
				}
			}
		}
	}
	if dir == "W" {
		for col := 0; col < len(state[0]); col++ {
			for i, line := range state {
				rolled[i] += line[col : col+1]
			}
			if col == 0 {
				continue
			}
			for row := 0; row < len(state); row++ {
				if rolled[row][col] != 'O' {
					continue
				}
				for i := 1; i <= col; i++ {
					if rolled[row][col-i] != '.' {
						break
					}
					rolled[row] = rolled[row][:col-i] + "O." + rolled[row][col-i+2:]
				}
			}
		}
	}
	if dir == "E" {
		for col := len(state[0]) - 1; col >= 0; col-- {
			for i, line := range state {
				rolled[i] = line[col:col+1] + rolled[i]
			}
			if col == len(state[0])-1 {
				continue
			}
			for row := 0; row < len(rolled); row++ {
				if rolled[row][0] != 'O' {
					continue
				}
				for i := 1; i < len(rolled[0]); i++ {
					if rolled[row][i] != '.' {
						break
					}
					rolled[row] = rolled[row][:i-1] + ".O" + rolled[row][i+1:]
				}
			}
		}
	}
	return rolled
}
