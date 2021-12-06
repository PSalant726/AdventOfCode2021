package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	for i := 0; i < 80; i++ {
		newLanternfish := []int{}
		for i, fish := range input {
			if fish == 0 {
				input[i] = 6
				newLanternfish = append(newLanternfish, 8)
			} else {
				input[i]--
			}
		}

		input = append(input, newLanternfish...)
	}

	return len(input)
}
