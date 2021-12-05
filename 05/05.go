package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	board := NewBoard()

	for _, line := range input {
		startX := line[0][0]
		startY := line[0][1]
		endX := line[1][0]
		endY := line[1][1]

		switch {
		case startX == endX:
			board.DrawVertical(startX, startY, endY)
		case startY == endY:
			board.DrawHorizontal(startX, startY, endX)
		}
	}

	return board.AtLeastTwoOverlaps
}
