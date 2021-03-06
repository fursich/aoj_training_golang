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
// 	`5 12
//    0 1 4
//    0 2 3
//    1 1 2
//    1 3 4
//    1 1 4
//    1 3 2
//    0 1 3
//    1 2 4
//    1 3 0
//    0 0 4
//    1 0 2
//    1 3 0
//   `)

type DisjointSets struct {
	parent []int
	rank   []int
}

func NewDisjointSets(n int) *DisjointSets {
	p := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
	}

	return &DisjointSets{parent: p, rank: r}
}

func (ds *DisjointSets) FindSet(x int) int {
	for ds.parent[x] != x {
		x = ds.parent[x]
	}
	return x
}

func (ds *DisjointSets) Same(x int, y int) bool {
	return ds.FindSet(x) == ds.FindSet(y)
}

func (ds *DisjointSets) Unite(x int, y int) {
	ds.link(ds.FindSet(x), ds.FindSet(y))
}

func (ds *DisjointSets) link(x int, y int) {
	if ds.rank[x] > ds.rank[y] {
		ds.parent[y] = x
	} else {
		ds.parent[x] = y
		if ds.rank[x] == ds.rank[y] {
			ds.rank[y]++
		}
	}
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func execCommand(ds *DisjointSets) {
	com := scanInt()
	x, y := scanInt(), scanInt()

	switch com {
	case 0:
		ds.Unite(x, y)
	case 1:
		if ds.Same(x, y) {
			fmt.Fprintln(wr, 1)
		} else {
			fmt.Fprintln(wr, 0)
		}
	default:
		panic("unknown commands")
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, q := scanInt(), scanInt()
	ds := NewDisjointSets(n)

	for i := 0; i < q; i++ {
		execCommand(ds)
	}
}
