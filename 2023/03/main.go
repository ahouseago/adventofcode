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

var re = regexp.MustCompile(`\d+`)
var symbolsRe = regexp.MustCompile(`[^\d\.]`)

type span struct {
	value int
	row   int
	span  []int
}

type symbol struct {
	isGear bool
	pos    int
}

func main() {
	digitPositionsPerLine := make(map[int][]span)
	symbolsPerRow := make(map[int][]symbol)
	for row, line := range getInputLines() {
		digitPositions := re.FindAllStringIndex(line, -1)
		var spans []span
		for _, pos := range digitPositions {
			str := line[pos[0]:pos[len(pos)-1]]
			value, err := strconv.ParseInt(str, 10, 0)
			if err != nil {
				log.Fatal(err, row, line, str)
			}
			spans = append(spans, span{int(value), row, pos})
		}
		digitPositionsPerLine[row] = spans
		symbolPositions := getInitialValues(symbolsRe.FindAllStringIndex(line, -1))
		var symbols []symbol
		for _, pos := range symbolPositions {
			symbols = append(symbols, symbol{line[pos] == '*', pos})
		}
		symbolsPerRow[row] = symbols
	}

	part2 := 0
	validNumbers := make(map[string]int)
	for row, symbols := range symbolsPerRow {
		above, sameLine, below := digitPositionsPerLine[row-1], digitPositionsPerLine[row], digitPositionsPerLine[row+1]
		for _, symbol := range symbols {
			var adjacentNumbers []int
			for _, numberPos := range append(above, below...) {
				if symbol.pos >= numberPos.span[0]-1 && symbol.pos <= numberPos.span[len(numberPos.span)-1] {
					validNumbers[fmt.Sprintf("%d:%d", numberPos.row, numberPos.span[0])] = numberPos.value
					if symbol.isGear {
						adjacentNumbers = append(adjacentNumbers, numberPos.value)
					}
				}
			}
			for _, numberPos := range sameLine {
				if symbol.pos == numberPos.span[0]-1 || symbol.pos == numberPos.span[len(numberPos.span)-1] {
					validNumbers[fmt.Sprintf("%d:%d", numberPos.row, numberPos.span[0])] = numberPos.value
					if symbol.isGear {
						adjacentNumbers = append(adjacentNumbers, numberPos.value)
					}
				}
			}
			if len(adjacentNumbers) == 2 {
				part2 += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	part1 := 0
	for _, num := range validNumbers {
		part1 += num
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func getInitialValues(vals [][]int) []int {
	out := make([]int, len(vals))
	for i, val := range vals {
		out[i] = val[0]
	}
	return out
}
