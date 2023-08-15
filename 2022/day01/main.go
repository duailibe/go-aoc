package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	calories := []int{}
	current := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			calories = append(calories, current)
			current = 0
		} else {
			num, _ := strconv.Atoi(line)
			current += num
		}
	}
	slices.Sort(calories)
	slices.Reverse(calories)
	fmt.Println("Part 1:", calories[0])
	fmt.Println("Part 2:", sum(calories[0:3]))
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
