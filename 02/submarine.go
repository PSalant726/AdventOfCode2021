package main

type Submarine struct {
	Depth      int
	Horizontal int
}

func NewSubmarine() *Submarine {
	return &Submarine{0, 0}
}

func (s *Submarine) MoveForward(distance int) {
	s.Horizontal += distance
}

func (s *Submarine) MoveDown(distance int) {
	s.Depth += distance
}

func (s *Submarine) MoveUp(distance int) {
	s.Depth -= distance
}
