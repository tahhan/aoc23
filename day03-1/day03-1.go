package day03_1

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
	row   int
	start int
	end   int
	value string
}

//type Point struct {
//	x int
//	y int
//}

func isAdjacent(number Number, engine [][]string) bool {
	i := number.row
	var shouldBeAdded = false
	for j := number.start; j <= number.end; j++ {
		if len(engine[i]) > j+1 && !isDigit(engine[i][j+1]) && engine[i][j+1] != "." {
			//fmt.Println("right")
			shouldBeAdded = true
		}

		if len(engine) > i+1 && len(engine[i]) > j+1 && !isDigit(engine[i+1][j+1]) && engine[i+1][j+1] != "." {
			//fmt.Println("bottom right")
			shouldBeAdded = true
		}

		if len(engine) > i+1 && !isDigit(engine[i+1][j]) && engine[i+1][j] != "." {
			//fmt.Println("bottom")
			shouldBeAdded = true
		}

		if len(engine) > i+1 && j != 0 && !isDigit(engine[i+1][j-1]) && engine[i+1][j-1] != "." {
			//fmt.Println("bottom left")
			shouldBeAdded = true
		}

		if j != 0 && !isDigit(engine[i][j-1]) && engine[i][j-1] != "." {
			//fmt.Println("left")
			shouldBeAdded = true
		}

		if i != 0 && j != 0 && !isDigit(engine[i-1][j-1]) && engine[i-1][j-1] != "." {
			//fmt.Println("upper left")
			shouldBeAdded = true
		}

		if i != 0 && !isDigit(engine[i-1][j]) && engine[i-1][j] != "." {
			//fmt.Println("upper")
			shouldBeAdded = true
		}

		if i != 0 && len(engine[i]) > j+1 && !isDigit(engine[i-1][j+1]) && engine[i-1][j+1] != "." {
			//fmt.Println("upper right")
			shouldBeAdded = true
		}
	}
	//fmt.Println(number, shouldBeAdded)
	return shouldBeAdded
}
func GearRatios() int {
	var engine [][]string
	utils.MapFileLines("day03-1/input03-1.txt", func(line string) {
		engine = append(engine, strings.Split(line, ""))
	})

	var engineParts []Number

	//engine := [][]string{{"#", "6", "1", "7"},
	//	{".", ".", ".", "#"}}

	//fmt.Println(engine)
	//var adjacencyIndices []Point
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
			//
			////fmt.Println(engineCell, i, j)
			//shouldBeAdded := false
			//if isDigit(engineCell) {
			//
			//	if len(engine[i]) > j+1 && !isDigit(engine[i][j+1]) && engine[i][j+1] != "." {
			//		//fmt.Println("right")
			//		shouldBeAdded = true
			//	}
			//
			//	if len(engine) > i+1 && len(engine[i]) > j+1 && !isDigit(engine[i+1][j+1]) && engine[i+1][j+1] != "." {
			//		//fmt.Println("bottom right")
			//		shouldBeAdded = true
			//	}
			//
			//	if len(engine) > i+1 && !isDigit(engine[i+1][j]) && engine[i+1][j] != "." {
			//		//fmt.Println("bottom")
			//		shouldBeAdded = true
			//	}
			//
			//	if len(engine) > i+1 && j != 0 && !isDigit(engine[i+1][j-1]) && engine[i+1][j-1] != "." {
			//		//fmt.Println("bottom left")
			//		shouldBeAdded = true
			//	}
			//
			//	if j != 0 && !isDigit(engine[i][j-1]) && engine[i][j-1] != "." {
			//		//fmt.Println("left")
			//		shouldBeAdded = true
			//	}
			//
			//	if i != 0 && j != 0 && !isDigit(engine[i-1][j-1]) && engine[i-1][j-1] != "." {
			//		//fmt.Println("upper left")
			//		shouldBeAdded = true
			//	}
			//
			//	if i != 0 && !isDigit(engine[i-1][j]) && engine[i-1][j] != "." {
			//		//fmt.Println("upper")
			//		shouldBeAdded = true
			//	}
			//
			//	if i != 0 && len(engine[i]) > j+1 && !isDigit(engine[i-1][j+1]) && engine[i-1][j+1] != "." {
			//		//fmt.Println("upper right")
			//		shouldBeAdded = true
			//	}
			//}
			//
			//if shouldBeAdded {
			//	if index := slices.IndexFunc(adjacencyIndices, func(point Point) bool {
			//		return point.x == i && point.y > j-2 && point.y < j+2
			//	}); index == -1 {
			//		adjacencyIndices = append(adjacencyIndices, Point{i, j})
			//	}
			//
			//}
		}
		if number.start != -1 {
			number.end = len(engine[i]) - 1
			engineParts = append(engineParts, number)
		}
	}

	var totalCount = 0
	for _, numberPart := range engineParts {
		if isAdjacent(numberPart, engine) {
			if num, err := strconv.Atoi(numberPart.value); err == nil {
				totalCount += num
			}
		}
	}

	//totalCount := 0
	//for _, point := range adjacencyIndices {
	//	i := point.x
	//	j := point.y
	//	currntNumber := engine[i][j]
	//	firstDigit := ""
	//	secondDigit := ""
	//	thirdDigit := ""
	//	fourthDigit := ""
	//
	//	if j > 1 && isDigit(engine[i][j-2]) && isDigit(engine[i][j-1]) {
	//		firstDigit = engine[i][j-2]
	//	}
	//
	//	if j != 0 && isDigit(engine[i][j-1]) {
	//		secondDigit = engine[i][j-1]
	//	}
	//
	//	if len(engine[i]) > j+1 && isDigit(engine[i][j+1]) {
	//		thirdDigit = engine[i][j+1]
	//	}
	//
	//	if len(engine[i]) > j+2 && isDigit(engine[i][j+2]) && isDigit(engine[i][j+1]) {
	//		fourthDigit = engine[i][j+2]
	//	}
	//
	//	if num, err := strconv.Atoi(firstDigit + secondDigit + currntNumber + thirdDigit + fourthDigit); err == nil {
	//		fmt.Println(num)
	//		totalCount += num
	//	}
	//}

	return totalCount
}
