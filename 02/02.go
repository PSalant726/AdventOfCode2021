package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	sub := &Submarine{}

	for _, command := range input {
		switch command.Direction {
		case "forward":
			sub.MoveForward(command.Distance)
		case "down":
			sub.MoveDown(command.Distance)
		case "up":
			sub.MoveUp(command.Distance)
		}
	}

	return sub.Horizontal * sub.Depth
}

func part2() int {
	sub := &Submarine{}

	for _, command := range input {
		switch command.Direction {
		case "forward":
			sub.MoveForwardWithAim(command.Distance)
		case "down":
			sub.AimDown(command.Distance)
		case "up":
			sub.AimUp(command.Distance)
		}
	}

	return sub.Horizontal * sub.Depth
}
