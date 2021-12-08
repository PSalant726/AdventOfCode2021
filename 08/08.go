package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	counts := copyMap(digitCounts)
	for _, pattern := range input {
		signal := NewSignal(pattern)
		signal.DecodeUniquesInOutput()

		for k, v := range signal.ValueDigitCounts {
			counts[k] += v
		}
	}

	return counts[1] + counts[4] + counts[7] + counts[8]
}
