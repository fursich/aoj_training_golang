package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)

var stdin = os.Stdin

// var stdin = bytes.NewBufferString(`
// 10
// 4 1 3 2 16 9 10 14 8 7
// `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Heap []int

func (hp Heap) Print() {
	for i := 0; i < len(hp); i++ {
		fmt.Printf(" %d", hp[i])
	}
	fmt.Println()
}

func (hp Heap) key(idx int) int {
	return hp[idx]
}

func (hp Heap) isValidIndex(idx int) bool {
	return 0 <= idx && idx < len(hp)
}

func id(idx int) int {
	return idx + 1
}

func parentIndex(idx int) int {
	return (idx+1)/2 - 1
}

func leftIndex(idx int) int {
	return (idx+1)*2 - 1
}

func rightIndex(idx int) int {
	return (idx + 1) * 2
}

func (hp Heap) BuildMaxHeap() {
	for i := len(hp)/2 - 1; i >= 0; i-- {
		hp.MaxHeapify(i)
	}
}

func (hp Heap) MaxHeapify(idx int) {
	var largestIndex int
	l := leftIndex(idx)
	r := rightIndex(idx)

	if hp.isValidIndex(l) && hp.key(l) > hp.key(idx) {
		largestIndex = l
	} else {
		largestIndex = idx
	}
	if hp.isValidIndex(r) && hp.key(r) > hp.key(largestIndex) {
		largestIndex = r
	}

	if largestIndex != idx {
		hp[largestIndex], hp[idx] = hp[idx], hp[largestIndex]
		hp.MaxHeapify(largestIndex)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	hp := make(Heap, n)

	for i := 0; i < n; i++ {
		hp[i] = scanInt()
	}

	hp.BuildMaxHeap()
	hp.Print()
}
