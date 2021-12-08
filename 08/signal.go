package main

import "strings"

var segmentCounts = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

type Signal struct {
	ValueDigitCounts map[int]int
	output           []string
}

func NewSignal(pattern string) *Signal {
	inputOutput := strings.Split(pattern, " | ")

	return &Signal{
		ValueDigitCounts: copyMap(digitCounts),
		output:           strings.Split(inputOutput[1], " "), // Always has length 4
	}
}

func (s *Signal) DecodeUniquesInOutput() {
	for _, encoded := range s.output {
		if values := segmentCounts[len(encoded)]; len(values) == 1 {
			s.ValueDigitCounts[values[0]]++
		}
	}
}
