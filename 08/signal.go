package main

import "strings"

type Signal struct {
	Value            int
	ValueDigitCounts map[int]int
	digitsToPatterns map[int]string
	input            []string
	output           []string
	patternsToDigits map[string]int
}

func NewSignal(pattern string) *Signal {
	inputOutput := strings.Split(pattern, " | ")

	signal := &Signal{
		ValueDigitCounts: copyMap(digitCounts),
		digitsToPatterns: make(map[int]string, 10),
		input:            strings.Split(inputOutput[0], " "),
		output:           strings.Split(inputOutput[1], " "),
		patternsToDigits: make(map[string]int, 10),
	}

	signal.decode()
	return signal
}

func (s *Signal) decode() {
	for i, encoded := range s.input {
		encoded = alphabetize(encoded)
		s.input[i] = encoded

		switch len(encoded) {
		case 2:
			s.patternsToDigits[encoded] = 1
			s.digitsToPatterns[1] = encoded
		case 3:
			s.patternsToDigits[encoded] = 7
			s.digitsToPatterns[7] = encoded
		case 4:
			s.patternsToDigits[encoded] = 4
			s.digitsToPatterns[4] = encoded
		case 7:
			s.patternsToDigits[encoded] = 8
			s.digitsToPatterns[8] = encoded
		}
	}

	for _, encoded := range s.input {
		switch len(encoded) {
		case 5:
			if diff(encoded, s.digitsToPatterns[7]) == 2 {
				s.patternsToDigits[encoded] = 3
				s.digitsToPatterns[3] = encoded
			} else if intersections(encoded, s.digitsToPatterns[4]) == 3 {
				s.patternsToDigits[encoded] = 5
				s.digitsToPatterns[5] = encoded
			} else {
				s.patternsToDigits[encoded] = 2
				s.digitsToPatterns[2] = encoded
			}
		case 6:
			if diff(encoded, s.digitsToPatterns[4]) == 2 {
				s.patternsToDigits[encoded] = 9
				s.digitsToPatterns[9] = encoded
			} else if diff(encoded, s.digitsToPatterns[7]) == 3 {
				s.patternsToDigits[encoded] = 0
				s.digitsToPatterns[0] = encoded
			} else {
				s.patternsToDigits[encoded] = 6
				s.digitsToPatterns[6] = encoded
			}
		}
	}

	valueDigits := make([]int, 4)
	for i, encoded := range s.output {
		encoded = alphabetize(encoded)
		valueDigits[i] = s.patternsToDigits[encoded]
		s.ValueDigitCounts[s.patternsToDigits[encoded]]++
	}

	s.Value = joinIntSlice(valueDigits)
}
