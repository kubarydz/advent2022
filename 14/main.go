package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 24 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	input = utils.ReadInput("sample")
	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 93 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)

}

type point struct {
	x, y int
}

func calculateV1(input []byte) int {
	cave, lowest := buildCave(input)

	source := point{500, 0}
	sandCounter := 0
	sand := source
	for {
		if sand.y > lowest {
			break
		}

		_, ok := cave[point{x: sand.x, y: sand.y + 1}]
		if !ok {
			sand.y++
			continue
		}

		if _, leftOk := cave[point{x: sand.x - 1, y: sand.y + 1}]; !leftOk {
			sand.x--
			continue
		}
		if _, rightOk := cave[point{x: sand.x + 1, y: sand.y + 1}]; !rightOk {
			sand.x++
			continue
		}
		cave[sand] = struct{}{}
		sandCounter++
		sand = source
	}

	return sandCounter
}

func strToPoint(coords string) point {
	xy := strings.Split(coords, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		panic(err)
	}
	return point{
		x: x,
		y: y,
	}
}

func getLineRange(a, b int) []int {
	line := []int{}
	small := a
	big := b
	if a > b {
		small = b
		big = a

	}
	for i := small; i <= big; i++ {
		line = append(line, i)
	}
	return line
}

func calculateV2(input []byte) int {
	cave, lowest := buildCave(input)

	floor := lowest + 2

	source := point{500, 0}
	sandCounter := 0
	sand := source
	for {
		if sand.y+1 >= floor {
			cave[point{x: sand.x, y: sand.y + 1}] = struct{}{}
			cave[point{x: sand.x - 1, y: sand.y + 1}] = struct{}{}
			cave[point{x: sand.x + 1, y: sand.y + 1}] = struct{}{}
		}
		_, ok := cave[point{x: sand.x, y: sand.y + 1}]
		if !ok {
			sand.y++
			continue
		}

		if _, leftOk := cave[point{x: sand.x - 1, y: sand.y + 1}]; !leftOk {
			sand.x--
			continue
		}
		if _, rightOk := cave[point{x: sand.x + 1, y: sand.y + 1}]; !rightOk {
			sand.x++
			continue
		}
		cave[sand] = struct{}{}
		sandCounter++
		if sand.y == 0 {
			break
		}
		sand = source
	}

	return sandCounter
}

func buildCave(input []byte) (map[point]struct{}, int) {
	chunked := utils.ChunkInput(input)

	cave := map[point]struct{}{}

	lowest := 0
	for _, line := range chunked {
		points := strings.Split(string(line), " -> ")
		prev := strToPoint(points[0])
		for _, pStr := range points[1:] {
			p := strToPoint(pStr)
			if p.y > lowest {
				lowest = p.y
			}
			if prev.x == p.x {
				for _, y := range getLineRange(p.y, prev.y) {
					cave[point{x: p.x, y: y}] = struct{}{}
				}
			}
			if prev.y == p.y {
				for _, x := range getLineRange(p.x, prev.x) {
					cave[point{x: x, y: p.y}] = struct{}{}
				}
			}
			prev = p
		}
	}
	return cave, lowest
}
