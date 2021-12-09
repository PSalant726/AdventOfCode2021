package main

type Point struct {
	Exists bool
	Height int
}

func NewPoint(height int) Point {
	return Point{
		Height: height,
		Exists: true,
	}
}

func (p Point) LowPoint(adjacents ...Point) bool {
	isLowPoint := true

	for _, adjacent := range adjacents {
		if adjacent.Exists && p.Height >= adjacent.Height {
			isLowPoint = false
		}
	}

	return isLowPoint
}
