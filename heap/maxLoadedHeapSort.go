package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)

// var stdin = os.Stdin

var stdin = bytes.NewBufferString(`
8
1 2 3 5 9 12 15 23
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
	s := make([]string, len(hp))

	for i := 0; i < len(hp); i++ {
		s[i] = fmt.Sprint(hp[i])
	}

	fmt.Println(strings.Join(s, " "))
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

func (hp *Heap) BuildMaxHeap() {
	for i := len(*hp)/2 - 1; i >= 0; i-- {
		hp.MaxHeapify(i)
	}
}

func (hp *Heap) MaxHeapify(idx int) {
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
		(*hp)[largestIndex], (*hp)[idx] = (*hp)[idx], (*hp)[largestIndex]
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

func (hp *Heap) HeapSort() *Heap {
	hp.BuildMaxHeap()

	newHp := NewHeap(len(*hp))
	size := len(*hp)
	for i := 0; i < size; i++ {
		(*newHp)[i] = hp.ExtractMax()
	}

	return newHp
}

func (hp *Heap) ExtractMax() int {
	mx := (*hp)[0]
	size := len(*hp)

	(*hp)[0] = (*hp)[size-1]
	*hp = (*hp)[:size-1]

	if size-1 > 0 {
		hp.MaxHeapify(0)
	}
	return mx
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()

	hp0 := make([]int, n)
	for i := 0; i < n; i++ {
		hp0[i] = scanInt()
	}

	sort.Ints(hp0)
	hp := NewHeap(n)

	for i := 0; i < n; i++ {
		(*hp)[i] = hp0[n-1-i]
	}

	for i, j := parentIndex(n-1), 1; i > 0 && hp.isValidIndex(j); i, j = parentIndex(i), j+1 {
		(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
	}

	hp.BuildMaxHeap()
	hp.Print()
}
