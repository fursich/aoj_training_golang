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

// var stdin = bytes.NewBufferString(
// 	`4
//    1 2 2 4
//    2 1 4
//    3 0
//    4 1 3
// `)

type Queue struct {
	baseList [MAX_QUEUE_SIZE]int
	head     int
	length   int
}

var queue *Queue
var visited []bool
var dist []int
var adj [][]int

const MAX_QUEUE_SIZE = 100000

func NewQueue() *Queue {
	q := Queue{head: 0, length: 0}
	return &q
}

func (q *Queue) Enqueue(v int) {
	if q.length >= MAX_QUEUE_SIZE {
		panic("too may elements. cannot enqueue any further")
	}

	idx := (q.head + q.length) % MAX_QUEUE_SIZE
	q.length++
	q.baseList[idx] = v
}

func (q *Queue) Dequeue() int {
	if q.length <= 0 {
		panic("Cannot find an element to dequeue")
	}

	idx := q.head
	q.head = (idx + 1) % MAX_QUEUE_SIZE
	q.length--
	return q.baseList[idx]
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	visited = make([]bool, n)
	dist = make([]int, n)
	adj = make([][]int, n)
	queue = NewQueue()

	for i := 0; i < n; i++ {
		scanInt() // not needed
		k := scanInt()
		adj[i] = make([]int, k)
		dist[i] = -1

		for j := 0; j < k; j++ {
			adj[i][j] = scanInt() - 1
		}
	}

	runBfs()
	printDistances(n)
}

func printDistances(n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintln(wr, i+1, dist[i])
	}
}

func runBfs() {
	visit(0, 0)

	for queue.length > 0 {
		i := queue.Dequeue()
		for _, j := range adj[i] {
			if visited[j] {
				continue
			}

			visit(j, dist[i]+1)
		}
	}
}

func visit(i int, d int) {
	visited[i] = true
	queue.Enqueue(i)
	dist[i] = d
}
