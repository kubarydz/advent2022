package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	resp := calculateV1(input, 20)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 26 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV1(input, 2000000)
	fmt.Printf("input 1 solution: %d\n", resp)

	input = utils.ReadInput("sample")
	resp = calculateV2(input, 20)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 56000011 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV2(input, 4000000)
	fmt.Printf("input 2 solution: %d\n", resp)

}

type point struct {
	x, y int
}

func calculateV1(input []byte, rowNr int) int {
	chunked := utils.ChunkInput(input)

	noBeacon := map[int]struct{}{}

	for _, row := range chunked {
		sensor, beacon := getPoints(row)
		distance := manhattanDistance(sensor, beacon)
		distToRow := dist(sensor.y, rowNr)
		safeRowSpread := distance - distToRow
		if safeRowSpread <= 0 {
			continue
		}
		for i := sensor.x - safeRowSpread; i < sensor.x+safeRowSpread; i++ {
			noBeacon[i] = struct{}{}
		}
	}

	return len(noBeacon)
}

func calculateV2(input []byte, rowNr int) int {
	chunked := utils.ChunkInput(input)

	shapes := map[point]int{}

	for _, row := range chunked {
		sensor, beacon := getPoints(row)
		distance := manhattanDistance(sensor, beacon)
		shapes[sensor] = distance

	}

	for shape, distance := range shapes {
		// build perimiter just outside of sensor reach. There is only one point not in sensors range, so it has to be just outside of at least one of them
		perimiter := []point{}

		for i := distance; i >= -1; i-- {
			y := shape.y + distance - i
			x := shape.x - i - 1
			if x < 0 || x > rowNr || y < 0 || y >= rowNr {
				continue
			}
			perimiter = append(perimiter, point{x: x, y: y})
		}

		for i := distance - 1; i >= 0; i-- {
			y := shape.y + distance - i
			x := shape.x + i + 1
			if x < 0 || x > rowNr || y < 0 || y >= rowNr {
				continue
			}
			perimiter = append(perimiter, point{x: x, y: y})
		}

		for i := distance - 1; i >= -1; i-- {
			y := shape.y - (distance - i)
			x := shape.x - i - 1
			if x < 0 || x > rowNr || y < 0 || y >= rowNr {
				continue
			}
			perimiter = append(perimiter, point{x: x, y: y})
		}

		for i := distance; i >= 0; i-- {
			y := shape.y - (distance - i)
			x := shape.x + i + 1
			if x < 0 || x > rowNr || y < 0 || y >= rowNr {
				continue
			}
			perimiter = append(perimiter, point{x: x, y: y})
		}

		for _, p := range perimiter {
			found := false
			for checkShape, checkDist := range shapes {
				if checkShape == shape {
					continue
				}

				if checkShape.y-(checkDist-dist(checkShape.x, p.x)) <= p.y &&
					checkShape.y+(checkDist-dist(checkShape.x, p.x)) >= p.y {
					found = true
					break
				}
			}

			if !found {
				return p.x*4000000 + p.y
			}
		}
	}

	return 0
}

func getPoints(row []byte) (point, point) {
	stringRow := string(row[12:])
	split := strings.Split(stringRow, ": closest beacon is at x=")
	sensorRow := strings.Split(split[0], ", y=")
	sensorX, err := strconv.Atoi(sensorRow[0])
	if err != nil {
		panic(err)
	}
	sensorY, err := strconv.Atoi(sensorRow[1])
	if err != nil {
		panic(err)
	}

	beaconRow := strings.Split(split[1], ", y=")
	beaconX, err := strconv.Atoi(beaconRow[0])
	if err != nil {
		panic(err)
	}
	beaconY, err := strconv.Atoi(beaconRow[1])
	if err != nil {
		panic(err)
	}

	return point{x: sensorX, y: sensorY}, point{x: beaconX, y: beaconY}
}

func manhattanDistance(a, b point) int {
	return dist(a.x, b.x) + dist(a.y, b.y)
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
