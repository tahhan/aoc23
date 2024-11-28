package day04_1

import (
	"aoc23/utils"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func stringArrayToIntArray(stringArray []string) []int {
	var numArray []int
	for _, numString := range stringArray {
		if num, numErr := strconv.Atoi(numString); numErr == nil {
			numArray = append(numArray, num)
		}
	}
	return numArray
}

func parseCardLine(line string) Card {
	card := Card{}
	cardLineSplit := strings.Split(line, ":")
	cardNumberString, _ := strings.CutPrefix(cardLineSplit[0], "Card ")
	cardId, _ := strconv.Atoi(strings.Trim(cardNumberString, " "))
	card.id = cardId

	numbersLine := strings.Split(cardLineSplit[1], "|")
	card.winningNumbers = stringArrayToIntArray(strings.Split(numbersLine[0], " "))
	card.cardNumbers = stringArrayToIntArray(strings.Split(numbersLine[1], " "))

	return card
}

func winningCardCount(card Card) int {
	cardCount := 0
	for _, winningNum := range card.winningNumbers {
		for _, cardNum := range card.cardNumbers {
			if winningNum == cardNum {
				cardCount++
			}
		}
	}

	//fmt.Println(cardCount)
	return cardCount
}
func Scratchcards() int {
	var cards []Card
	utils.MapFileLines("day04-1/input04-1.txt", func(line string) {
		cards = append(cards, parseCardLine(line))
	})

	//fmt.Println(cards)
	totalCount := 0
	for _, card := range cards {
		winningCardCount := winningCardCount(card)
		if winningCardCount != 0 {
			totalCount += int(math.Pow(float64(2), float64(winningCardCount-1)))
		}
	}

	return totalCount
}
