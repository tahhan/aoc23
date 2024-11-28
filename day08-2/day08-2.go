package day08_2

import (
	"aoc23/utils"
	"math"
	"regexp"
	"strings"
)

type Node struct {
	leftElement  string
	rightElement string
}

func HauntedWasteland2() int {

	nodes := map[string]Node{}
	var instructions []string
	utils.MapFileLines("day08-2/input08-2.txt", func(line string) {
		if r := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`); r.MatchString(line) {
			matches := r.FindAllStringSubmatch(line, -1)
			nodes[matches[0][1]] = Node{leftElement: matches[0][2], rightElement: matches[0][3]}
		} else if r := regexp.MustCompile(`^L|R`); r.MatchString(line) {
			for _, instruction := range line {
				instructions = append(instructions, string(instruction))
			}
		}
	})

	//fmt.Println(instructions, nodes)
	var pathsCounters []int
	for element, _ := range nodes {
		if strings.HasSuffix(element, "A") {
			currentElement := element
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
					if strings.HasSuffix(currentElement, "Z") {
						zzzFound = true
						break
					}
				}
			}
			pathsCounters = append(pathsCounters, stepCounter)
		}
	}

	return lcmAll(pathsCounters)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
func lcmAll(numArr []int) int {
	acc := 1
	for _, v := range numArr {
		acc = lcm(acc, v)
	}
	return acc
}
