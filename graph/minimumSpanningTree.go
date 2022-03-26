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
//   -1 2 3 1 -1
//   2 -1 -1 4 -1
//   3 -1 -1 1 1
//   1 4 1 -1 3
//   -1 -1 1 3 -1
// `)

const INFINITY = 10000

var adj [][]int
var didVisit []bool
var minWeight []int

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())

	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()

	didVisit = make([]bool, n)
	adj = make([][]int, n)
	minWeight = make([]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)

		for j := 0; j < n; j++ {
			w := scanInt()
			if w < 0 {
				w = INFINITY
			}
			adj[i][j] = w
		}

		minWeight[i] = INFINITY
	}

	w := seekMst(n)
	fmt.Fprintln(wr, w)
}

func seekMst(n int) int {
	var u int
	var totalWeight int

	minWeight[0] = 0

	for {
		minCost := INFINITY
		for i := 0; i < n; i++ {
			if didVisit[i] {
				continue
			}
			if minWeight[i] < minCost {
				minCost = minWeight[i]
				u = i
			}
		}

		if minCost == INFINITY {
			return totalWeight
		}

		totalWeight += minCost
		didVisit[u] = true

		for i := 0; i < n; i++ {
			if didVisit[i] {
				continue
			}
			if adj[u][i] < minWeight[i] {
				minWeight[i] = adj[u][i]
			}
		}
	}
}
