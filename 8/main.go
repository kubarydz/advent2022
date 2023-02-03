package main

import (
	"fmt"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 21 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 8 {
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

	sumVisible := 0
	for i := 1; i < len(chunked)-1; i++ {
		for j := 1; j < len(chunked[0])-1; j++ {
			if isVisibleFromLeft(chunked[i], j) ||
				isVisibleFromRight(chunked[i], j) ||
				isVisibleFromTop(chunked, i, j) ||
				isVisibleFromBottom(chunked, i, j) {
				sumVisible++
			}
		}
	}

	sumVisible += len(chunked)*2 + len(chunked[0])*2 - 4

	return sumVisible
}

func isVisibleFromLeft(row []byte, pos int) bool {
	for i := pos - 1; i >= 0; i-- {
		if row[i] >= row[pos] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(row []byte, pos int) bool {
	for i := pos + 1; i < len(row); i++ {
		if row[i] >= row[pos] {
			return false
		}
	}
	return true
}

func isVisibleFromTop(input [][]byte, row, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if input[i][col] >= input[row][col] {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(input [][]byte, row, col int) bool {
	for i := row + 1; i < len(input); i++ {
		if input[i][col] >= input[row][col] {
			return false
		}
	}
	return true
}

func calculateV2(input []byte) int {
	chunked := utils.ChunkInput(input)

	highestScore := 0
	for i := 0; i < len(chunked); i++ {
		for j := 1; j < len(chunked[0]); j++ {
			scenicScore := calculateScenicScore(chunked, i, j)
			if scenicScore > highestScore {
				highestScore = scenicScore
			}
		}
	}
	return highestScore
}

func calculateScenicScore(input [][]byte, row, col int) int {
	if row == 0 || col == 0 || row == len(input)-1 || col == len(input[0])-1 {
		return 0
	}

	return scenicScoreLeft(input[row], col) * scenicScoreRight(input[row], col) * scenicScoreTop(input, row, col) * scenicScoreBottom(input, row, col)
}

func scenicScoreLeft(input []byte, pos int) int {
	score := 0
	for i := pos - 1; i >= 0; i-- {
		score++
		if input[i] >= input[pos] {
			return score
		}
	}
	return score
}

func scenicScoreRight(input []byte, pos int) int {
	score := 0
	for i := pos + 1; i < len(input); i++ {
		score++
		if input[i] >= input[pos] {
			return score
		}
	}
	return score
}

func scenicScoreTop(input [][]byte, row, col int) int {
	score := 0
	for i := row - 1; i >= 0; i-- {
		score++
		if input[i][col] >= input[row][col] {
			return score
		}
	}
	return score
}

func scenicScoreBottom(input [][]byte, row, col int) int {
	score := 0
	for i := row + 1; i < len(input); i++ {
		score++
		if input[i][col] >= input[row][col] {
			return score
		}
	}
	return score
}
