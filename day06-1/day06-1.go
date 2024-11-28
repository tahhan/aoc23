package day06_1

import (
	"aoc23/utils"
	"regexp"
	"strconv"
	"strings"
)

func Boats() int {
	var times []int
	var distances []int
	utils.MapFileLines("day06-1/input06-1.txt", func(line string) {
		if timesString, found := strings.CutPrefix(line, "Time:"); found {
			r := regexp.MustCompile(` *(\d+) *`)
			matches := r.FindAllStringSubmatch(timesString, -1)
			for _, v := range matches {
				if num, err := strconv.Atoi(v[1]); err == nil {
					times = append(times, num)
				}
			}
		} else if distancesString, found := strings.CutPrefix(line, "Distance:"); found {
			r := regexp.MustCompile(` *(\d+) *`)
			matches := r.FindAllStringSubmatch(distancesString, -1)
			for _, v := range matches {
				if num, err := strconv.Atoi(v[1]); err == nil {
					distances = append(distances, num)
				}
			}
		}
	})

	totalMulti := 1
	for timeIndex, time := range times {
		distance := distances[timeIndex]

		waysTowin := 0
		for i := 0; i <= time; i++ {
			calculatedDistance := i * (time - i)
			if calculatedDistance > distance {
				waysTowin++
			}
		}
		totalMulti *= waysTowin

	}

	return totalMulti
}
