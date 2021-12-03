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

func positionInts(input string) []int {
	inputArr := strings.Split(input, "")
	retArr := make([]int, len(inputArr))

	for i, char := range inputArr {
		retArr[i], _ = strconv.Atoi(char)
	}

	return retArr
}
