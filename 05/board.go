package main

type Board struct {
	AtLeastTwoOverlaps int
	Rows               [][]int
}

func NewBoard() *Board {
	board := &Board{Rows: make([][]int, 999)}

	for i := range board.Rows {
		board.Rows[i] = make([]int, 999)
	}

	return board
}

func (b *Board) DrawDiagonal(startX, startY, endX, endY int) {
	y := startY
	switch {
	case startX <= endX && startY <= endY:
		for i := startX; i <= endX; i++ {
			b.plot(i, y)
			y++
		}
	case startX <= endX && startY > endY:
		for i := startX; i <= endX; i++ {
			b.plot(i, y)
			y--
		}
	case startX > endX && startY <= endY:
		for i := startX; i >= endX; i-- {
			b.plot(i, y)
			y++
		}
	default:
		for i := startX; i >= endX; i-- {
			b.plot(i, y)
			y--
		}
	}
}

func (b *Board) DrawHorizontal(startX, startY, endX int) {
	if startX <= endX {
		for i := startX; i <= endX; i++ {
			b.plot(i, startY)
		}
	} else {
		for i := startX; i >= endX; i-- {
			b.plot(i, startY)
		}
	}
}

func (b *Board) DrawVertical(startX, startY, endY int) {
	if startY <= endY {
		for i := startY; i <= endY; i++ {
			b.plot(startX, i)
		}
	} else {
		for i := startY; i >= endY; i-- {
			b.plot(startX, i)
		}
	}
}

func (b *Board) plot(x, y int) {
	if b.Rows[y][x] >= 2 {
		b.AtLeastTwoOverlaps--
	}

	b.Rows[y][x]++

	if b.Rows[y][x] >= 2 {
		b.AtLeastTwoOverlaps++
	}
}
