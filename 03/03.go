package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int64 {
	positionTotals := positionInts(input[0])
	for _, v := range input[1:] {
		inputTotals := positionInts(v)
		for i, val := range inputTotals {
			positionTotals[i] += val
		}
	}

	gammaArr := make([]string, len(positionTotals))
	epsilonArr := make([]string, len(positionTotals))

	for i, v := range positionTotals {
		gammaArr[i] = "0"
		epsilonArr[i] = "1"

		if v > len(input)/2 {
			gammaArr[i] = "1"
			epsilonArr[i] = "0"
		}
	}

	gamma := binToDec(strings.Join(gammaArr, ""))
	epsilon := binToDec(strings.Join(epsilonArr, ""))

	return gamma * epsilon
}

func part2() int64 {
	o2GeneratorVals, co2ScrubberVals := partitionByPositionVal(input, 0)

	position := 1
	for len(o2GeneratorVals) > 1 {
		o2GeneratorVals, _ = partitionByPositionVal(o2GeneratorVals, position)
		position++
	}

	position = 1
	for len(co2ScrubberVals) > 1 {
		_, co2ScrubberVals = partitionByPositionVal(co2ScrubberVals, position)
		position++
	}

	return binToDec(o2GeneratorVals[0]) * binToDec(co2ScrubberVals[0])
}
