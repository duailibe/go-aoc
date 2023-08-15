package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	Part1()
	Part2()
}

func Part1() {
	points := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		opponent, mine, _ := strings.Cut(line, " ")

		opponentIndex := strings.Index("ABC", opponent)
		mineIndex := strings.Index("XYZ", mine)

		points += mineIndex + 1

		if mineIndex == opponentIndex {
			points += 3
		} else if winnerIndex := (opponentIndex + 1) % 3; mineIndex == winnerIndex {
			points += 6
		}
	}
	fmt.Println("Part 1:", points)
}

func Part2() {
	points := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		opponent, outcome, _ := strings.Cut(line, " ")

		opponentIndex := strings.Index("ABC", opponent)

		if outcome == "X" {
			points += mod(opponentIndex-1, 3) + 1
		} else if outcome == "Y" {
			points += 3 + opponentIndex + 1
		} else {
			points += 6 + (opponentIndex+1)%3 + 1
		}
	}
	fmt.Println("Part 2:", points)
}

func mod(a, b int) int {
	return (a%b + b) % b
}
