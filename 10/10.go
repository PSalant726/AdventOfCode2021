package main

import (
	"fmt"
	"sort"
	"strings"
)

var openers = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var closers = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func main() {
	syntaxErrorScore, middleCompletionScore := parseNavSubsystem()

	fmt.Printf("Part 1: %d\n", syntaxErrorScore)
	fmt.Printf("Part 2: %d\n", middleCompletionScore)
}

func parseNavSubsystem() (int, int) {
	var completionScores []int
	var syntaxErrorScore int
	var syntaxErrorScores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	for _, line := range input {
		var charStack []string
		var corrupted bool

		chars := strings.Split(line, "")
		for _, char := range chars {
			if _, isOpener := closers[char]; isOpener {
				charStack = append(charStack, char)
			} else if charStack[len(charStack)-1] == openers[char] {
				charStack = charStack[:len(charStack)-1]
			} else {
				syntaxErrorScore += syntaxErrorScores[char]
				corrupted = true
				break
			}
		}

		if !corrupted {
			completionScores = append(completionScores, completionScore(charStack))
		}
	}

	sort.Ints(completionScores)

	return syntaxErrorScore, completionScores[len(completionScores)/2]
}

func completionScore(remainingStack []string) int {
	var completionScore int
	var scores = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	for i := len(remainingStack) - 1; i >= 0; i-- {
		completionChar := closers[remainingStack[i]]
		completionScore = completionScore*5 + scores[completionChar]
	}

	return completionScore
}
