package main

import (
	"bufio"
	"errors"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pipes := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pipes = append(pipes, scanner.Text())
	}

	startPos, err := getStartPos(pipes)
	if err != nil {
		panic(err)
	}

	prevPos := startPos
	currentPos := startPos

	loopLength := 1

	if canGoUp(currentPos, pipes) {
		currentPos = Coord{currentPos.row - 1, currentPos.col}
	} else if canGoDown(currentPos, pipes) {
		currentPos = Coord{currentPos.row + 1, currentPos.col}
	} else if canGoLeft(currentPos, pipes) {
		currentPos = Coord{currentPos.row, currentPos.col - 1}
	}

	for true {
		if currentPos == startPos {
			break
		}

		loopLength++
		// currentPipe := pipes[currentPos.row][currentPos.col]
		// println(string(currentPipe))

		tmpPos := currentPos
		currentPos = getNextPipe(currentPos, prevPos, pipes)
		prevPos = tmpPos
	}

	println(loopLength / 2)

}

func partTwo() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

func getStartPos(pipes []string) (Coord, error) {
	for i, row := range pipes {
		for j, col := range row {
			if col == 'S' {
				return Coord{i, j}, nil
			}
		}
	}
	return Coord{-1, -1}, errors.New("Starting position not found")
}

func getNextPipe(currentPos Coord, prevPos Coord, pipes []string) Coord {
	// Check up
	if prevPos.row != currentPos.row-1 && canGoUp(currentPos, pipes) {
		return Coord{currentPos.row - 1, currentPos.col}
	}

	// Check right
	if prevPos.col != currentPos.col+1 && canGoRight(currentPos, pipes) {
		return Coord{currentPos.row, currentPos.col + 1}
	}

	// Check down
	if prevPos.row != currentPos.row+1 && canGoDown(currentPos, pipes) {
		return Coord{currentPos.row + 1, currentPos.col}
	}

	// Check left
	if prevPos.col != currentPos.col-1 && canGoLeft(currentPos, pipes) {
		return Coord{currentPos.row, currentPos.col - 1}
	}

	return currentPos
}

func canGoUp(currentPos Coord, pipes []string) bool {
	if currentPos.row <= 0 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row-1][currentPos.col]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['|'] = struct{}{}
	validNextPipes['7'] = struct{}{}
	validNextPipes['F'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['|'] = struct{}{}
	validCurrentPipes['L'] = struct{}{}
	validCurrentPipes['J'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoDown(currentPos Coord, pipes []string) bool {
	if currentPos.row >= len(pipes) {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row+1][currentPos.col]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['|'] = struct{}{}
	validNextPipes['L'] = struct{}{}
	validNextPipes['J'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['|'] = struct{}{}
	validCurrentPipes['7'] = struct{}{}
	validCurrentPipes['F'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoLeft(currentPos Coord, pipes []string) bool {
	if currentPos.col <= 0 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row][currentPos.col-1]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['-'] = struct{}{}
	validNextPipes['L'] = struct{}{}
	validNextPipes['F'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['-'] = struct{}{}
	validCurrentPipes['7'] = struct{}{}
	validCurrentPipes['J'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoRight(currentPos Coord, pipes []string) bool {
	if currentPos.col >= len(pipes[currentPos.row])-1 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row][currentPos.col+1]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['-'] = struct{}{}
	validNextPipes['7'] = struct{}{}
	validNextPipes['J'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['-'] = struct{}{}
	validCurrentPipes['L'] = struct{}{}
	validCurrentPipes['F'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

type Coord struct {
	row int
	col int
}
