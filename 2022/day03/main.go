package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	input = strings.TrimSpace(input)
	Part1()
	Part2()
}

func Part1() {
	priorities := 0
	for _, line := range strings.Split(input, "\n") {
		first := makeRucksack(line[0 : len(line)/2])
		second := makeRucksack(line[len(line)/2:])

		for value, priority := range first {
			if rucksackContains(second, value) {
				priorities += priority
			}
		}
	}
	fmt.Println("Part 1:", priorities)
}

func Part2() {
	rucksacks := []rucksack{}
	for _, line := range strings.Split(input, "\n") {
		rucksacks = append(rucksacks, makeRucksack(line))
	}

	priorities := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for value, priority := range rucksacks[i] {
			if rucksackContains(rucksacks[i+1], value) && rucksackContains(rucksacks[i+2], value) {
				priorities += priority
			}
		}
	}
	fmt.Println("Part 2:", priorities)
}

type rucksack = map[rune]int

func makeRucksack(str string) map[rune]int {
	rucksack := map[rune]int{}
	for _, rune := range str {
		rucksack[rune] = strings.IndexRune(alphabet, rune) + 1
	}
	return rucksack
}

func rucksackContains(rucksack rucksack, item rune) bool {
	_, exists := rucksack[item]
	return exists
}
