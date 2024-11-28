package day01_2

import (
	"aoc23/utils"
	"fmt"
	"strconv"
	"strings"
)

func Trebuchet2() int {
	numLiteral := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	var totalCount = 0

	utils.MapFileLines("day01-2/input01-2.txt", func(line string) {
		var firstNumLiteralIndex = -1
		var lastNumLiteralIndex = -1
		var firstNum = 0
		var lastNum = 0
		for numL, num := range numLiteral {
			numIndex := strings.Index(line, numL)
			if numIndex != -1 {
				if firstNumLiteralIndex == -1 || numIndex < firstNumLiteralIndex {
					firstNumLiteralIndex = numIndex
					firstNum = num
				}
			}

			lastNumIndex := strings.LastIndex(line, numL)
			if lastNumIndex != -1 && lastNumIndex > lastNumLiteralIndex {
				lastNumLiteralIndex = lastNumIndex
				lastNum = num
			}
		}
		firstNumIndex := firstNumLiteralIndex
		for charIndex, char := range line {
			if num, numErr := strconv.Atoi(string(char)); numErr == nil {
				if firstNum == 0 || charIndex < firstNumIndex {
					firstNum = num
					firstNumIndex = charIndex
				}

				if charIndex > lastNumLiteralIndex {
					lastNum = num
				}
			}
		}
		if finalNumber, numErr := strconv.Atoi(fmt.Sprintf("%d%d", firstNum, lastNum)); numErr == nil {
			totalCount += finalNumber
		}
	})

	return totalCount
}
