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

type GameSet struct {
	blueCount  int
	redCount   int
	greenCount int
}

type Game struct {
	id   int
	sets []GameSet
}

func (g *GameSet) isWithinLimits(blueLimit int, redLimit int, greenLimit int) bool {
	return g.blueCount <= blueLimit && g.redCount <= redLimit && g.greenCount <= greenLimit
}

func (g *Game) isGameValid() bool {
	for _, set := range g.sets {
		// blue, red, green
		if !set.isWithinLimits(14, 12, 13) {
			return false
		}
	}
	return true
}

func (g *Game) minimumPlayableSet() GameSet {
	minRed, minBlue, minGreen := 0, 0, 0
	for _, set := range g.sets {
		if set.redCount > minRed {
			minRed = set.redCount
		}
		if set.blueCount > minBlue {
			minBlue = set.blueCount
		}
		if set.greenCount > minGreen {
			minGreen = set.greenCount
		}
	}
	return GameSet{
		redCount:   minRed,
		blueCount:  minBlue,
		greenCount: minGreen,
	}
}

func parseGame(line string) Game {
	gameSets := []GameSet{}

	gameFromSetsSplit := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.Split(gameFromSetsSplit[0], " ")[1])
	gameSetsPart := gameFromSetsSplit[1]
	panicIfError(err)

	gameSetsSplit := strings.Split(gameSetsPart, ";")

	for _, set := range gameSetsSplit {
		gamePlayedSet := GameSet{}
		gamePlayedSetSplit := strings.Split(set, ",")
		for _, setSplit := range gamePlayedSetSplit {
			cubesPulls := strings.Split(setSplit, ",")

			for _, cubePull := range cubesPulls {
				cubePullSplit := strings.Split(cubePull, " ")
				cubeColor := cubePullSplit[2]

				cubeCount, err := strconv.Atoi(cubePullSplit[1])
				panicIfError(err)

				switch cubeColor {
				case "blue":
					gamePlayedSet.blueCount = cubeCount
				case "red":
					gamePlayedSet.redCount = cubeCount
				case "green":
					gamePlayedSet.greenCount = cubeCount
				}
			}
		}
		gameSets = append(gameSets, gamePlayedSet)
	}

	game := Game{
		id:   gameId,
		sets: gameSets,
	}
	return game
}

func (g *Game) getPower() int {
	minimunSet := g.minimumPlayableSet()
	return minimunSet.blueCount * minimunSet.redCount * minimunSet.greenCount
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

	for fileScanner.Scan() {
		// For each line...
		line := fileScanner.Text()
		game := parseGame(line)
		// Part A
		// if game.isGameValid() {
		// 	fmt.Println(game)
		// 	count = count + game.id
		// }
		// Part B
		power := game.getPower()
		count = count + power
	}

	panicIfError(fileScanner.Err())

	fmt.Println(count)
}
