package day03_2

import (
	"aoc23/utils"
	"strconv"
	"strings"
)

func isDigit(engineCell string) bool {
	if _, err := strconv.Atoi(engineCell); err == nil {
		return true
	}
	return false
}

type Number struct {
	row              int
	start            int
	end              int
	value            string
	isAdjacentToGear bool
	gearCoords       Point
}

type Point struct {
	x int
	y int
}

func isAdjacent(number *Number, engine [][]string) bool {
	i := number.row
	var shouldBeAdded = false
	for j := number.start; j <= number.end; j++ {
		if len(engine[i]) > j+1 && !isDigit(engine[i][j+1]) && engine[i][j+1] != "." {
			//fmt.Println("right")
			if engine[i][j+1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i, j + 1}
			}
			shouldBeAdded = true
		}

		if len(engine) > i+1 && len(engine[i]) > j+1 && !isDigit(engine[i+1][j+1]) && engine[i+1][j+1] != "." {

			if engine[i+1][j+1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i + 1, j + 1}
			}

			shouldBeAdded = true
		}

		if len(engine) > i+1 && !isDigit(engine[i+1][j]) && engine[i+1][j] != "." {
			//fmt.Println("bottom")
			if engine[i+1][j] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i + 1, j}
			}
			shouldBeAdded = true
		}

		if len(engine) > i+1 && j != 0 && !isDigit(engine[i+1][j-1]) && engine[i+1][j-1] != "." {
			//fmt.Println("bottom left")
			if engine[i+1][j-1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i + 1, j - 1}
			}
			shouldBeAdded = true
		}

		if j != 0 && !isDigit(engine[i][j-1]) && engine[i][j-1] != "." {
			//fmt.Println("left")
			if engine[i][j-1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i, j - 1}
			}
			shouldBeAdded = true
		}

		if i != 0 && j != 0 && !isDigit(engine[i-1][j-1]) && engine[i-1][j-1] != "." {
			//fmt.Println("upper left")
			if engine[i-1][j-1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i - 1, j - 1}
			}
			shouldBeAdded = true
		}

		if i != 0 && !isDigit(engine[i-1][j]) && engine[i-1][j] != "." {
			//fmt.Println("upper")
			if engine[i-1][j] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i - 1, j}
			}
			shouldBeAdded = true
		}

		if i != 0 && len(engine[i]) > j+1 && !isDigit(engine[i-1][j+1]) && engine[i-1][j+1] != "." {
			//fmt.Println("upper right")
			if engine[i-1][j+1] == "*" {
				number.isAdjacentToGear = true
				number.gearCoords = Point{i - 1, j + 1}
			}
			shouldBeAdded = true
		}
	}
	return shouldBeAdded
}
func GearRatios2() int {
	var engine [][]string
	utils.MapFileLines("day03-2/input03-2.txt", func(line string) {
		engine = append(engine, strings.Split(line, ""))
	})

	var engineParts []Number

	for i, engineRow := range engine {
		number := Number{row: i, start: -1, end: -1, value: ""}
		for j, engineCell := range engineRow {
			if isDigit(engineCell) {
				if number.value == "" {
					number.start = j
				}
				number.value += engineCell
			} else {
				if number.start != -1 {
					number.end = j - 1
					engineParts = append(engineParts, number)
					number = Number{row: i, start: -1, end: -1, value: ""}
				}
			}
		}
		if number.start != -1 {
			number.end = len(engine[i]) - 1
			engineParts = append(engineParts, number)
		}
	}

	var totalCount = 0
	for numberPartIndex := range engineParts {
		numberPart := &engineParts[numberPartIndex]
		if isAdjacent(numberPart, engine) {
			if num, err := strconv.Atoi(numberPart.value); err == nil {
				totalCount += num
			}
		}
	}

	gearsMap := map[Point][]Number{}
	for _, numberPart := range engineParts {
		if numberPart.isAdjacentToGear {
			if _, exists := gearsMap[numberPart.gearCoords]; exists {
				gearsMap[numberPart.gearCoords] = append(gearsMap[numberPart.gearCoords], numberPart)
			} else {
				gearsMap[numberPart.gearCoords] = []Number{numberPart}
			}
		}
	}

	//fmt.Println(gearsMap)
	var gearRatio = 0
	for _, gear := range gearsMap {
		if len(gear) == 2 {
			firstNum, _ := strconv.Atoi(gear[0].value)
			secondNum, _ := strconv.Atoi(gear[1].value)
			gearRatio += firstNum * secondNum
		}
	}

	return gearRatio
}
