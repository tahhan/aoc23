package day09_1

import (
	"aoc23/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MirageMaintenance() int {
	var numberArrs [][]int
	utils.MapFileLines("day09-1/input09-1.txt", func(line string) {
		numbersStrArr := strings.Split(line, " ")
		var numbers []int
		for _, numberStr := range numbersStrArr {
			if num, err := strconv.Atoi(numberStr); err == nil {
				numbers = append(numbers, num)
			}
		}
		numberArrs = append(numberArrs, numbers)
	})

	sum := 0
	for _, numArr := range numberArrs {
		if len(numArr) != 0 {
			sum += nextNumber(numArr)
		}
	}

	return sum
}

func differenceArr(numArr []int) []int {
	var diffArr []int
	for i := 1; i < len(numArr); i++ {
		diffArr = append(diffArr, int(math.Abs(float64(numArr[i]-numArr[i-1]))))
	}
	return diffArr
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}
func nextNumber(numArr []int) int {
	if len(numArr) < 3 {
		fmt.Println(numArr)
		return 0
	}
	if abs(numArr[len(numArr)-1]-numArr[len(numArr)-2]) == 0 {
		return numArr[len(numArr)-1]
	}
	return numArr[len(numArr)-1] + nextNumber(differenceArr(numArr))
}
