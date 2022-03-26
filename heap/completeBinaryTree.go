package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)
var stdin = os.Stdin

// var stdin2 = bytes.NewBufferString(`
// 5
// 7 8 1 2 3
// `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Heap []int

func (keys *Heap) id(idx int) int {
	return idx + 1
}

func (keys *Heap) key(idx int) int {
	return (*keys)[idx]
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

func (hp *Heap) PrintNode(idx int) {
	fmt.Fprintf(wr, "node %d: key = %d, ", hp.id(idx), hp.key(idx))

	if parentIndex(idx) >= 0 {
		fmt.Fprintf(wr, "parent key = %d, ", hp.key(parentIndex(idx)))
	}

	if leftIndex(idx) < len(*hp) {
		fmt.Fprintf(wr, "left key = %d, ", hp.key(leftIndex(idx)))
	}

	if rightIndex(idx) < len(*hp) {
		fmt.Fprintf(wr, "right key = %d, ", hp.key(rightIndex(idx)))
	}

	fmt.Fprintln(wr)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	hp := make(Heap, n)

	for i := 0; i < n; i++ {
		hp[i] = scanInt()
	}

	for i := 0; i < n; i++ {
		hp.PrintNode(i)
	}
}
