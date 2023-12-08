package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func Distance(timePressed int, raceTime int) int {
	return timePressed * (raceTime - timePressed)
}

func AbcFormula(a, b, c float64) (float64, float64) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		panic("No real roots")
	}
	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return x1, x2
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
	times := strings.Fields(lines[0])[1:]
	targetDistances := strings.Fields(lines[1])[1:]

	var count float64 = 0

	for i, time := range times {
		targetDistance, err := strconv.Atoi(targetDistances[i])
		panicIfError(err)
		time, err := strconv.Atoi(time)
		panicIfError(err)

		x1, x2 := AbcFormula(-1, float64(time), -(float64(targetDistance) + 0.0000001))
		winningRange := math.Floor(x2) - math.Ceil(x1) + 1

		count *= winningRange

		if count == 0 {
			count = winningRange
		}
	}

	fmt.Println(count)

	panicIfError(fileScanner.Err())
}
