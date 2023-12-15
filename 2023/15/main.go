package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type lens struct {
	label    string
	focalLen int
}

func main() {
	part1 := 0
	boxes := make([][]lens, 256)
	for _, str := range strings.Split(strings.TrimSpace(input), ",") {
		part1 += hash(str)
		split := strings.Split(str, "=")
		if len(split) == 1 {
			label := strings.TrimSuffix(split[0], "-")
			key := hash(label)
			for i, lens := range boxes[key] {
				if lens.label == label {
					boxes[key] = append(boxes[key][:i], boxes[key][i+1:]...)
					break
				}
			}
			continue
		}
		label := split[0]
		focalLen, _ := strconv.Atoi(split[1])
		key := hash(label)
		inBox := false
		for i, lens := range boxes[key] {
			if lens.label == label {
				boxes[key][i].focalLen = focalLen
				inBox = true
				break
			}
		}
		if !inBox {
			boxes[key] = append(boxes[key], lens{label, focalLen})
		}
	}
	part2 := 0
	for i, box := range boxes {
		for idx, lens := range box {
			part2 += (i + 1) * (idx + 1) * lens.focalLen
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func hash(str string) int {
	var out int
	for _, r := range str {
		out += int(r)
		out *= 17
		out %= 256
	}
	return out
}
