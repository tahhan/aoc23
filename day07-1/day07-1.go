package day07_1

import (
	"aoc23/utils"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []int
	bid      int
	strength int
}

func parseCards(hand string) ([]int, int) {
	cardsMap := map[int]int{}
	var cards []int
	strength := 0
	for _, cardString := range hand {
		if cardNumber, err := strconv.Atoi(string(cardString)); err == nil {
			if cardNumberCount, found := cardsMap[cardNumber]; found {
				cardsMap[cardNumber] = cardNumberCount + 1
			} else {
				cardsMap[cardNumber] = 1
			}
			cards = append(cards, cardNumber)
		} else {
			cardNumber := 0
			switch string(cardString) {
			case "T":
				cardNumber = 10
			case "J":
				cardNumber = 11
			case "Q":
				cardNumber = 12
			case "K":
				cardNumber = 13
			case "A":
				cardNumber = 14
			default:
				cardNumber = 0
			}

			if cardNumberCount, found := cardsMap[cardNumber]; found {
				cardsMap[cardNumber] = cardNumberCount + 1
			} else {
				cardsMap[cardNumber] = 1
			}
			cards = append(cards, cardNumber)
		}
	}

	highestCount := 0
	for _, cardCount := range cardsMap {
		highestCount = max(highestCount, cardCount)
	}

	// calculate strength
	switch len(cardsMap) {
	case 5:
		strength = 1
	case 4:
		strength = 2
	case 3:
		if highestCount == 3 {
			// Three of a kind
			strength = 4
		} else {
			// Two pair
			strength = 3
		}
	case 2:
		if highestCount == 4 {
			// Four of a kind
			strength = 6
		} else {
			// Full house
			strength = 5
		}
	case 1:
		strength = 7
	}

	return cards, strength
}

func CamelCards() int {
	var hands []Hand
	utils.MapFileLines("day07-1/input07-1.txt", func(line string) {
		handLine := strings.Split(line, " ")
		cards, strength := parseCards(handLine[0])
		bid, _ := strconv.Atoi(handLine[1])
		hand := Hand{cards: cards, bid: bid, strength: strength}
		hands = append(hands, hand)
	})

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.strength == b.strength {
			for cardIndex, card := range a.cards {
				if card == b.cards[cardIndex] {
					continue
				}

				if card > b.cards[cardIndex] {
					return 1
				}
				return -1
			}
		}

		return cmp.Compare(a.strength, b.strength)
	})

	//fmt.Println(hands)
	totalCount := 0
	for handIndex, hand := range hands {
		totalCount += hand.bid * (handIndex + 1)
	}

	return totalCount
}
