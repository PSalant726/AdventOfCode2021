package main

import "fmt"

func main() {
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

	fmt.Printf("Part 1: %d\n", part1(octopod))
}

func part1(octopod map[string]*Octopus) int {
	var flashes, step int
	for step < 100 {
		flashes += simulateStep(octopod)
		step++
	}

	return flashes
}

func simulateStep(octopod map[string]*Octopus) int {
	shouldFlash := map[string]*Octopus{}
	for _, octopus := range octopod {
		octopus.Energy++
		if octopus.Energy > 9 {
			shouldFlash[octopus.Id] = octopus
		}
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

	return flashes
}
