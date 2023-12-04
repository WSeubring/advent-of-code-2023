package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

type ScratchCard struct {
	id             int
	winningNumbers map[int]bool
	myNumbers      []int
}

func (scratchCard *ScratchCard) Score() int {
	score := 0
	for _, number := range scratchCard.myNumbers {
		if scratchCard.winningNumbers[number] {
			// Part A
			// score = score * 2
			// if score == 0 {
			// 	score = 1
			// }
			score++
		}
	}
	return score
}

func parseScratchCard(line string) ScratchCard {
	cardIdFromNumberSplit := strings.Split(line, ":")

	cardIdString := cardIdFromNumberSplit[0]
	cardIdParts := strings.Split(cardIdString, " ")
	cardId, err := strconv.Atoi(cardIdParts[len(cardIdParts)-1])
	panicIfError(err)

	cardParts := strings.Split(cardIdFromNumberSplit[1], "|")
	winningNumbersPart := cardParts[0]

	// Map the string of numbers to a array of ints
	myNumbersAsString := strings.Split(cardParts[1], " ")
	myNumbers := make([]int, len(myNumbersAsString))
	for i, number := range myNumbersAsString {
		if number == "" {
			continue
		}
		numberInt, err := strconv.Atoi(number)
		panicIfError(err)
		myNumbers[i] = numberInt
	}

	// Map the string of numbers to a map of ints
	winningNumbersDict := make(map[int]bool)
	winningNumbers := strings.Split(winningNumbersPart, " ")
	for _, number := range winningNumbers {
		if number == "" {
			continue
		}
		numberInt, err := strconv.Atoi(number)
		panicIfError(err)
		winningNumbersDict[numberInt] = true
	}

	// Build the card game
	game := ScratchCard{
		id:             cardId,
		winningNumbers: winningNumbersDict,
		myNumbers:      myNumbers,
	}

	return game
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	scratchCards := []ScratchCard{}

	for fileScanner.Scan() {
		// For each line...
		line := fileScanner.Text()
		scratchCard := parseScratchCard(line)
		scratchCards = append(scratchCards, scratchCard)
	}
	panicIfError(fileScanner.Err())

	// Calculate the score for each card id
	scorePerId := make(map[int]int)
	for _, card := range scratchCards {
		scorePerId[card.id] = card.Score()
	}

	// Add starting cards in count
	cardsPerId := make(map[int]int)
	for _, card := range scratchCards {
		cardsPerId[card.id] = 1
	}

	// Calculate the number of cards won for each card id
	for _, card := range scratchCards {
		score := scorePerId[card.id]

		for i := 1; i <= score; i++ {
			cardsPerId[card.id+i] = cardsPerId[card.id+i] + cardsPerId[card.id]
		}
	}

	sum := 0
	for _, value := range cardsPerId {
		sum = sum + value
	}
	fmt.Println(sum)

}
