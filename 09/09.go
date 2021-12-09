package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	var totalRisk int

	for i, row := range input {
		for j, height := range row {
			var up, down, left, right Point
			point := NewPoint(height)

			if i > 0 {
				up = NewPoint(input[i-1][j])
			}

			if i < len(input)-1 {
				down = NewPoint(input[i+1][j])
			}

			if j > 0 {
				left = NewPoint(row[j-1])
			}

			if j < len(row)-1 {
				right = NewPoint(row[j+1])
			}

			if point.LowPoint(up, down, left, right) {
				totalRisk += height + 1
			}
		}
	}

	return totalRisk
}
