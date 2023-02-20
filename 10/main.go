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
	if resp != 13140 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	fmt.Printf("sample 2 solution:\n")
	calculateV2(input)

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	fmt.Printf("input 2 solution:\n")
	calculateV2(input)
}

func calculateV1(input []byte) int {
	moves := utils.ChunkInput(input)

	cycle := 1
	signalStrength := 1
	sumOfSignals := 0
	for _, move := range moves {
		if (cycle-20)%40 == 0 {
			sumOfSignals += signalStrength * cycle
		}

		moveStr := string(move[:4])

		if moveStr == "noop" {
			cycle++
		}
		if moveStr == "addx" {
			cycle++
			if (cycle-20)%40 == 0 {
				sumOfSignals += signalStrength * cycle
			}
			cycle++
			signal, _ := strconv.Atoi(string(move[5:]))
			signalStrength += signal
		}
	}

	return sumOfSignals
}

func calculateV2(input []byte) {
	chunked := utils.ChunkInput(input)

	lines := [6][40]byte{}
	spriteMiddle := 1
	currentLine := 0
	cycle := 1
	linePosition := 0
	for _, chunk := range chunked {
		if linePosition == 40 {
			currentLine++
			linePosition = 0
		}

		drawPixel(&lines, spriteMiddle, currentLine, linePosition)
		linePosition++

		moveStr := string(chunk[:4])
		if moveStr == "noop" {
			cycle++
		}
		if moveStr == "addx" {
			cycle++
			if linePosition == 40 {
				currentLine++
				linePosition = 0
			}
			drawPixel(&lines, spriteMiddle, currentLine, linePosition)
			cycle++
			linePosition++
			signal, _ := strconv.Atoi(string(chunk[5:]))
			spriteMiddle += signal
		}
	}

	drawBoard(&lines)
}

func drawPixel(lines *[6][40]byte, spriteMiddle, currentLine, linePosition int) {
	if spriteMiddle == linePosition || spriteMiddle == linePosition-1 || spriteMiddle == linePosition+1 {
		lines[currentLine][linePosition] = '#'
	} else {
		lines[currentLine][linePosition] = '.'
	}

}

func drawBoard(lines *[6][40]byte) {
	for _, line := range lines {
		for _, char := range line {
			fmt.Printf("%c", rune(char))
		}
		fmt.Printf("\n")
	}
}
