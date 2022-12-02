package main

import (
	"fmt"
	"github.com/kubarydz/advent2022/utils"
	"strconv"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	highestElf, highestCalories := getHighestElf(input)
	fmt.Printf("elf %v has %v calories\n", highestElf, highestCalories)
	if highestCalories != 24000 {
		fmt.Printf("test nr 1 failed\n")
		return
	}
	te, hc := getTopElfs(input)
	fmt.Printf("elfs %d, %d, %d have combined calories of %d\n", te[0], te[1], te[2], hc)
	if hc != 45000 {
		fmt.Printf("test nr 2 failed\n")
		return

	}

	fmt.Printf("inputs\n")
	input = utils.ReadInput("input")
	highestElf, highestCalories = getHighestElf(input)
	fmt.Printf("elf %v has %v calories\n", highestElf, highestCalories)
	te, hc = getTopElfs(input)
	fmt.Printf("elfs %d, %d, %d have combined calories of %d\n", te[0], te[1], te[2], hc)
}

func getHighestElf(input []byte) (int, int) {
	elfs := map[int]int{}
	elfCounter := 1
	previousChar := '1'
	calsLenght := 0
	calsBuffer := ""
	for i, b := range input {
		if b == '\n' || i == (len(input)-1) {
			if previousChar == '\n' {
				elfCounter++
				calsLenght = 0
				continue
			}
			elf := elfs[elfCounter]
			newCalories, _ := strconv.Atoi(calsBuffer)
			elf += newCalories
			elfs[elfCounter] = elf
			calsLenght = 0
			previousChar = rune(b)
			calsBuffer = ""
			continue
		}
		calsBuffer += string(b)
		calsLenght++
		previousChar = rune(b)
	}

	highestCalories := 0
	highestElf := 0
	for e, c := range elfs {
		if c > highestCalories {
			highestCalories = c
			highestElf = e
		}
	}
	return highestElf, highestCalories
}

func getTopElfs(input []byte) ([3]int, int) {
	elfs := map[int]int{}
	elfCounter := 1
	previousChar := '1'
	calsBuffer := ""
	for i, b := range input {
		if b == '\n' || i == (len(input)-1) {
			if previousChar == '\n' {
				elfCounter++
				continue
			}
			elf := elfs[elfCounter]
			newCalories, _ := strconv.Atoi(calsBuffer)
			elf += newCalories
			elfs[elfCounter] = elf
			previousChar = rune(b)
			calsBuffer = ""
			continue
		}
		calsBuffer += string(b)
		previousChar = rune(b)
	}

	highestCalories := [3]int{0, 0, 0}
	highestElf := [3]int{0, 0, 0}
	for e, c := range elfs {
		if c > highestCalories[0] {
			highestElf[2] = highestElf[1]
			highestCalories[2] = highestCalories[1]

			highestElf[1] = highestElf[0]
			highestCalories[1] = highestCalories[0]

			highestCalories[0] = c
			highestElf[0] = e
			continue
		}

		if c > highestCalories[1] {
			highestElf[2] = highestElf[1]
			highestCalories[2] = highestCalories[1]

			highestCalories[1] = c
			highestElf[1] = e
			continue
		}

		if c > highestCalories[2] {
			highestCalories[2] = c
			highestElf[2] = e
			continue
		}
	}
	return highestElf, highestCalories[0] + highestCalories[1] + highestCalories[2]
}
