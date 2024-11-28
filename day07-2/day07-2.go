package day07_2

import (
	"aoc23/utils"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	originalHand  string
	cards         []int
	replacedCards []int
	bid           int
	strength      int
}

func highestNumber(cards []int) int {
	highestNumber := cards[0]
	for _, card := range cards {
		highestNumber = max(highestNumber, card)
	}
	return highestNumber
}

func parseCards(hand string) ([]int, int, []int) {
	var cards []int
	for _, cardString := range hand {
		if cardNumber, err := strconv.Atoi(string(cardString)); err == nil {
			cards = append(cards, cardNumber)
		} else {
			cardNumber := 0
			switch string(cardString) {
			case "T":
				cardNumber = 10
			case "J":
				cardNumber = 1
			case "Q":
				cardNumber = 12
			case "K":
				cardNumber = 13
			case "A":
				cardNumber = 14
			default:
				cardNumber = 0
			}
			cards = append(cards, cardNumber)
		}
	}

	cardsMap := map[int]int{}
	for _, card := range cards {
		if cardNumberCount, found := cardsMap[card]; found {
			cardsMap[card] = cardNumberCount + 1
		} else {
			cardsMap[card] = 1
		}
	}

	originalHighestCount := 0
	origianlHighestCountNumber := 0
	for cardNumber, cardCount := range cardsMap {
		if cardNumber != 1 && cardCount > originalHighestCount {
			originalHighestCount = cardCount
			origianlHighestCountNumber = cardNumber
		}

	}

	var cardsReplaced []int
	if slices.Contains(cards, 1) {
		//highestNumber := highestNumber(cards)
		// replace it with the highest card number
		for _, card := range cards {
			if card == 1 {
				cardsReplaced = append(cardsReplaced, origianlHighestCountNumber)
			} else {
				cardsReplaced = append(cardsReplaced, card)
			}
		}
	} else {
		cardsReplaced = cards
	}

	cardsMapReplaced := map[int]int{}
	for _, card := range cardsReplaced {
		if cardNumberCount, found := cardsMapReplaced[card]; found {
			cardsMapReplaced[card] = cardNumberCount + 1
		} else {
			cardsMapReplaced[card] = 1
		}
	}

	highestCount := 0
	for _, cardCount := range cardsMapReplaced {
		highestCount = max(highestCount, cardCount)
	}

	strength := 0
	// calculate strength
	switch len(cardsMapReplaced) {
	case 5:
		strength = 1
		//fmt.Println(hand, "High card")
	case 4:
		strength = 2
		//fmt.Println("One pair", hand)
	case 3:
		if highestCount == 3 {
			// Three of a kind
			strength = 4
			//fmt.Println("Three of a kind", hand)
		} else {
			// Two pair
			strength = 3
			//fmt.Println("Two pair", hand)
		}
	case 2:
		if highestCount == 4 {
			// Four of a kind
			strength = 6
			//fmt.Println("Four of a kind", hand)
		} else {
			// Full house
			strength = 5
			//fmt.Println("Full house", hand)
		}
	case 1:
		strength = 7
		//fmt.Println("Five of a kind", hand)
	}

	return cards, strength, cardsReplaced
}

func CamelCards2() int {
	var hands []Hand
	utils.MapFileLines("day07-2/moreTestInput07-2.txt", func(line string) {
		handLine := strings.Split(line, " ")
		cards, strength, replacedCards := parseCards(handLine[0])
		bid, _ := strconv.Atoi(handLine[1])
		hand := Hand{originalHand: handLine[0], cards: cards, replacedCards: replacedCards, bid: bid, strength: strength}
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

	totalCount := 0
	for handIndex, hand := range hands {
		//fmt.Println(hand)
		totalCount += hand.bid * (handIndex + 1)
	}

	return totalCount
}
