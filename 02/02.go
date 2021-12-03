package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	sub := &Submarine{}

	for _, step := range input {
		distance, _ := step[1].(int)
		switch step[0] {
		case "forward":
			sub.MoveForward(distance)
		case "down":
			sub.MoveDown(distance)
		case "up":
			sub.MoveUp(distance)
		}
	}

	return sub.Horizontal * sub.Depth
}

func part2() int {
	sub := &Submarine{}

	for _, step := range input {
		distance, _ := step[1].(int)
		switch step[0] {
		case "forward":
			sub.MoveForwardWithAim(distance)
		case "down":
			sub.AimDown(distance)
		case "up":
			sub.AimUp(distance)
		}
	}

	return sub.Horizontal * sub.Depth
}
