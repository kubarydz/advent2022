package main

import (
	"fmt"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateRPSScore(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 15 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateScoreV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 12 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	resp = calculateRPSScore(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	resp = calculateScoreV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)
}

func calculateRPSScore(input []byte) int {
	var chunks [][]byte

	for i := 0; i < len(input); i += 4 {
		chunks = append(chunks, input[i:i+3])
	}

	score := 0
	for _, chunk := range chunks {
		score += getSingleScore(chunk[0], chunk[2])
	}
	return score
}

func getSingleScore(a, b byte) int {
	winLoseScore := getWinLoseScore(a, b)
	switch b {
	case 'X':
		return winLoseScore + 1
	case 'Y':
		return winLoseScore + 2
	case 'Z':
		return winLoseScore + 3
	}

	return 0
}

func getWinLoseScore(a, b byte) int {
	if b == 'X' {
		b = 'A'
	} else if b == 'Y' {
		b = 'B'
	} else {
		b = 'C'
	}

	if a == b {
		return 3
	}
	if a == 'A' && b == 'B' {
		return 6
	}
	if a == 'B' && b == 'C' {
		return 6
	}
	if a == 'C' && b == 'A' {
		return 6
	}

	return 0
}

func calculateScoreV2(input []byte) int {
	var chunks [][]byte

	for i := 0; i < len(input); i += 4 {
		chunks = append(chunks, input[i:i+3])
	}

	score := 0
	for _, chunk := range chunks {
		score += getSingleV2Score(chunk[0], chunk[2])
	}
	return score
}

func getSingleV2Score(a, b byte) int {
	winLosePoints := 0
	if b == 'Y' {
		winLosePoints = 3
	} else if b == 'Z' {
		winLosePoints = 6
	}
	shapePoints := 1
	if b == 'Y' {
		shapePoints += int(a - 'A')
	}
	if b == 'Z' {
		switch a {
		case 'A':
			shapePoints = 2
		case 'B':
			shapePoints = 3
		}
	}
	if b == 'X' {
		switch a {
		case 'A':
			shapePoints = 3
		case 'C':
			shapePoints = 2
		}
	}
	return winLosePoints + shapePoints
}
