package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	idealPosition := median(input)

	var fuel int
	for _, position := range input {
		fuel += difference(position, idealPosition)
	}

	return fuel
}

func part2() int {
	idealPosition := mean(input)

	var fuel int
	for _, position := range input {
		fuel += triangularNumber(difference(position, idealPosition))
	}

	return fuel
}

func difference(a, b int) int {
	if a >= b {
		return a - b
	}

	return b - a
}

func mean(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}

	return sum / len(slice)
}

func median(slice []int) int {
	sort.Ints(slice)
	return slice[len(slice)/2]
}

// https://en.wikipedia.org/wiki/Triangular_number
func triangularNumber(num int) int {
	return num * (num + 1) / 2
}
