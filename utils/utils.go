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
