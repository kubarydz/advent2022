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
	if resp != 2 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 4 {
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
	for _, c := range chunks {
		firstDash := 0
		secondDash := 0
		comma := 0
		for i, b := range c {
			if b == ',' {
				comma = i
				continue
			}
			if b == '-' {
				if firstDash == 0 {
					firstDash = i
					continue
				}
				secondDash = i
				break
			}
		}
		firstStart, _ := strconv.Atoi(string(c[:firstDash]))
		firstFinish, _ := strconv.Atoi(string(c[firstDash+1 : comma]))

		secondStart, _ := strconv.Atoi(string(c[comma+1 : secondDash]))
		secondFinish, _ := strconv.Atoi(string(c[secondDash+1:]))

		if (firstStart <= secondStart && firstFinish >= secondFinish) ||
			(secondStart <= firstStart && secondFinish >= firstFinish) {
			sum++
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
	for _, c := range chunks {
		firstDash := 0
		secondDash := 0
		comma := 0
		for i, b := range c {
			if b == ',' {
				comma = i
				continue
			}
			if b == '-' {
				if firstDash == 0 {
					firstDash = i
					continue
				}
				secondDash = i
				break
			}
		}
		firstStart, _ := strconv.Atoi(string(c[:firstDash]))
		firstFinish, _ := strconv.Atoi(string(c[firstDash+1 : comma]))

		secondStart, _ := strconv.Atoi(string(c[comma+1 : secondDash]))
		secondFinish, _ := strconv.Atoi(string(c[secondDash+1:]))

		if (firstStart <= secondStart && secondStart <= firstFinish) ||
			(secondStart <= firstStart && firstStart <= secondFinish) {
			sum++
		}
	}

	return sum

}
