package main

import "fmt"

func main() {
	boards := make([]*Board, len(bingoBoards))
	for i, board := range bingoBoards {
		boards[i] = NewBoard(board)
	}

	fmt.Printf("Part 1: %d\n", part1(boards))
	// Part 1 modifies `boards`, but because of the
	// nature of the questions, this does not interfere
	// with the calculation of part 2's answer.
	fmt.Printf("Part 2: %d\n", part2(boards))
}

func part1(boards []*Board) int {
	for _, number := range drawn {
		for _, board := range boards {
			board.Mark(number)
			if board.Bingo {
				return board.Score(number)
			}
		}
	}

	return 0
}

func part2(boards []*Board) int {
	var lastDrawnNumber int
	var lastWinningBoard *Board

	for _, number := range drawn {
		for _, board := range boards {
			if board.Bingo {
				continue
			}

			board.Mark(number)
			if board.Bingo {
				lastWinningBoard = board
				lastDrawnNumber = number
			}
		}
	}

	return lastWinningBoard.Score(lastDrawnNumber)
}
