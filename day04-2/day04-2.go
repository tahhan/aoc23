package day04_2

import (
	"aoc23/utils"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
	winnings       int
	copies         int
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

	return cardCount
}
func Scratchcards2() int {
	var cards []Card
	utils.MapFileLines("day04-2/input04-2.txt", func(line string) {
		cards = append(cards, parseCardLine(line))
	})

	//fmt.Println(cards)
	cardsMap := map[int]*Card{}
	for cardIndex := range cards {
		card := &cards[cardIndex]
		winningCardCount := winningCardCount(*card)
		card.winnings = winningCardCount
		cardsMap[card.id] = card
	}

	for cardId := range cardsMap {
		card := cardsMap[cardId]
		cardWinningsRec(cardsMap, card)
	}
	totalCount := 0

	for _, v := range cardsMap {
		totalCount += 1 + v.copies
	}

	//totalCount = 1 + cardWinningsRec(cardsMap, cardsMap[1])
	//fmt.Println(cardsMap)
	return totalCount
}

//func cardWinningsRec(cardsMap map[int]Card, card Card) int {
//	if card.winnings == 0 {
//		return 0
//	}
//
//	sum := 1
//	for i := 1; i < card.winnings; i++ {
//		sum += cardWinningsRec(cardsMap, cardsMap[card.id+i])
//	}
//
//	return sum
//}

func cardWinningsRec(cardsMap map[int]*Card, card *Card) {
	if card.winnings == 0 {
		return
	}

	if card.winnings != 0 {
		for i := 1; i <= card.winnings; i++ {
			copiedCard := cardsMap[card.id+i]
			copiedCard.copies++
			cardWinningsRec(cardsMap, copiedCard)
		}
	}
}
