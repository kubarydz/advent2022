package main

import (
	"encoding/json"
	"fmt"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 13 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV1(input)
	fmt.Printf("input 1 solution: %d\n", resp)

	input = utils.ReadInput("sample")
	resp = calculateV2(input)
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 140 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	input = utils.ReadInput("input")
	resp = calculateV2(input)
	fmt.Printf("input 2 solution: %d\n", resp)

}

func calculateV1(input []byte) int {
	chunked := utils.ChunkInput(input)

	res := 0
	for i := 0; i < len(chunked); i += 3 {
		var left, right any
		json.Unmarshal(chunked[i], &left)
		json.Unmarshal(chunked[i+1], &right)
		if compare(left, right) < 0 {
			res += i/3 + 1
		}
	}

	return res
}

func compare(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if res := compare(l[i], r[i]); res != 0 {
			return res
		}
	}

	return len(l) - len(r)
}

func calculateV2(input []byte) int {
	var divider2, divider6 any
	json.Unmarshal([]byte("[[2]]"), &divider2)
	json.Unmarshal([]byte("[[6]]"), &divider6)

	chunked := utils.ChunkInput(input)
	index2 := 1
	index6 := 2
	for i := 0; i < len(chunked); i++ {
		if len(chunked[i]) == 0 {
			continue
		}
		var left any
		json.Unmarshal(chunked[i], &left)
		if compare(left, divider2) < 0 {
			index2++
		}
		if compare(left, divider6) < 0 {
			index6++
		}
	}
	return index2 * index6
}
