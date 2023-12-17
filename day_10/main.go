package main

import (
	"bufio"
	"fmt"
	"os"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func findStart(pipes [][]Pipe) Pipe {
	for _, row := range pipes {
		for _, pipe := range row {
			if pipe.isStart() {
				return pipe
			}
		}
	}
	panic("No start found")
}

type Pipe struct {
	x1       int
	y1       int
	x2       int
	y2       int
	posX     int
	posY     int
	distance int
}

func (pipe *Pipe) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", pipe.x1, pipe.y1, pipe.x2, pipe.y2)
}

func (pipe *Pipe) isStart() bool {
	return pipe.x1 == 10 && pipe.y1 == 10 && pipe.x2 == 10 && pipe.y2 == 10
}

func (pipe *Pipe) isGround() bool {
	return pipe.x1 == -10 && pipe.y1 == -10 && pipe.x2 == -10 && pipe.y2 == -10
}

func parseGridToPipes(grid [][]rune) [][]Pipe {

	pipes := make([][]Pipe, len(grid))
	for y, row := range grid {
		pipes[y] = make([]Pipe, len(row))
		for x, char := range row {
			pipesMap := map[rune](Pipe){
				'|': Pipe{0, -1, 0, 1, x, y, 0},
				'-': Pipe{-1, 0, 1, 0, x, y, 0},
				'L': Pipe{-1, 0, 0, -1, x, y, 0},
				'J': Pipe{0, -1, 1, 0, x, y, 0},
				'7': Pipe{0, 1, -1, 0, x, y, 0},
				'F': Pipe{1, 0, 0, 1, x, y, 0},
				'.': Pipe{-10, -10, -10, -10, x, y, 0},
				'S': Pipe{10, 10, 10, 10, x, y, 0},
			}
			if pipe, ok := pipesMap[char]; ok {
				pipes[y][x] = pipe
			} else {
				pipes[y][x] = Pipe{-1, -1, -1, -1, x, y, 0}
			}
		}
	}

	return pipes
}

func (pipe *Pipe) isConnected(nextPipe Pipe) bool {
	return (pipe.x1 == nextPipe.x1 && pipe.y1 == nextPipe.y1) ||
		(pipe.x1 == nextPipe.x2 && pipe.y1 == nextPipe.y2) ||
		(pipe.x2 == nextPipe.x1 && pipe.y2 == nextPipe.y1) ||
		(pipe.x2 == nextPipe.x2 && pipe.y2 == nextPipe.y2) ||
		(pipe.isStart() && !nextPipe.isGround()) ||
		(!pipe.isGround() && nextPipe.isStart())
}

func (pipe *Pipe) getConnectedPipes(pipes [][]Pipe) []Pipe {
	connectedPipes := make([]Pipe, 0)

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if pipe.posY+i < 0 || pipe.posY+i >= len(pipes) || pipe.posX+j < 0 || pipe.posX+j >= len(pipes[pipe.posY+i]) {
				// Out of bounds
				continue
			}

			if i == 0 && j == 0 {
				// Do not self check
				continue
			}
			if pipe.isConnected(pipes[pipe.posY+i][pipe.posX+j]) {
				connectedPipes = append(connectedPipes, pipes[pipe.posY+i][pipe.posX+j])
			}
		}
	}
	return connectedPipes
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)

	panicIfError(err)
	// Close the file when we leave the scope of the current function,
	defer file.Close()

	// Make a buffer to keep chunks that are read.
	fileScanner := bufio.NewScanner(file)

	grid := make([][]rune, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
	}
	panicIfError(fileScanner.Err())

	pipes := parseGridToPipes(grid)
	startPipe := findStart(pipes)

	// Queue for pipes to check (BFS)
	queue := make([]Pipe, 0)
	queue = append(queue, startPipe)

	for len(queue) > 0 {
		// Get the first pipe in the queue
		currentPipeCopy := queue[0]
		currentPipe := pipes[currentPipeCopy.posY][currentPipeCopy.posX]

		queue = queue[1:]

		// Get the connected pipes of the current pipe
		connectedPipes := currentPipe.getConnectedPipes(pipes)

		// For each connected pipe
		for _, connectedPipe := range connectedPipes {
			// If the pipe is not visited
			if connectedPipe.distance == 0 {
				// Mark the pipe as visited
				pipes[connectedPipe.posY][connectedPipe.posX].distance = currentPipe.distance + 1

				// Append the connected pipes of the current pipe
				queue = append(queue, connectedPipe)
			}
		}
	}

	// Print the grid
	for _, row := range pipes {
		for _, pipe := range row {
			fmt.Printf("%d", pipe.distance%10)
		}
		fmt.Println()
	}

	// Find max distance
	maxDistance := 0
	for _, row := range pipes {
		for _, pipe := range row {
			if pipe.distance > maxDistance {
				maxDistance = pipe.distance
			}
		}
	}

	fmt.Println(maxDistance)
}
