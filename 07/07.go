package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	idealPosition := median(input)

	var fuel int
	for _, position := range input {
		if position >= idealPosition {
			fuel += position - idealPosition
		} else {
			fuel += idealPosition - position
		}
	}

	return fuel
}

func median(slice []int) int {
	sort.Ints(slice)
	return slice[len(slice)/2]
}
