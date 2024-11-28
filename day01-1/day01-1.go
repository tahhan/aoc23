package day01_1

import (
	"aoc23/utils"
	"fmt"
	"strconv"
)

func Trebuchet() int {

	var totalCount = 0
	utils.MapFileLines("day01-1/input01-1.txt", func(line string) {
		var firstNum = 0
		var lastNum = 0
		for _, char := range line {
			if num, numErr := strconv.Atoi(string(char)); numErr == nil {
				if firstNum == 0 {
					firstNum = num
				}
				lastNum = num
			}
		}
		if finalNumber, numErr := strconv.Atoi(fmt.Sprintf("%d%d", firstNum, lastNum)); numErr == nil {
			totalCount += finalNumber
		}
	})
	return totalCount
}
