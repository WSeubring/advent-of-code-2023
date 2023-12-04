// Note: Day 3, was heavily inspired by Copilot
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	grid := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
	}

	count := 0
	for i, line := range grid {
		numberString := ""
		for j, cell := range line {
			if unicode.IsDigit(cell) {
				numberString += string(cell)
			} else {
				if numberString != "" {
					count += processNumber(numberString, i, j, grid)
					numberString = ""
				}
			}
		}
		if numberString != "" {
			count += processNumber(numberString, i, len(line), grid)
		}
	}

	fmt.Printf("count: %d\n", count)
}

func processNumber(numberString string, i, j int, grid [][]rune) int {
	number, _ := strconv.Atoi(numberString)
	for k := 0; k < len(numberString); k++ {
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				ni, nj := i+di, j-len(numberString)+k+dj
				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) && !unicode.IsDigit(grid[ni][nj]) && grid[ni][nj] != '.' {
					return number
				}
			}
		}
	}
	return 0
}
