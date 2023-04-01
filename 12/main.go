package main

import (
	"fmt"

	"github.com/kubarydz/advent2022/utils"
)

func main() {
	input := utils.ReadInput("sample")

	fmt.Printf("samples\n")
	resp := calculateV1(input)
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 31 {
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

type point struct {
	x, y int
}

var moves = []point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func calculateV1(input []byte) int {
	chunked := utils.ChunkInput(input)

	start := findPosition('S', chunked)
	chunked[start.x][start.y] = 'a'
	end := findPosition('E', chunked)
	chunked[end.x][end.y] = 'z'

	seen := map[point]int{start: 0}

	queue := NewQueue(start)

	for queue.Size() != 0 {
		current := queue.PopLast()
		for _, move := range moves {
			movePoint := point{current.x + move.x, current.y + move.y}
			if !validMove(current, movePoint, chunked) {
				continue
			}
			if dist, ok := seen[movePoint]; !ok || dist > seen[current]+1 {
				seen[movePoint] = seen[current] + 1
				queue.Put(movePoint)
			}
		}
	}
	return seen[end]
}

func validMove(current, p point, input [][]byte) bool {
	if p.x < 0 || p.x > len(input)-1 || p.y < 0 || p.y > len(input[0])-1 {
		return false
	}
	if rune(input[current.x][current.y])-rune(input[p.x][p.y]) < -1 {
		return false
	}
	return true
}

func calculateV2(input []byte) int {
	return 0
}

func findPosition(symbol byte, input [][]byte) point {
	for i, row := range input {
		for j, val := range row {
			if val == symbol {
				return point{i, j}
			}
		}
	}
	return point{}
}

type Queue[T any] struct {
	First *QueueNode[T]
	Last  *QueueNode[T]
}

type QueueNode[T any] struct {
	Previous *QueueNode[T]
	Next     *QueueNode[T]

	Val T
}

func (q *Queue[T]) PopLast() T {
	last := q.Last
	q.Last = q.Last.Previous
	if q.Last != nil {
		q.Last.Next = nil
	}
	return last.Val
}

func (q *Queue[T]) Put(elem T) {
	node := QueueNode[T]{
		Previous: q.Last,
		Val:      elem,
	}
	if q.Last == nil {
		q.First = &node
		q.Last = &node
		return
	}
	q.Last.Next = &node
	q.Last = &node
}

func (q *Queue[T]) Size() int {
	if q.Last == nil {
		return 0
	}
	s := 0
	for n := q.First; n != nil; n = n.Next {
		s++
	}
	return s
}

func NewQueue[T any](elem T) *Queue[T] {
	node := QueueNode[T]{
		Val: elem,
	}
	return &Queue[T]{
		First: &node,
		Last:  &node,
	}
}
