package main

import (
	"sort"
	"strings"
)

var digitCounts = map[int]int{
	0: 0,
	1: 0,
	2: 0,
	3: 0,
	4: 0,
	5: 0,
	6: 0,
	7: 0,
	8: 0,
	9: 0,
}

func alphabetize(str string) string {
	sorted := sort.StringSlice(strings.Split(str, ""))
	sorted.Sort()
	return strings.Join([]string(sorted), "")
}

func copyMap(origin map[int]int) map[int]int {
	destination := make(map[int]int, len(origin))
	for k, v := range origin {
		destination[k] = v
	}

	return destination
}

// diff returns the number of
// characters in a that are not in b.
func diff(a, b string) int {
	for _, c := range b {
		a = strings.Replace(a, string(c), "", 1)
	}

	return len(a)
}

// intersections returns the number of
// characters in a that are also in b.
func intersections(a, b string) int {
	var count int
	for _, char := range b {
		if strings.Contains(a, string(char)) {
			count++
		}
	}

	return count
}

// joinIntSlice creates a single integer
// represented by the integers in s.
func joinIntSlice(s []int) int {
	var result int
	place := 1

	for i := len(s) - 1; i >= 0; i-- {
		result += s[i] * place
		place *= 10
	}

	return result
}
