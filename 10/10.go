package main

import (
	"fmt"
	"strings"
)

var openersToClosers = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var closersToOpeners = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var closerScores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

func part1() int {
	var syntaxErrorScore int

	for _, line := range input {
		var charStack []string
		chars := strings.Split(line, "")

		for _, char := range chars {
			if _, charIsOpener := openersToClosers[char]; charIsOpener {
				charStack = append(charStack, char)
			} else if score, charIsCloser := closerScores[char]; charIsCloser {
				if charStack[len(charStack)-1] == closersToOpeners[char] {
					charStack = charStack[:len(charStack)-1]
				} else {
					syntaxErrorScore += score
					break
				}
			}
		}
	}

	return syntaxErrorScore
}
