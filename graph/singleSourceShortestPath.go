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

const INFINITY = 100000000

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
			adj[i][j] = INFINITY
		}

		minWeight[i] = INFINITY
	}

	for i := 0; i < n; i++ {
		u := scanInt()
		m := scanInt()
		for j := 0; j < m; j++ {
			v := scanInt()
			c := scanInt()
			adj[u][v] = c
		}
	}

	seekSssp(n)
	printWeight(n)
}

func printWeight(n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintln(wr, i, minWeight[i])
	}
}

func seekSssp(n int) {
	var minCost, u int

	minWeight[0] = 0

	for {
		minCost = INFINITY

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
			return
		}

		didVisit[u] = true

		for i := 0; i < n; i++ {
			if didVisit[i] {
				continue
			}

			if minWeight[u]+adj[u][i] < minWeight[i] {
				minWeight[i] = minWeight[u] + adj[u][i]
			}
		}
	}
}
