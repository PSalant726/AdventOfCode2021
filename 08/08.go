package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	counts := copyMap(digitCounts)
	for _, pattern := range input {
		for k, v := range NewSignal(pattern).ValueDigitCounts {
			counts[k] += v
		}
	}

	return counts[1] + counts[4] + counts[7] + counts[8]
}

func part2() int {
	var outputTotal int
	for _, pattern := range input {
		outputTotal += NewSignal(pattern).Value
	}

	return outputTotal
}
