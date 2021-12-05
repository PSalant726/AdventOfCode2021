package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	boards := make([]*Board, len(bingoBoards))
	for i, board := range bingoBoards {
		boards[i] = NewBoard(board)
	}

	for _, number := range drawn {
		for _, board := range boards {
			if _, ok := board.Values[number]; ok {
				board.Mark(number)
				if board.Bingo {
					return board.Score(number)
				}
			}
		}
	}

	return 0
}
