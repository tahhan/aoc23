package day06_2

import (
	"aoc23/utils"
	"regexp"
	"strconv"
	"strings"
)

func Boats2() int {
	var time int
	var distance int
	utils.MapFileLines("day06-2/input06-2.txt", func(line string) {
		if timesString, found := strings.CutPrefix(line, "Time:"); found {
			r := regexp.MustCompile(` *(\d+) *`)
			matches := r.FindAllStringSubmatch(timesString, -1)
			timeString := ""
			for _, v := range matches {
				timeString += v[1]
			}
			num, _ := strconv.Atoi(timeString)
			time = num
		} else if distancesString, found := strings.CutPrefix(line, "Distance:"); found {
			r := regexp.MustCompile(` *(\d+) *`)
			matches := r.FindAllStringSubmatch(distancesString, -1)
			distanceString := ""
			for _, v := range matches {
				distanceString += v[1]
			}
			num, _ := strconv.Atoi(distanceString)
			distance = num
		}
	})

	waysToWin := 0
	for i := 0; i <= time; i++ {
		calculatedDistance := i * (time - i)
		if calculatedDistance > distance {
			waysToWin++
		}
	}

	return waysToWin
}
