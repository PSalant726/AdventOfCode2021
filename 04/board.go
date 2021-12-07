package main

type Board struct {
	Bingo   bool
	columns rowsOrColumns
	rows    rowsOrColumns
	values  map[int]struct{}
}

type rowsOrColumns struct {
	values [][]int
}

func NewBoard(numbers [][]int) *Board {
	boardLength := len(numbers)
	board := &Board{
		columns: rowsOrColumns{make([][]int, boardLength)},
		rows:    rowsOrColumns{numbers},
		values:  make(map[int]struct{}, boardLength*boardLength),
	}

	for _, row := range numbers {
		for j, num := range row {
			board.values[num] = struct{}{}
			board.columns.values[j] = append(board.columns.values[j], num)
		}
	}

	return board
}

func (b *Board) Mark(drawn int) {
	if _, ok := b.values[drawn]; !ok {
		return
	}

	delete(b.values, drawn)
	b.mark(drawn, b.rows)
	b.mark(drawn, b.columns)
}

func (b Board) Score(drawn int) int {
	var sum int
	for number := range b.values {
		sum += number
	}

	return sum * drawn
}

func (b *Board) mark(drawn int, rowsOrColumns rowsOrColumns) {
	for i, rowOrColumn := range rowsOrColumns.values {
		for j, number := range rowOrColumn {
			if number == drawn {
				rowsOrColumns.values[i] = remove(rowOrColumn, j)

				if len(rowsOrColumns.values[i]) == 0 {
					b.Bingo = true
				}

				return
			}
		}
	}
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
