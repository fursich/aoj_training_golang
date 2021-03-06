package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(rd)
var wr = bufio.NewWriter(os.Stdout)

var rd = os.Stdin

// var rd = bytes.NewBufferString(
// 	`6
//    2 1
//    2 2
//    4 2
//    6 2
//    3 3
//    5 4
//    2
//    2 4 0 4
//    4 10 2 5
//   `)

type Node struct {
	point *Point
	left  int
	right int
}

type Tree []*Node

func NewTree(n int) *Tree {
	var t Tree = make([]*Node, n)

	for i := 0; i < n; i++ {
		t[i] = &Node{left: -1, right: -1}
	}

	return &t
}

func (t *Tree) index(id int) *Node {
	return (*t)[id]
}

var ans []int

func (t *Tree) find(id int, sx, tx, sy, ty int, depth int) {
	if id < 0 {
		return
	}

	pt := t.index(id).point
	x := pt.x
	y := pt.y

	if sx <= x && x <= tx && sy <= y && y <= ty {
		ans = append(ans, pt.id)
	}

	if depth%2 == 0 {
		if sx <= x {
			t.find(t.index(id).left, sx, tx, sy, ty, depth+1)
		}
		if x <= tx {
			t.find(t.index(id).right, sx, tx, sy, ty, depth+1)
		}
	} else {
		if sy <= y {
			t.find(t.index(id).left, sx, tx, sy, ty, depth+1)
		}
		if y <= ty {
			t.find(t.index(id).right, sx, tx, sy, ty, depth+1)
		}
	}
}

var idx int = 0

func (t *Tree) buildKDTreeX(pts []Point, l, r int) int {
	if l >= r {
		return -1
	}

	mid := (l + r) / 2
	i := idx
	idx++

	sx := SorterX{
		points: pts[l:r],
		length: r - l,
	}
	sort.Sort(&sx)

	(*t)[i].point = &pts[mid]
	(*t)[i].left = t.buildKDTreeY(pts, l, mid)
	(*t)[i].right = t.buildKDTreeY(pts, mid+1, r)

	return i
}

func (t *Tree) buildKDTreeY(pts []Point, l, r int) int {
	if l >= r {
		return -1
	}

	mid := (l + r) / 2
	i := idx
	idx++

	sy := SorterY{
		points: pts[l:r],
		length: r - l,
	}
	sort.Sort(&sy)

	(*t)[i].point = &pts[mid]
	(*t)[i].left = t.buildKDTreeX(pts, l, mid)
	(*t)[i].right = t.buildKDTreeX(pts, mid+1, r)

	return i
}

type Point struct {
	id int
	x  int
	y  int
}

type SorterX struct {
	points []Point
	length int
}

func (sx *SorterX) Len() int {
	return sx.length
}

func (sx *SorterX) Swap(i, j int) {
	sx.points[i], sx.points[j] = sx.points[j], sx.points[i]
}

func (sx *SorterX) Less(i, j int) bool {
	return sx.points[i].x < sx.points[j].x
}

type SorterY struct {
	points []Point
	length int
}

func (sy *SorterY) Len() int {
	return sy.length
}

func (sy *SorterY) Swap(i, j int) {
	sy.points[i], sy.points[j] = sy.points[j], sy.points[i]
}

func (sy *SorterY) Less(i, j int) bool {
	return sy.points[i].y < sy.points[j].y
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Buffer(make([]byte, 10000000), math.MaxInt64)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	points := make([]Point, n)
	tree := NewTree(n)

	for i := 0; i < n; i++ {
		points[i].id = i
		points[i].x = scanInt()
		points[i].y = scanInt()
	}

	root := tree.buildKDTreeX(points, 0, n)

	q := scanInt()
	ans = make([]int, 0, n)
	for i := 0; i < q; i++ {
		sx, tx, sy, ty := scanInt(), scanInt(), scanInt(), scanInt()

		tree.find(root, sx, tx, sy, ty, 0)
		sort.Ints(ans)
		printResult()
		ans = ans[0:0]
	}
}

func printResult() {
	for i := 0; i < len(ans); i++ {
		fmt.Fprintln(wr, ans[i])
	}
	fmt.Fprintln(wr)
}
