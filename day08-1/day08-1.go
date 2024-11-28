package day08_1

import (
	"aoc23/utils"
	"regexp"
)

type Node struct {
	leftElement  string
	rightElement string
}

func HauntedWasteland() int {

	nodes := map[string]Node{}
	var instructions []string
	utils.MapFileLines("day08-1/input08-1.txt", func(line string) {
		if r := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`); r.MatchString(line) {
			matches := r.FindAllStringSubmatch(line, -1)
			nodes[matches[0][1]] = Node{leftElement: matches[0][2], rightElement: matches[0][3]}
		} else if r := regexp.MustCompile(`^L|R`); r.MatchString(line) {
			for _, instruction := range line {
				instructions = append(instructions, string(instruction))
			}
		}
	})

	//fmt.Println(instructions)
	currentElement := "AAA"
	stepCounter := 0
	zzzFound := false
	for !zzzFound {
		for i := 0; i < len(instructions); i++ {
			if instructions[i] == "L" {
				currentElement = nodes[currentElement].leftElement
			} else {
				currentElement = nodes[currentElement].rightElement
			}

			//fmt.Println(instructions[i], currentElement)
			stepCounter++
			if currentElement == "ZZZ" {
				zzzFound = true
				break
			}
		}
	}
	return stepCounter
}
