package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)
var stdin = os.Stdin
var wr = bufio.NewWriter(os.Stdout)

// var stdin2 = bytes.NewBufferString(`
// 10
// insert 30
// insert 88
// insert 12
// insert 1
// insert 20
// find 12
// insert 17
// insert 25
// find 16
// print
// `)

func scanStr() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanCommand() {
	s := scanStr()
	switch s {
	case "insert":
		k := scanInt()
		root = root.Insert(k)
	case "print":
		root.PrintInOrdered()
		fmt.Fprintln(wr)
		root.PrintPreOrdered()
		fmt.Fprintln(wr)
	case "find":
		k := scanInt()
		if root.Find(k) {
			fmt.Fprintln(wr, "yes")
		} else {
			fmt.Fprintln(wr, "no")
		}
	default:
		panic("cannot interpret given command: " + s)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	for i := 0; i < n; i++ {
		scanCommand()
	}
}

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

var root *Node = nil

func NewNode(key int, p *Node) *Node {
	return &Node{key: key, parent: p}
}

func (nd *Node) PrintPreOrdered() {
	fmt.Fprintf(wr, " %d", nd.key)

	if nd.left != nil {
		nd.left.PrintPreOrdered()
	}
	if nd.right != nil {
		nd.right.PrintPreOrdered()
	}
}

func (nd *Node) PrintInOrdered() {
	if nd.left != nil {
		nd.left.PrintInOrdered()
	}

	fmt.Fprintf(wr, " %d", nd.key)

	if nd.right != nil {
		nd.right.PrintInOrdered()
	}
}

func (rt *Node) Insert(key int) *Node {
	var y *Node = nil
	x := rt
	for x != nil {
		y = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	nd := NewNode(key, y)

	if y == nil {
		return nd
	}

	if key < y.key {
		y.left = nd
	} else {
		y.right = nd
	}

	return rt
}

func (nd *Node) Find(key int) bool {
	if nd == nil {
		return false
	}
	if nd.key == key {
		return true
	}

	if key < nd.key {
		return nd.left.Find(key)
	} else {
		return nd.right.Find(key)
	}
}
