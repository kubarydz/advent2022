package main

import (
	"fmt"
	"strconv"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 88 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 36 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)
}

func calculateV1(input []byte) int {
	chunked := utils.ChunkInput(input)

	headPosition := intPair{0, 0}
	tailPosition := intPair{0, 0}
	allTailPositions := map[intPair]struct{}{tailPosition: {}}
	for _, move := range chunked {
		nextMove := getMovementDirection(move[0])
		nrOfMoves, _ := strconv.Atoi(string(move[2:]))
		for i := 0; i < nrOfMoves; i++ {
			headPosition.x += nextMove.x
			headPosition.y += nextMove.y
			if shouldTailMove(headPosition, tailPosition) {
				tailPosition = newTailPosition(headPosition, tailPosition)
				allTailPositions[tailPosition] = struct{}{}
			}
		}
	}

	return len(allTailPositions)
}

func calculateV2(input []byte) int {
	chunked := utils.ChunkInput(input)

	tailsPositions := [10]intPair{}
	allTailPositions := map[intPair]struct{}{tailsPositions[9]: {}}
	for _, move := range chunked {
		nextMove := getMovementDirection(move[0])
		nrOfMoves, _ := strconv.Atoi(string(move[2:]))
		for i := 0; i < nrOfMoves; i++ {
			tailsPositions[0].x += nextMove.x
			tailsPositions[0].y += nextMove.y

			if !shouldTailMove(tailsPositions[1], tailsPositions[0]) {
				continue
			}

			for i := 1; i < len(tailsPositions); i++ {
				if shouldTailMove(tailsPositions[i-1], tailsPositions[i]) {
					tailsPositions[i] = newTailPosition(tailsPositions[i-1], tailsPositions[i])
				}
			}
			allTailPositions[tailsPositions[9]] = struct{}{}
		}
	}

	return len(allTailPositions)
}

type intPair struct {
	x int
	y int
}

func getMovementDirection(b byte) intPair {
	x, y := 0, 0
	switch b {
	case 'R':
		x = 1
	case 'L':
		x = -1
	case 'U':
		y = 1
	case 'D':
		y = -1
	}
	return intPair{x: x, y: y}
}

func shouldTailMove(headPosition, tailPosition intPair) bool {
	return utils.Abs(headPosition.x-tailPosition.x) > 1 ||
		utils.Abs(headPosition.y-tailPosition.y) > 1
}

func newTailPosition(headPosition, tailPosition intPair) intPair {
	newTailPosition := tailPosition
	if headPosition.x != tailPosition.x {
		if headPosition.x > tailPosition.x {
			newTailPosition.x++
		} else {
			newTailPosition.x--
		}
	}

	if headPosition.y != tailPosition.y {
		if headPosition.y > tailPosition.y {
			newTailPosition.y++
		} else {
			newTailPosition.y--
		}
	}

	return newTailPosition
}
