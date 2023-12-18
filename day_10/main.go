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
	left     int
	top      int
	right    int
	bottom   int
	x        int
	y        int
	distance int
}

func (pipe *Pipe) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", pipe.left, pipe.top, pipe.right, pipe.bottom)
}

func (pipe *Pipe) isStart() bool {
	return pipe.left == 1 && pipe.top == 1 && pipe.right == 1 && pipe.bottom == 1
}

func (pipe *Pipe) isGround() bool {
	return pipe.left == 0 && pipe.top == 0 && pipe.right == 0 && pipe.bottom == 0
}

func parseGridToPipes(grid [][]rune) [][]Pipe {

	pipes := make([][]Pipe, len(grid))
	for y, row := range grid {
		pipes[y] = make([]Pipe, len(row))
		for x, char := range row {
			pipesMap := map[rune](Pipe){
				// Top-bottom
				'|': Pipe{left: 0, top: 1, right: 0, bottom: 1, x: x, y: y, distance: 0},
				// Left-right
				'-': Pipe{left: 1, top: 0, right: 1, bottom: 0, x: x, y: y, distance: 0},
				// top-right
				'L': Pipe{left: 0, top: 1, right: 1, bottom: 0, x: x, y: y, distance: 0},
				// left-top
				'J': Pipe{left: 1, top: 1, right: 0, bottom: 0, x: x, y: y, distance: 0},
				// left-bottom
				'7': Pipe{left: 1, top: 0, right: 0, bottom: 1, x: x, y: y, distance: 0},
				// bottom-right
				'F': Pipe{left: 0, top: 0, right: 1, bottom: 1, x: x, y: y, distance: 0},

				'.': Pipe{0, 0, 0, 0, x, y, 0},
				'S': Pipe{1, 1, 1, 1, x, y, 0},
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
	if nextPipe.x < pipe.x {
		return (pipe.left == 1 && nextPipe.right == 1)
	}

	if nextPipe.x > pipe.x {
		return (pipe.right == 1 && nextPipe.left == 1)
	}

	if nextPipe.y < pipe.y {
		return (pipe.top == 1 && nextPipe.bottom == 1)
	}

	if nextPipe.y > pipe.y {
		return (pipe.bottom == 1 && nextPipe.top == 1)
	}

	return false
}

func (pipe *Pipe) getConnectedPipes(pipes [][]Pipe) []Pipe {
	connectedPipes := make([]Pipe, 0)

	for _, offset := range []int{-1, 1} {
		// Check vertical neighbours (up and down)
		if pipe.y+offset >= 0 && pipe.y+offset < len(pipes) {
			if pipe.isConnected(pipes[pipe.y+offset][pipe.x]) {
				connectedPipes = append(connectedPipes, pipes[pipe.y+offset][pipe.x])
			}
		}

		// Check horizontal neighbours (left and right)
		if pipe.x+offset >= 0 && pipe.x+offset < len(pipes[pipe.y]) {
			if pipe.isConnected(pipes[pipe.y][pipe.x+offset]) {
				connectedPipes = append(connectedPipes, pipes[pipe.y][pipe.x+offset])
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
		currentPipeCopy := queue[0]
		currentPipe := pipes[currentPipeCopy.y][currentPipeCopy.x]

		queue = queue[1:]

		connectedPipes := currentPipe.getConnectedPipes(pipes)

		for _, connectedPipe := range connectedPipes {
			if connectedPipe.distance == 0 {
				pipes[connectedPipe.y][connectedPipe.x].distance = currentPipe.distance + 1

				queue = append(queue, connectedPipe)
			}
		}
	}

	// Print the grid
	for _, row := range pipes {
		for _, pipe := range row {
			if pipe.isStart() {
				fmt.Printf("S")
				continue
			}
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

	fmt.Println(pipes[3][1].isConnected(pipes[3][2]))
	fmt.Println(pipes[2][3].isConnected(pipes[3][3]))

	fmt.Println(maxDistance)
}
