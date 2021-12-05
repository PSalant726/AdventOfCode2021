package main

type Board struct {
	Bingo   bool
	Columns [][]int
	Rows    [][]int
	Values  map[int]struct{}
}

func NewBoard(numbers [][]int) *Board {
	board := &Board{
		Columns: make([][]int, 5),
		Rows:    numbers,
		Values:  make(map[int]struct{}, 25),
	}

	for _, row := range numbers {
		for j, num := range row {
			board.Values[num] = struct{}{}
			board.Columns[j] = append(board.Columns[j], num)
		}
	}

	return board
}

func (b *Board) Mark(drawn int) {
	delete(b.Values, drawn)

	for i, row := range b.Rows {
		for j, number := range row {
			if number == drawn {
				b.Rows[i] = remove(row, j)

				if len(b.Rows[i]) == 0 {
					b.Bingo = true
				}

				break
			}
		}
	}

	for i, column := range b.Columns {
		for j, number := range column {
			if number == drawn {
				b.Columns[i] = remove(column, j)

				if len(b.Columns[i]) == 0 {
					b.Bingo = true
				}

				break
			}
		}
	}
}

func (b Board) Score(drawn int) int {
	var sum int
	for number := range b.Values {
		sum += number
	}

	return sum * drawn
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
