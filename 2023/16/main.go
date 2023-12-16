package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type node struct {
	pos [2]int
	dir string
}

func main() {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Println("Part 1:", energised(grid, node{[2]int{0, 0}, "R"}))

	highest := 0
	for i := 0; i < len(grid); i++ {
		highest = max(energised(grid, node{[2]int{i, 0}, "R"}), highest)
		highest = max(energised(grid, node{[2]int{i, len(grid[0]) - 1}, "L"}), highest)
	}
	for i := 0; i < len(grid[0]); i++ {
		highest = max(energised(grid, node{[2]int{0, i}, "D"}), highest)
		highest = max(energised(grid, node{[2]int{len(grid) - 1, i}, "U"}), highest)
	}
	fmt.Println("Part 2:", highest)
}

func energised(grid []string, start node) int {
	queue := []node{start}
	visited := make(map[[2]int][]string)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.pos[0] < 0 || current.pos[0] >= len(grid) ||
			current.pos[1] < 0 || current.pos[1] >= len(grid[0]) ||
			includes(visited[current.pos], current.dir) {
			continue
		}
		visited[current.pos] = append(visited[current.pos], current.dir)
		mirror := grid[current.pos[0]][current.pos[1]]
		switch current.dir {
		case "U":
			switch mirror {
			case '\\':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] - 1}, "L"})
			case '-':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] - 1}, "L"})
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] + 1}, "R"})
			case '/':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] + 1}, "R"})
			default:
				queue = append(queue, node{[2]int{current.pos[0] - 1, current.pos[1]}, "U"})
			}
		case "D":
			switch mirror {
			case '\\':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] + 1}, "R"})
			case '-':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] - 1}, "L"})
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] + 1}, "R"})
			case '/':
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] - 1}, "L"})
			default:
				queue = append(queue, node{[2]int{current.pos[0] + 1, current.pos[1]}, "D"})
			}
		case "L":
			switch mirror {
			case '\\':
				queue = append(queue, node{[2]int{current.pos[0] - 1, current.pos[1]}, "U"})
			case '|':
				queue = append(queue, node{[2]int{current.pos[0] - 1, current.pos[1]}, "U"})
				queue = append(queue, node{[2]int{current.pos[0] + 1, current.pos[1]}, "D"})
			case '/':
				queue = append(queue, node{[2]int{current.pos[0] + 1, current.pos[1]}, "D"})
			default:
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] - 1}, "L"})
			}
		case "R":
			switch mirror {
			case '\\':
				queue = append(queue, node{[2]int{current.pos[0] + 1, current.pos[1]}, "D"})
			case '|':
				queue = append(queue, node{[2]int{current.pos[0] - 1, current.pos[1]}, "U"})
				queue = append(queue, node{[2]int{current.pos[0] + 1, current.pos[1]}, "D"})
			case '/':
				queue = append(queue, node{[2]int{current.pos[0] - 1, current.pos[1]}, "U"})
			default:
				queue = append(queue, node{[2]int{current.pos[0], current.pos[1] + 1}, "R"})
			}
		}
	}
	return len(visited)
}

func includes(dirs []string, dir string) bool {
	for _, d := range dirs {
		if d == dir {
			return true
		}
	}
	return false
}
