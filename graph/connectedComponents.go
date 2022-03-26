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
// 	`10 9
//    0 1
//    0 2
//    3 4
//    5 7
//    5 6
//    6 7
//    6 8
//    7 8
//    8 9
//    3
//    0 1
//    5 9
//    1 3
//    `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())

	return i
}

type Adj []map[int]bool

var adj *Adj
var colors []int

func NewAdj(n int) *Adj {
	adj := make(Adj, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}

	return &adj
}

func IsConnected(i int, j int) bool {
	if colors[i] == 0 || colors[j] == 0 {
		return false
	}
	return colors[i] == colors[j]
}

func (adj *Adj) Connect(i int, j int) {
	(*adj)[i][j] = true
	(*adj)[j][i] = true
}

func buildAdj(n int, m int) {
	adj = NewAdj(n)

	for i := 0; i < m; i++ {
		p := scanInt()
		q := scanInt()
		adj.Connect(p, q)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	m := scanInt()
	buildAdj(n, m)

	colors = make([]int, n)

	paintColors(n)
	assessConnections()
}

func assessConnections() {
	n := scanInt()
	for i := 0; i < n; i++ {
		u := scanInt()
		v := scanInt()
		if IsConnected(u, v) {
			fmt.Fprintln(wr, "yes")
		} else {
			fmt.Fprintln(wr, "no")
		}
	}
}

func paintColors(n int) {
	for i := 0; i < n; i++ {
		visit(i, i+1)
	}
}

func visit(i int, cl int) {
	if didVisit(i) {
		return
	}
	colors[i] = cl

	for j := range (*adj)[i] {
		if !didVisit(j) {
			visit(j, cl)
		}
	}
}

func didVisit(i int) bool {
	return colors[i] != 0
}
