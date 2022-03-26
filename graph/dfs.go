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
// 	`6
//    1 2 2 4
//    2 1 5
//    3 2 5 6
//    4 0
//    5 1 4
//    6 1 6
//   `)

// var stdin = bytes.NewBufferString(
// 	`4
//    1 1 2
//    2 1 4
//    3 0
//    4 1 3
//   `)

// var stdin = bytes.NewBufferString(
// 	`6
//    1 2 2 3
//    2 2 3 4
//    3 1 5
//    4 1 6
//    5 1 6
//    6 0
//   `)

const (
	UNKNOWN = iota
	FOUND
	FINISHED
)
const MAX_STACK_LENGTH = 10000

type Stack []int

var st = NewStack()
var clock int = 1

func NewStack() *Stack {
	s := make(Stack, 0, MAX_STACK_LENGTH)
	return &s
}

func (s *Stack) pop() int {
	if s.size() == 0 {
		panic("stack is empty. cannot pop any further")
	}

	v := (*s)[s.size()-1]
	*s = (*s)[:s.size()-1]

	return v
}

func (s *Stack) push(v int) {
	*s = append(*s, v)
}

func (s *Stack) top() int {
	return (*s)[s.size()-1]
}

func (s *Stack) size() int {
	return len(*s)
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func printTime(time [][]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintf(wr, "%d %d %d\n", i+1, time[i][0], time[i][1])
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		scanInt() // id (not neccesary)
		k := scanInt()

		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			adj[i][j] = scanInt() - 1
		}
	}

	time := visitVertices(adj, n)
	printTime(time, n)
}

func visitVertices(adj [][]int, n int) [][]int {

	time := make([][]int, n)
	status := make([]int, n)

	for i := 0; i < n; i++ {
		status[i] = UNKNOWN
		time[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		dfsVisit(i, adj, time, status)
	}

	return time
}

func dfsVisit(i int, adj [][]int, time [][]int, status []int) {
	if status[i] != UNKNOWN {
		return
	}

	st.push(i)
	for {
		if st.size() == 0 {
			break
		}
		p := st.top()

		if status[p] == UNKNOWN {
			status[p] = FOUND
			time[p][0] = clock
			clock++
		}

		if len(adj[p]) == 0 {
			st.pop()
			status[p] = FINISHED
			time[p][1] = clock
			clock++
			continue
		}

		q := adj[p][0]
		adj[p] = adj[p][1:]

		if status[q] == UNKNOWN {
			st.push(q)
		}
	}
}
