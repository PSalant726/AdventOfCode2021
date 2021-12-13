package main

type Point struct {
	Height int
	X      int
	Y      int
}

func NewPoint(height, x, y int) Point {
	return Point{height, x, y}
}

func (p Point) Adjacents(board [][]int) []Point {
	points := []Point{}

	if p.Y > 0 {
		points = append(points, NewPoint(board[p.Y-1][p.X], p.X, p.Y-1))
	}

	if p.Y < len(board)-1 {
		points = append(points, NewPoint(board[p.Y+1][p.X], p.X, p.Y+1))
	}

	if p.X > 0 {
		points = append(points, NewPoint(board[p.Y][p.X-1], p.X-1, p.Y))
	}

	if p.X < len(board[p.Y])-1 {
		points = append(points, NewPoint(board[p.Y][p.X+1], p.X+1, p.Y))
	}

	return points
}

func (p Point) BasinSize(board [][]int, seen map[Point]struct{}) int {
	size := 1

	for _, nextPoint := range p.Adjacents(board) {
		if _, ok := seen[nextPoint]; ok {
			continue
		}

		if nextPoint.Height < 9 && nextPoint.Height >= p.Height {
			seen[nextPoint] = struct{}{}
			size += nextPoint.BasinSize(board, seen)
		}
	}

	return size
}

func (p Point) LowPoint(adjacents ...Point) bool {
	isLowPoint := true

	for _, adjacent := range adjacents {
		if p.Height >= adjacent.Height {
			isLowPoint = false
		}
	}

	return isLowPoint
}
