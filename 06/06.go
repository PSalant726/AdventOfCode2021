package main

import "fmt"

func main() {
	school := lanternSchool(input)
	fmt.Printf("Part 1: %d\n", simulateLanternfish(80, school))
	fmt.Printf("Part 2: %d\n", simulateLanternfish(256, school))
}

// lanternSchool creates a slice of length 9, where each index
// contains the count of lanternfish in the school with that
// index's number of days remaining until reproduction.
func lanternSchool(lanternfish []int) []int {
	school := make([]int, 9)

	for _, fish := range lanternfish {
		school[fish]++
	}

	return school
}

func simulateLanternfish(days int, school []int) int {
	var day int
	for day < days {
		newFish := school[0]
		school = append(school[1:], newFish)
		school[6] += newFish
		day++
	}

	return sumSlice(school)
}

func sumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}

	return sum
}
