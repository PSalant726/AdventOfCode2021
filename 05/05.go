package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
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

func part2() int {
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
		default:
			board.DrawDiagonal(startX, startY, endX, endY)
		}
	}

	return board.AtLeastTwoOverlaps
}
