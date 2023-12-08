package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
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

	instructionSequence := lines[0]

	type Direction struct {
		left  string
		right string
	}

	mapDirections := make(map[string]Direction)
	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		key := fields[0]
		left := strings.Trim(fields[2], "(),")
		right := strings.Trim(fields[3], "()),")

		mapDirections[key] = Direction{left: left, right: right}
	}

	stepCount := 0

	start := "AAA"
	end := "ZZZ"
	current := start
	for i := 0; i < len(instructionSequence); {
		fmt.Printf("current: %v\n", current)
		if current == end {
			break
		}

		instruction := string(instructionSequence[i])

		currentDirection := mapDirections[current]
		if instruction == "L" {
			current = currentDirection.left
		} else if instruction == "R" {
			current = currentDirection.right
		}
		stepCount++

		if i == len(instructionSequence)-1 && current != end {
			i = 0
		} else {
			i++
		}
	}

	fmt.Println(stepCount)

	panicIfError(fileScanner.Err())
}
