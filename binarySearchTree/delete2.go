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
// 18
// insert 8
// insert 2
// insert 3
// insert 7
// insert 22
// insert 1
// find 1
// find 2
// find 3
// find 4
// find 5
// find 6
// find 7
// find 8
// print
// delete 3
// delete 7
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
	case "find":
		k := scanInt()
		if root.IsPresent(k) {
			fmt.Fprintln(wr, "yes")
		} else {
			fmt.Fprintln(wr, "no")
		}
	case "delete":
		k := scanInt()
		root = root.Delete(k)
	case "print":
		root.PrintInOrdered()
		fmt.Fprintln(wr)
		root.PrintPreOrdered()
		fmt.Fprintln(wr)
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

const NIL = -1

func (nd *Node) PrintNodes() {
	var p, l, r int
	if nd.parent == nil {
		p = NIL
	} else {
		p = nd.parent.key
	}
	if nd.left == nil {
		l = NIL
	} else {
		l = nd.left.key
	}
	if nd.right == nil {
		r = NIL
	} else {
		r = nd.right.key
	}
	fmt.Fprintf(wr, " %d <%d, %d, %d>", nd.key, p, l, r)

	if nd.left != nil {
		nd.left.PrintNodes()
	}

	if nd.right != nil {
		nd.right.PrintNodes()
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

func (nd *Node) IsPresent(key int) bool {
	return nd.Find(key) != nil
}

func (nd *Node) Find(key int) *Node {
	if nd == nil || nd.key == key {
		return nd
	}

	if key < nd.key {
		return nd.left.Find(key)
	} else {
		return nd.right.Find(key)
	}
}

func (rt *Node) Delete(key int) *Node {
	nd := rt.Find(key)
	if nd == nil {
		panic("Cannot delete. The key does not exist")
	}

	newRoot := nd.DeleteOne()
	if nd == rt {
		return newRoot
	} else {
		return rt
	}
}

func (nd *Node) DeleteOne() *Node {
	if nd.left == nil && nd.right == nil {

		if nd.parent == nil {
			return nil
		}
		if nd.parent.left == nd {
			nd.parent.left = nil
		} else {
			nd.parent.right = nil
		}
		return nd.parent
	}

	if nd.left == nil {
		nd.right.parent = nd.parent

		if nd.parent == nil {
			return nd.right
		}
		if nd.parent.left == nd {
			nd.parent.left = nd.right
		} else {
			nd.parent.right = nd.right
		}
		return nd.parent
	}

	if nd.right == nil {
		nd.left.parent = nd.parent

		if nd.parent == nil {
			return nd.left
		}
		if nd.parent.left == nd {
			nd.parent.left = nd.left
		} else {
			nd.parent.right = nd.left
		}
		return nd.parent
	}

	nd2 := nd.FindNextWithInOrder()
	nd.key = nd2.key
	nd2.DeleteOne()
	return nd
}

func (nd *Node) FindNextWithInOrder() *Node {
	var l *Node = nil

	for c := nd.right; c != nil; c = c.left {
		l = c
	}

	return l
}
