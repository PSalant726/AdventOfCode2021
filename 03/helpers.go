package main

import (
	"strconv"
	"strings"
)

func binToDec(number string) int64 {
	var result int64
	var bit int
	n := len(number) - 1

	for n >= 0 {
		if number[n] == '1' {
			result += (1 << (bit))
		}

		n = n - 1
		bit++
	}

	return result
}

func partitionByPositionVal(inputs []string, position int) ([]string, []string) {
	mostCommon := []string{}
	leastCommon := []string{}

	var positionTotal int
	for _, input := range inputs {
		inputInts := positionInts(input)
		positionTotal += inputInts[position]

		if inputInts[position] == 1 {
			mostCommon = append(mostCommon, input)
		} else {
			leastCommon = append(leastCommon, input)
		}
	}

	if positionTotal < len(inputs)/2 {
		mostCommon, leastCommon = leastCommon, mostCommon
	}

	return mostCommon, leastCommon
}

func positionInts(input string) []int {
	inputArr := strings.Split(input, "")
	retArr := make([]int, len(inputArr))

	for i, char := range inputArr {
		retArr[i], _ = strconv.Atoi(char)
	}

	return retArr
}
