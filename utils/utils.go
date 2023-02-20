package utils

import (
	"fmt"
	"os"
)

func ReadInput(filename string) []byte {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("cannot open %v, error: %v\n", filename, err)
		panic("cannot open file")
	}
	return bytes
}

func ChunkInput(input []byte) [][]byte {
	var chunks [][]byte
	lastStart := 0
	for i, b := range input {
		if b == '\n' {
			chunks = append(chunks, input[lastStart:i])
			lastStart = i + 1

		}
	}
	return chunks
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
