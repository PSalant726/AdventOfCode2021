package main

type Submarine struct {
	Aim        int
	Depth      int
	Horizontal int
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

func (s *Submarine) AimDown(distance int) {
	s.Aim += distance
}

func (s *Submarine) AimUp(distance int) {
	s.Aim -= distance
}

func (s *Submarine) MoveForwardWithAim(distance int) {
	s.Horizontal += distance
	s.Depth += s.Aim * distance
}
