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

func isNumeric(c rune) bool {
	// Check if the character is within the numeric ascii range
	return c >= '0' && c <= '9'
}

func FindFirstDigit(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if isNumeric(c) {
			return string(runes[i])
		}
	}

	// Should not occur in the AOC exercise
	panic(fmt.Sprintf("No digit found in string %s", s))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ParseWritenDigitsToNumber(line string) string {
	wordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Smallest word is 3 characters
	i := 0
	for i < len(line)-2 {
		for word, digit := range wordToDigit {
			if strings.HasPrefix(line[i:], word) {
				// Append the last character of the word to the digit to avoid replacing other valid digits from the end
				lastCharOfWord := word[len(word)-1:]
				line = strings.Replace(line, word, digit+lastCharOfWord, 1)
				i = 0
				break
			}
		}
		i++
	}

	return line
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	count := 0

	// For each line...
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parsedLine := ParseWritenDigitsToNumber(line)

		firstDigitChar := FindFirstDigit(parsedLine)
		lastDigitChar := FindFirstDigit(Reverse(parsedLine))
		digitString := firstDigitChar + lastDigitChar

		result, err := strconv.Atoi(digitString)
		panicIfError(err)

		count = count + result
	}

	panicIfError(fileScanner.Err())

	fmt.Println(count)
}
