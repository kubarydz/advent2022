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
	if resp != 10 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 29 {
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
	return findDistinctCharacters(input, 4)
}

func calculateV2(input []byte) int {
	return findDistinctCharacters(input, 14)
}

func findDistinctCharacters(input []byte, distinctSize int) int {
	marker := 0
	seen := make(map[byte]int)
	nextPossibleMarker := distinctSize

	for i := 0; i < len(input); i++ {
		if _, ok := seen[input[i]]; !ok && i >= nextPossibleMarker {
			marker = i
			break
		}

		if seen[input[i]]+distinctSize > nextPossibleMarker {
			nextPossibleMarker = seen[input[i]] + distinctSize
		}

		seen[input[i]] = i
		if i < distinctSize {
			continue
		}

		if whenSeen := seen[input[i-distinctSize+1]]; whenSeen == i-distinctSize+1 {
			delete(seen, input[i-distinctSize+1])
		}
	}
	return marker + 1
}
