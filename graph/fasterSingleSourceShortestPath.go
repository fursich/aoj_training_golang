package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(rd)
var wr = bufio.NewWriter(os.Stdout)

var rd = os.Stdin

// var rd = bytes.NewBufferString(
// 	`5
//   0 3 2 3 3 1 1 2
//   1 2 0 2 3 4
//   2 3 0 3 3 1 4 1
//   3 4 2 1 0 1 1 4 4 3
//   4 2 2 1 3 3
// `)

type PriorityQueue []*Vertex
type Vertex struct {
	id         int
	distance   int
	queueIndex int
}

type Edge struct {
	id       int
	distance int
}

var pq *PriorityQueue
var adj [][]*Edge
var vertices []*Vertex
var didVisit []bool

func NewPriorityQueue(n int) *PriorityQueue {
	pq := make(PriorityQueue, n)
	return &pq
}

func NewVertex(id int, dist int, queueIndex int) *Vertex {
	return &Vertex{id: id, distance: dist, queueIndex: queueIndex}
}

func (v *Vertex) LessThan(u *Vertex) bool {
	return v.distance < u.distance
}

func parentIndex(n int) int {
	return (n - 1) / 2
}

func leftIndex(n int) int {
	return n*2 + 1
}

func rightIndex(n int) int {
	return n*2 + 2
}

func (pq *PriorityQueue) Top() *Vertex {
	return (*pq)[0]
}

func (pq *PriorityQueue) Buttom() *Vertex {
	return (*pq)[pq.Size()-1]
}

func (pq *PriorityQueue) Size() int {
	return len(*pq)
}

func (pq *PriorityQueue) LessThan(i int, j int) bool {
	return (*pq)[i].LessThan((*pq)[j])
}

func (pq *PriorityQueue) Swap(i int, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]

	(*pq)[i].queueIndex, (*pq)[j].queueIndex = (*pq)[j].queueIndex, (*pq)[i].queueIndex
}

func (pq *PriorityQueue) ExtractMin() *Vertex {
	if pq.Size() < 1 {
		panic("Heap underflow")
	}

	min := pq.Top()
	pq.Swap(0, pq.Size()-1)
	*pq = (*pq)[0 : pq.Size()-1]

	pq.MinHeapify(0)
	return min
}

func (pq *PriorityQueue) DecreaseKey(i int, dist int) {
	if dist > (*pq)[i].distance {
		panic("current key is smaller than the given key")
	}

	(*pq)[i].distance = dist

	for i > 0 {
		j := parentIndex(i)
		if pq.LessThan(j, i) {
			break
		}

		pq.Swap(i, j)
		i = j
	}
}

func (pq *PriorityQueue) BuildMinHeap() {
	for i := pq.Size()/2 - 1; i >= 0; i-- {
		pq.MinHeapify(i)
	}
}

func (pq *PriorityQueue) MinHeapify(i int) {
	l := leftIndex(i)
	r := rightIndex(i)

	smallest := i
	if l < pq.Size() && pq.LessThan(l, i) {
		smallest = l
	}
	if r < pq.Size() && pq.LessThan(r, smallest) {
		smallest = r
	}

	if smallest == i {
		return
	}

	pq.Swap(i, smallest)
	pq.MinHeapify(smallest)
}

const INFINITY = (1 << 63) - 1

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())

	return i
}

func main() {
	var v *Vertex

	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()

	pq = NewPriorityQueue(n)
	vertices = make([]*Vertex, n)
	didVisit = make([]bool, n)
	adj = make([][]*Edge, n)

	for i := 0; i < n; i++ {
		id := scanInt()
		m := scanInt()

		if id == 0 {
			v = NewVertex(id, 0, 0)
		} else {
			v = NewVertex(id, INFINITY, id)
		}
		(*pq)[id] = v
		vertices[id] = v

		adj[id] = make([]*Edge, m)
		for j := 0; j < m; j++ {
			adj[id][j] = &Edge{id: scanInt(), distance: scanInt()}
		}
	}
	pq.BuildMinHeap()

	calculateSssp(n)

	for _, v := range vertices {
		fmt.Fprintln(wr, v.id, v.distance)
	}
}

func calculateSssp(n int) {
	for pq.Size() > 0 {
		u := pq.ExtractMin()
		didVisit[u.id] = true

		for _, v := range adj[u.id] {
			if didVisit[v.id] {
				continue
			}

			d := vertices[u.id].distance + v.distance
			if d < vertices[v.id].distance {
				idx := vertices[v.id].queueIndex
				pq.DecreaseKey(idx, d)
			}
		}
	}
}
