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
	input = strings.TrimSpace(input)
	contained := 0
	overlaps := 0
	for _, line := range strings.Split(input, "\n") {
		first, second, _ := strings.Cut(line, ",")
		firstSection := parseSection(first)
		secondSection := parseSection(second)
		if firstSection.Contains(secondSection) || secondSection.Contains(firstSection) {
			contained++
		}
		if firstSection.Overlaps(secondSection) {
			overlaps++
		}
	}
	fmt.Println("Part 1:", contained)
	fmt.Println("Part 2:", overlaps)
}

type section struct {
	start, end int
}

func parseSection(str string) section {
	startStr, endStr, _ := strings.Cut(str, "-")
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)
	return section{start, end}
}

func (s section) Contains(other section) bool {
	return s.start <= other.start && s.end >= other.end
}

func (s section) Overlaps(other section) bool {
	return min(s.end, other.end) >= max(s.start, other.start)
}
