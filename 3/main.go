package main

import (
	"fmt"

	"github.com/kubarydz/advent2022/utils"
	"golang.org/x/exp/slices"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 157 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 70 {
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
	var chunks [][]byte

	lastStart := 0
	for i, b := range input {
		if b == '\n' {
			chunks = append(chunks, input[lastStart:i])
			lastStart = i + 1
		}
	}

	sum := 0
	for _, chunk := range chunks {
		first := chunk[:len(chunk)/2]
		second := chunk[len(chunk)/2:]

		used := map[byte]struct{}{}
		for _, f := range first {
			for _, s := range second {
				if f == s {
					_, ok := used[f]
					if ok {
						continue
					}
					used[f] = struct{}{}
					if f >= 'a' && f <= 'z' {
						sum += int(f) - 'a' + 1
					} else {
						sum += int(f) - 'A' + 27
					}
				}
			}
		}
	}
	return sum
}

func calculateV2(input []byte) int {
	var chunks [][]byte

	lastStart := 0
	for i, b := range input {
		if b == '\n' {
			chunks = append(chunks, input[lastStart:i])
			lastStart = i + 1
		}
	}

	sum := 0
	for i := 0; i < len(chunks); i += 3 {
		possibleMatches := []byte{}
		for _, f := range chunks[i] {
			for _, s := range chunks[i+1] {
				if f == s {
					if !slices.Contains(possibleMatches, f) {
						possibleMatches = append(possibleMatches, f)
					}

				}
			}
		}

		for _, f := range possibleMatches {
			for _, s := range chunks[i+2] {
				if f == s {
					if f >= 'a' && f <= 'z' {
						sum += int(f) - 'a' + 1
					} else {
						sum += int(f) - 'A' + 27
					}
					break
				}

			}
		}
	}
	return sum
}
