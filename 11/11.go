package main

import "fmt"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	octopod := map[string]*Octopus{}
	for row := range input {
		for col := range input[row] {
			octopus := NewOctopus(row, col)
			octopod[octopus.Id] = octopus
		}
	}

	for _, octopus := range octopod {
		octopus.PopulateAdjacents()
	}

	var flashes, step int
	for step < 100 {
		flashed, _ := simulateStep(octopod)
		flashes += flashed
		step++
	}

	return flashes
}

func part2() int {
	octopod := map[string]*Octopus{}
	for row := range input {
		for col := range input[row] {
			octopus := NewOctopus(row, col)
			octopod[octopus.Id] = octopus
		}
	}

	for _, octopus := range octopod {
		octopus.PopulateAdjacents()
	}

	var step int
	for {
		if _, allFlashed := simulateStep(octopod); allFlashed {
			return step
		}

		step++
	}
}

func simulateStep(octopod map[string]*Octopus) (int, bool) {
	shouldFlash := map[string]*Octopus{}
	for _, octopus := range octopod {
		octopus.Energy++
		if octopus.Energy > 9 {
			shouldFlash[octopus.Id] = octopus
		}
	}

	var allFlashed bool
	if len(shouldFlash) == len(octopod) {
		allFlashed = true
	}

	var flashes int
	for len(shouldFlash) > 0 {
		for _, octopus := range shouldFlash {
			newFlashes := octopus.Flash()
			for _, new := range newFlashes {
				shouldFlash[new.Id] = new
			}

			delete(shouldFlash, octopus.Id)
			flashes++
		}
	}

	return flashes, allFlashed
}
