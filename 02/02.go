package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	sub := NewSubmarine()

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
