package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SequenceReading []int

func (s SequenceReading) GetDifferences() []int {
	result := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		result[i] = (s)[i+1] - (s)[i]
	}
	return result
}

func (s SequenceReading) IsAllZero() bool {
	for _, value := range s {
		if value != 0 {
			return false
		}
	}
	return true
}

func (s SequenceReading) Last() int {
	return s[len(s)-1]
}

func (s SequenceReading) First() int {
	return s[0]
}

func (s SequenceReading) NextValue() int {
	steps := make([][]int, 0)
	previous := s

	for !previous.IsAllZero() {
		steps = append(steps, previous.GetDifferences())
		previous = previous.GetDifferences()
	}

	lastValuesOfSteps := make([]int, len(steps))
	for i, step := range steps {
		lastValuesOfSteps[i] = step[len(step)-1]
	}

	change := 0
	for i := len(lastValuesOfSteps) - 1; i >= 0; i-- {
		change += lastValuesOfSteps[i]
	}

	return s.Last() + change
}

func (s SequenceReading) PreviousValue() int {
	steps := make([][]int, 0)
	previous := s

	for !previous.IsAllZero() {
		steps = append(steps, previous.GetDifferences())
		previous = previous.GetDifferences()
	}

	fmt.Println(steps)

	firstValuesOfSteps := make([]int, len(steps))
	for i, step := range steps {
		firstValuesOfSteps[i] = step[0]
	}

	change := firstValuesOfSteps[len(firstValuesOfSteps)-1]
	for i := len(firstValuesOfSteps) - 2; i >= 0; i-- {
		change = firstValuesOfSteps[i] - change
	}
	return s.First() - change
}

type InstabilitySensor struct {
	readings []SequenceReading
}

func (i *InstabilitySensor) SumNextValues() int {
	sum := 0
	for _, reading := range i.readings {
		sum += reading.NextValue()
	}
	return sum
}

func (i *InstabilitySensor) SumPreviousValues() int {
	sum := 0
	for _, reading := range i.readings {
		sum += reading.PreviousValue()
	}
	return sum
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func stringsToIntArray(s []string) []int {
	var result []int
	for _, numberString := range s {
		value, err := strconv.Atoi(numberString)
		panicIfError(err)

		result = append(result, value)
	}
	return result
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)

	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	instabilitySensor := InstabilitySensor{readings: make([]SequenceReading, 0)}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		instabilitySensor.readings = append(instabilitySensor.readings, stringsToIntArray(fields))
	}

	panicIfError(fileScanner.Err())

	// Part one
	fmt.Println(instabilitySensor.SumNextValues())
	// Part two
	fmt.Println(instabilitySensor.SumPreviousValues())

}
