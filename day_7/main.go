package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

type CamelCardHand struct {
	cardsCounts map[rune]int
	cards       []rune
	bid         int
}

type CamelCards struct {
	hands []CamelCardHand
}

func parseCamelCards(lines []string) CamelCards {
	hands := make([]CamelCardHand, len(lines))
	for i, line := range lines {
		hands[i] = parseCamelCardHand(line)
	}
	return CamelCards{hands}
}

func parseCamelCardHand(line string) CamelCardHand {
	parts := strings.Fields(line)
	cardsPart := parts[0]
	bidPart := parts[1]
	bid, err := strconv.Atoi(bidPart)
	panicIfError(err)

	cardsCount := make(map[rune]int)
	cards := make([]rune, len(cardsPart))
	for i, card := range cardsPart {
		cardsCount[card] += 1
		cards[i] = card
	}

	return CamelCardHand{cardsCount, cards, bid}
}

func (camelCardHand *CamelCardHand) Score() int {
	mostPairs := 0
	secondMostPairs := 0

	// Part two
	jokers := camelCardHand.cardsCounts['J']

	for key, count := range camelCardHand.cardsCounts {
		// Part two
		if key == 'J' {
			continue
		}

		if count > mostPairs {
			secondMostPairs = mostPairs
			mostPairs = count
		} else if count > secondMostPairs {
			secondMostPairs = count
		}
	}

	// return mostPairs*10 + secondMostPairs
	// part two
	return (mostPairs+jokers)*10 + secondMostPairs
}

type CamcelCardHandsByScore []CamelCardHand

func (hands CamcelCardHandsByScore) Len() int {
	return len(hands)
}

func (hands CamcelCardHandsByScore) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func (hands CamcelCardHandsByScore) Less(i, j int) bool {
	camcelCardRuneToValue := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		// 'J': 11,
		// Part two
		'J': 1,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	aScore := hands[i].Score()
	bScore := hands[j].Score()

	if aScore == bScore {
		for k := 0; k < len(hands[i].cards); k++ {
			aCard := hands[i].cards[k]
			bCard := hands[j].cards[k]
			if camcelCardRuneToValue[aCard] != camcelCardRuneToValue[bCard] {
				return camcelCardRuneToValue[aCard] < camcelCardRuneToValue[bCard]
			}
		}
	}
	return aScore < bScore
}

func (camelCards *CamelCards) SortHands() {
	sort.Sort(CamcelCardHandsByScore(camelCards.hands))
}

func (camelCards *CamelCards) TotalWinnings() uint64 {
	camelCards.SortHands()
	var sum uint64 = 0
	for i, hand := range camelCards.hands {
		sum = sum + uint64(hand.bid*(i+1))
	}
	return sum
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)

	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	camelCards := parseCamelCards(lines)

	fmt.Println(camelCards.TotalWinnings())

	panicIfError(fileScanner.Err())
}
