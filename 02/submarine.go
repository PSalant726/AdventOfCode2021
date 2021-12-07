package main

type Submarine struct {
	Depth      int
	Horizontal int
	aim        int
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
	s.aim += distance
}

func (s *Submarine) AimUp(distance int) {
	s.aim -= distance
}

func (s *Submarine) MoveForwardWithAim(distance int) {
	s.MoveForward(distance)
	s.Depth += s.aim * distance
}
