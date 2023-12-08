package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := strings.Split(lines[0], "")
	nodes := make(map[string][]string)
	currentNodes := []string{} // For part 2
	endNodes := make(map[string]bool)
	for _, line := range lines[2:] {
		key, left, right := line[:3], line[7:10], line[12:15]
		nodes[key] = []string{left, right}
		if strings.HasSuffix(key, "Z") {
			endNodes[key] = true
		}
		if strings.HasSuffix(key, "A") {
			currentNodes = append(currentNodes, key)
		}
	}

	current, steps := "AAA", 0
	for i := 0; i <= len(instructions); i++ {
		if i == len(instructions) {
			i = -1 // Reset to the start (-1 for post-increment)
			continue
		}
		if current == "ZZZ" {
			break
		}
		if instructions[i] == "L" {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}
		steps++
	}

	fmt.Println("Part 1:", steps)

	cycleLength := make(map[string]int)

	steps = 0
	for i := 0; i <= len(instructions); i++ {
		if i == len(instructions) {
			i = -1 // Reset to the start (-1 for post-increment)
			continue
		}
		nodeIndex := 0
		if instructions[i] == "R" {
			nodeIndex = 1
		}
		nextNodes := make([]string, len(currentNodes))
		for i, node := range currentNodes {
			if isEndNode := endNodes[node]; isEndNode {
				cycleLength[node] = steps
			}
			nextNodes[i] = nodes[node][nodeIndex]
		}
		// Assume that this is enough, might not if the cycle alternates between L and R
		if len(cycleLength) == len(endNodes) {
			break
		}
		currentNodes = nextNodes
		steps++
	}
	var lengths []int
	for _, length := range cycleLength {
		lengths = append(lengths, length)
	}
	fmt.Println("Part 2:", lcm(lengths[0], lengths[1], lengths[2:]...)) // Assume > 2 nodes
}

func hcf(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / hcf(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}
