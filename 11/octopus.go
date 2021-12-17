package main

import "fmt"

type Octopus struct {
	Energy int
	Id     string

	adjacents []*Octopus
	col       int
	row       int
}

func NewOctopus(row, col int) *Octopus {
	octopus := &Octopus{
		Energy: input[row][col],
		Id:     fmt.Sprintf("%d%d", row, col),

		adjacents: []*Octopus{},
		col:       col,
		row:       row,
	}

	return octopus
}

func (o *Octopus) Flash() []*Octopus {
	o.Energy = 0

	shouldFlash := []*Octopus{}
	for _, octopus := range o.adjacents {
		octopus.Energy++
		if octopus.Energy > 9 {
			shouldFlash = append(shouldFlash, octopus)
		}
	}

	return shouldFlash
}

func (o *Octopus) PopulateAdjacents() {
	isInTopRow := o.row == 0
	isInBottomRow := o.row == 9
	isInLeftColumn := o.col == 0
	isInRightColumn := o.col == 9

	if !isInTopRow {
		// Top Middle
		o.adjacents = append(o.adjacents, NewOctopus(o.row-1, o.col))

		if !isInLeftColumn {
			// Top Left
			o.adjacents = append(o.adjacents, NewOctopus(o.row-1, o.col-1))
		}

		if !isInRightColumn {
			// Top Right
			o.adjacents = append(o.adjacents, NewOctopus(o.row-1, o.col+1))
		}
	}

	if !isInLeftColumn {
		// Left
		o.adjacents = append(o.adjacents, NewOctopus(o.row, o.col-1))
	}

	if !isInRightColumn {
		// Right
		o.adjacents = append(o.adjacents, NewOctopus(o.row, o.col+1))
	}

	if !isInBottomRow {
		// Bottom Middle
		o.adjacents = append(o.adjacents, NewOctopus(o.row+1, o.col))

		if !isInLeftColumn {
			// Bottom Left
			o.adjacents = append(o.adjacents, NewOctopus(o.row+1, o.col-1))
		}

		if !isInRightColumn {
			// Bottom Right
			o.adjacents = append(o.adjacents, NewOctopus(o.row+1, o.col+1))
		}
	}
}
