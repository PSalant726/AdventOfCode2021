package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
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
