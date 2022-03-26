package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)

// var stdin = os.Stdin

var stdin = bytes.NewBufferString(`
insert 8
insert 2
extract
insert 10
extract
insert 11
extract
extract
end
`)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

type Heap []int

const MAXHEAPSIZE = 2000001

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

func NewHeap(n int) *Heap {
	hp := make(Heap, n, MAXHEAPSIZE)
	return &hp
}

func ExecuteNext(hp *Heap) bool {
	cmd := scanString()

	switch cmd {
	case "insert":
		hp.Insert(scanInt())
	case "extract":
		fmt.Fprintln(wr, hp.ExtractMax())
	case "end":
		return false
	default:
		panic("unknown command: " + cmd)
	}

	return true
}

func (hp *Heap) Insert(key int) {
	*hp = append(*hp, -1)
	hp.IncreaseKey(len(*hp)-1, key)
}

func (hp Heap) IncreaseKey(idx int, key int) {
	if key < hp.key(idx) {
		panic("cannot replace larger key " + fmt.Sprint(hp.key(idx)) + " with new key: " + fmt.Sprint(key))
	}
	hp[idx] = key

	for hp.isValidIndex(parentIndex(idx)) && hp.key(parentIndex(idx)) < hp.key(idx) {
		hp[idx], hp[parentIndex(idx)] = hp[parentIndex(idx)], hp[idx]
		idx = parentIndex(idx)
	}
}

func (hp *Heap) ExtractMax() int {
	mx := (*hp)[0]
	size := len(*hp)

	(*hp)[0] = (*hp)[size-1]
	*hp = (*hp)[:size-1]

	hp.MaxHeapify(0)
	return mx
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	hp := NewHeap(0)

	for {
		if !ExecuteNext(hp) {
			break
		}
	}
}
