package main

import (
	"fmt"
	"sort"
)

func main() {
	riskSum, lowPoints := part1()

	fmt.Printf("Part 1: %d\n", riskSum)
	fmt.Printf("Part 2: %d\n", part2(lowPoints))
}

func part1() (int, []Point) {
	var totalRisk int
	var lowPoints []Point

	for i, row := range input {
		for j, height := range row {
			point := NewPoint(height, j, i)
			adjacents := point.Adjacents(input)

			if point.LowPoint(adjacents...) {
				totalRisk += height + 1
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return totalRisk, lowPoints
}

func part2(lowPoints []Point) int {
	basins := make([]int, len(lowPoints))

	for i, lowPoint := range lowPoints {
		basins[i] = lowPoint.BasinSize(input, map[Point]struct{}{})
	}

	sort.Ints(basins)
	basins = basins[len(basins)-3:]

	return basins[0] * basins[1] * basins[2]
}
