package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	resp := calculateV1(input, "AA")
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 1651 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV1(input, "AA")
	fmt.Printf("input 1 solution: %d\n", resp)

	input = utils.ReadInput("sample")
	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 56000011 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)

}

func calculateV1(input []byte, firstValve string) int {
	lines := utils.ChunkInput(input)
	valves := getValves(lines)

	timer := 30

	done := make(chan bool)
	finalPressure := make(chan int)
	biggestPressure := 0

	nonZeroValves := 0
	for _, valve := range valves {
		if valve.flowRate > 0 {
			nonZeroValves++
		}
	}

	go func() {
		for {
			select {
			case p := <-finalPressure:
				if p > biggestPressure {
					biggestPressure = p
				}
			case <-done:
				return
			}
		}
	}()

	makeWay(valves[firstValve], timer, valves, []*valve{}, 0, finalPressure, nonZeroValves)
	done <- true

	return biggestPressure
}

func makeWay(current *valve, timer int, valves map[string]*valve, openValves []*valve, pressure int, finalPressure chan int, nonZeroValves int) {
	// fail-fast heuristics
	if timer < 20 && pressure < 50 {
		return
	}
	if timer < 25 && pressure < 10 {
		return
	}
	if timer < 10 && pressure < 500 {
		return
	}
	if timer < 5 && pressure < 800 {
		return
	}
	if timer < 2 && pressure < 1200 {
		return
	}

	if timer == 0 {
		finalPressure <- pressure
		return
	}

	isCurrentOpen := false
	pressureToAdd := 0
	for _, valve := range openValves {
		if valve == current {
			isCurrentOpen = true
		}
		pressureToAdd += valve.flowRate
	}
	pressure += pressureToAdd

	if !isCurrentOpen && current.flowRate >= 10 && timer >= 2 {
		newOpen := append([]*valve(nil), openValves...)
		newOpen = append(newOpen, current)

		pressure += current.flowRate
		pressure += pressureToAdd
		for _, tunnel := range current.tunnels {
			makeWay(tunnel, timer-2, valves, newOpen, pressure, finalPressure, nonZeroValves)
		}
		return
	}

	// fast track if all non-zero valves are open
	if len(openValves) == nonZeroValves {
		pressure += pressureToAdd * (timer - 1)
		finalPressure <- pressure
		return
	}

	timer--
	for _, tunnel := range current.tunnels {
		makeWay(tunnel, timer, valves, openValves, pressure, finalPressure, nonZeroValves)
	}
	if timer == 0 || isCurrentOpen || current.flowRate == 0 {
		return
	}

	timer--
	newOpen := append([]*valve(nil), openValves...)
	newOpen = append(newOpen, current)
	pressure += current.flowRate
	pressure += pressureToAdd

	for _, tunnel := range current.tunnels {
		makeWay(tunnel, timer, valves, newOpen, pressure, finalPressure, nonZeroValves)
	}
}

func getValves(input [][]byte) map[string]*valve {
	m := make(map[string]*valve, len(input))
	for _, line := range input {
		id := string(line[6:8])
		rate := line[23:25]
		if rate[1] == ';' {
			rate = rate[0:1]
		}
		rateInt, err := strconv.Atoi(string(rate))
		if err != nil {
			panic(err)
		}
		m[id] = &valve{id: id, flowRate: rateInt, tunnels: []*valve{}}
	}

	for _, line := range input {
		id := string(line[6:8])
		valves := strings.Split(string(line), "valve")[1]
		if valves[0] == 's' {
			valves = valves[1:]
		}
		valvesSplit := strings.Split(valves, ",")
		for _, v := range valvesSplit {
			v = strings.TrimSpace(v)
			tunnelFrom := m[id]
			tunnelFrom.tunnels = append(tunnelFrom.tunnels, m[v])
		}
	}
	return m
}

func calculateV2(input []byte) int {
	return 0
}

type valve struct {
	id       string
	flowRate int
	tunnels  []*valve
}
