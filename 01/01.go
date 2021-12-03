package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	var increases int
	for i, v := range input {
		if i > 0 && v > input[i-1] {
			increases++
		}
	}

	return increases
}

func part2() int {
	var increases int
	for i, v := range input {
		if i >= 3 {
			window1 := v + input[i-1] + input[i-2]
			window2 := window1 - v + input[i-3]

			if window1 > window2 {
				increases++
			}
		}
	}

	return increases
}
