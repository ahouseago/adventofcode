package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1 := 0
	var nonSmudged []int
	for gridNo, grid := range strings.Split(strings.TrimSpace(input), "\n\n") {
		rows := strings.Split(grid, "\n")
		lineOfSymmetry := 0
		for i := 1; i < len(rows); i++ {
			rowGood := true
			if rows[i] != rows[i-1] {
				continue
			}
			for offset := 2; i-offset >= 0 && offset+i <= len(rows); offset++ {
				if rows[i-offset] != rows[i+offset-1] {
					rowGood = false
					break
				}
			}
			if rowGood {
				lineOfSymmetry = i * 100
				nonSmudged = append(nonSmudged, lineOfSymmetry)
				break
			}
		}
		if lineOfSymmetry != 0 {
			part1 += lineOfSymmetry
			continue
		}
		width := len(rows[0])
		for i := 1; i < width; i++ {
			columnGood := true
			for _, row := range rows {
				if row[max(0, i-(width-i)):i] != Reverse(row[i:min(2*i, width)]) {
					columnGood = false
					break
				}
			}
			if columnGood {
				lineOfSymmetry = i
				nonSmudged = append(nonSmudged, lineOfSymmetry)
			}
		}
		if lineOfSymmetry == 0 {
			fmt.Printf("\nGrid %d: no line found:\n", gridNo)
			fmt.Println(grid)
		}
		part1 += lineOfSymmetry
	}

	// Just do it all again but check it's not the same as before
	part2 := 0
	for gridNo, grid := range strings.Split(strings.TrimSpace(input), "\n\n") {
		originalSymmetry := nonSmudged[gridNo]
		rows := strings.Split(grid, "\n")
		lineOfSymmetry := 0
		for i := 1; i < len(rows); i++ {
			rowGood := true
			rowDifferences := 0
			for j := 0; j < len(rows[i]); j++ {
				if rows[i][j] != rows[i-1][j] {
					rowDifferences++
				}
			}
			if rowDifferences > 1 {
				continue
			}
			for offset := 2; i-offset >= 0 && offset+i <= len(rows); offset++ {
				offsetRowDifferences := 0
				for j := 0; j < len(rows[i]); j++ {
					if rows[i-offset][j] != rows[i+offset-1][j] {
						offsetRowDifferences++
					}
				}
				if rowDifferences+offsetRowDifferences > 1 {
					rowGood = false
					break
				}
			}
			if rowGood && originalSymmetry != i*100 {
				lineOfSymmetry = i * 100
				break
			}
		}
		if lineOfSymmetry != 0 {
			part2 += lineOfSymmetry
			continue
		}
		width := len(rows[0])
		for i := 1; i < width; i++ {
			columnGood := true
			colDiffs := 0
			for _, row := range rows {
				left, right := row[max(0, i-(width-i)):i], Reverse(row[i:min(2*i, width)])
				for j := 0; j < len(left); j++ {
					if left[j] != right[j] {
						colDiffs++
					}
				}
				if colDiffs > 1 {
					columnGood = false
					break
				}
			}
			if columnGood && originalSymmetry != i {
				lineOfSymmetry = i
			}
		}
		if lineOfSymmetry == 0 {
			fmt.Printf("\nGrid %d: no line found:\n", gridNo)
			fmt.Println(grid)
		}
		part2 += lineOfSymmetry
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
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
