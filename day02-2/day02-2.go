package day02_2

import (
	"aoc23/utils"
	"strconv"
	"strings"
)

type Game struct {
	gameId        int
	blueCounts    []int
	maxBlueCount  int
	redCounts     []int
	maxRedCount   int
	greenCounts   []int
	maxGreenCount int
}

func parseGameLine(gameLine string) Game {
	game := Game{}
	// Game 37: 16 blue, 5 green, 18 red; 3 blue, 14 green, 1 red; 4 blue, 3 green, 14 red; 12 green, 7 red, 15 blue; 15 green, 11 blue, 2 red; 8 blue, 13 green, 6 red
	gameLineSplit := strings.Split(gameLine, ":")

	gameIdString, _ := strings.CutPrefix(gameLineSplit[0], "Game ")
	gameId, _ := strconv.Atoi(gameIdString)
	game.gameId = gameId

	gamesLine := gameLineSplit[1]
	rounds := strings.Split(gamesLine, ";")
	for _, round := range rounds {
		colorsCount := strings.Split(strings.Trim(round, " "), ",")
		for _, colorCount := range colorsCount {

			if blueCountString, found := strings.CutSuffix(strings.Trim(colorCount, " "), " blue"); found {
				if blueCount, err := strconv.Atoi(blueCountString); err == nil {
					game.blueCounts = append(game.blueCounts, blueCount)
					game.maxBlueCount = max(game.maxBlueCount, blueCount)
				}
			}

			if greenCountString, found := strings.CutSuffix(strings.Trim(colorCount, " "), " green"); found {
				if greenCount, err := strconv.Atoi(greenCountString); err == nil {
					game.greenCounts = append(game.greenCounts, greenCount)
					game.maxGreenCount = max(game.maxGreenCount, greenCount)
				}
			}

			if redCountString, found := strings.CutSuffix(strings.Trim(colorCount, " "), " red"); found {
				if redCount, err := strconv.Atoi(redCountString); err == nil {
					game.redCounts = append(game.redCounts, redCount)
					game.maxRedCount = max(game.maxRedCount, redCount)
				}
			}
		}
	}

	return game
}

func gamePower(gameLine string) int {
	game := parseGameLine(gameLine)
	// 12 red cubes, 13 green cubes, and 14 blue cubes

	return game.maxGreenCount * game.maxBlueCount * game.maxRedCount
}

func CubeConundrum2() int {
	totalCount := 0
	utils.MapFileLines("day02-2/input02-2.txt", func(line string) {
		if gamePower := gamePower(line); gamePower != 0 {
			totalCount += gamePower
		}
	})

	return totalCount
}
