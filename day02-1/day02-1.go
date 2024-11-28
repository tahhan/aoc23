package day02_1

import (
	"aoc23/utils"
	"strconv"
	"strings"
)

type Game struct {
	gameId      int
	blueCounts  []int
	redCounts   []int
	greenCounts []int
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
				}
			}

			if greenCountString, found := strings.CutSuffix(strings.Trim(colorCount, " "), " green"); found {
				if greenCount, err := strconv.Atoi(greenCountString); err == nil {
					game.greenCounts = append(game.greenCounts, greenCount)
				}
			}

			if redCountString, found := strings.CutSuffix(strings.Trim(colorCount, " "), " red"); found {
				if redCount, err := strconv.Atoi(redCountString); err == nil {
					game.redCounts = append(game.redCounts, redCount)
				}
			}
		}
	}

	return game
}

func isGamePossible(gameLine string) int {
	game := parseGameLine(gameLine)
	// 12 red cubes, 13 green cubes, and 14 blue cubes

	isPossible := true
	for _, redCount := range game.redCounts {
		if redCount > 12 {
			isPossible = false
		}
	}

	for _, greenCount := range game.greenCounts {
		if greenCount > 13 {
			isPossible = false
		}
	}

	for _, blueCount := range game.blueCounts {
		if blueCount > 14 {
			isPossible = false
		}
	}

	if isPossible {
		return game.gameId
	}
	return 0
}

func CubeConundrum() int {
	totalCount := 0
	utils.MapFileLines("day02-1/input02-1.txt", func(line string) {
		if gameId := isGamePossible(line); gameId != 0 {
			totalCount += gameId
		}
	})
	return totalCount
}
