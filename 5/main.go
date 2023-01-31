package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %s\n", resp)
	if resp != "CMZ" {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %s\n", resp)
	if resp != "MCD" {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %s\n", resp)

	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %s\n", resp)
}

func calculateV1(input []byte) string {
	chunks := utils.ChunkInput(input)
	numbersRowIndex := getNumbersRowIndex(chunks)

	crateStacks := getCrateStacks(chunks, numbersRowIndex)

	moves := chunks[numbersRowIndex+2:]

	for _, move := range moves {
		splitMove := strings.Split(string(move), " ")

		howManyMove, _ := strconv.Atoi(splitMove[1])
		moveFrom, _ := strconv.Atoi(splitMove[3])
		moveTo, _ := strconv.Atoi(splitMove[5])

		for i := 0; i < howManyMove; i++ {
			crateStacks[moveTo-1] = append(crateStacks[moveTo-1], crateStacks[moveFrom-1][len(crateStacks[moveFrom-1])-1])
			crateStacks[moveFrom-1] = crateStacks[moveFrom-1][:len(crateStacks[moveFrom-1])-1]
		}
	}

	topCrates := []byte{}
	for _, stack := range crateStacks {
		topCrates = append(topCrates, stack[len(stack)-1])
	}

	return string(topCrates)
}

func calculateV2(input []byte) string {
	chunks := utils.ChunkInput(input)
	numbersRowIndex := getNumbersRowIndex(chunks)

	crateStacks := getCrateStacks(chunks, numbersRowIndex)
	moves := chunks[numbersRowIndex+2:]

	for _, move := range moves {
		splitMove := strings.Split(string(move), " ")

		howManyMove, _ := strconv.Atoi(splitMove[1])
		moveFrom, _ := strconv.Atoi(splitMove[3])
		moveTo, _ := strconv.Atoi(splitMove[5])

		for i := howManyMove; i > 0; i-- {
			crateStacks[moveTo-1] = append(crateStacks[moveTo-1], crateStacks[moveFrom-1][len(crateStacks[moveFrom-1])-i])
		}
		crateStacks[moveFrom-1] = crateStacks[moveFrom-1][:len(crateStacks[moveFrom-1])-howManyMove]
	}

	topCrates := []byte{}
	for _, stack := range crateStacks {
		topCrates = append(topCrates, stack[len(stack)-1])
	}

	return string(topCrates)
}

func getNumbersRowIndex(chunks [][]byte) int {
	numbersRow := 0
	for i, row := range chunks {
		for _, col := range row {
			if unicode.IsNumber(rune(col)) {
				numbersRow = i
				break
			}
			if !unicode.IsSpace(rune(col)) {
				break
			}
		}
		if numbersRow != 0 {
			break
		}
	}
	return numbersRow
}

func getCrateStacks(chunks [][]byte, numbersRowIndex int) [][]byte {
	crateStacks := [][]byte{}
	for i := 0; i < len(chunks[0]); i++ {
		if !unicode.IsNumber(rune(chunks[numbersRowIndex][i])) {
			continue
		}

		stack := []byte{}
		for j := numbersRowIndex - 1; j >= 0; j-- {
			if unicode.IsSpace(rune(chunks[j][i])) {
				continue
			}
			stack = append(stack, chunks[j][i])
		}
		crateStacks = append(crateStacks, stack)

	}

	return crateStacks
}
