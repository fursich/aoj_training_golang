package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const NIL = -1

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)

var stdin = os.Stdin

// var stdin = bytes.NewBufferString(`
// 9
// 0 1 4
// 1 2 3
// 2 -1 -1
// 3 -1 -1
// 4 5 8
// 5 6 7
// 6 -1 -1
// 7 -1 -1
// 8 -1 -1
// `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	list := buildNodeList(n)

	rid := list.FindRoot()
	list.SetDepth(rid, 0)
	list.SetHeight(rid)
	list.Print()
}

type Node struct {
	parent int
	left   int
	right  int
	degree int
	depth  int
	height int
}

type NodeList []Node

func buildNodeList(n int) NodeList {
	nl := make(NodeList, n)
	for i := 0; i < n; i++ {
		nl[i].parent = NIL
		nl[i].left = NIL
		nl[i].right = NIL
	}

	for i := 0; i < n; i++ {
		p := scanInt()
		l := scanInt()
		r := scanInt()
		nl[p].left = l
		nl[p].right = r

		deg := 0
		if l != NIL {
			nl[l].parent = p
			deg++
		}
		if r != NIL {
			nl[r].parent = p
			deg++
		}
		nl[p].degree = deg
	}

	return nl
}

func (nl NodeList) FindRoot() int {
	for i := 0; i < len(nl); i++ {
		if nl[i].parent == NIL {
			return i
		}
	}
	return NIL
}

func (nl NodeList) SetDepth(id int, depth int) {
	nl[id].depth = depth

	if nl[id].left != NIL {
		nl.SetDepth(nl[id].left, depth+1)
	}
	if nl[id].right != NIL {
		nl.SetDepth(nl[id].right, depth+1)
	}
}

func (nl NodeList) SetHeight(id int) int {
	h1, h2 := 0, 0

	if nl[id].left != NIL {
		h1 = nl.SetHeight(nl[id].left) + 1
	}
	if nl[id].right != NIL {
		h2 = nl.SetHeight(nl[id].right) + 1
	}

	if h1 >= h2 {
		nl[id].height = h1
		return h1
	} else {
		nl[id].height = h2
		return h2
	}
}

func (nl NodeList) SiblingFor(i int) int {
	nd := nl[i]
	if nd.parent == NIL {
		return NIL
	}
	if nl[nd.parent].left == i {
		return nl[nd.parent].right
	} else {
		return nl[nd.parent].left
	}
}

func (nd *Node) NodeType() string {
	if nd.parent == NIL {
		return "root"
	}
	if nd.degree == 0 {
		return "leaf"
	}
	return "internal node"
}

func (nl NodeList) Print() {
	for i := 0; i < len(nl); i++ {
		fmt.Fprintf(wr, "node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n", i, nl[i].parent, nl.SiblingFor(i), nl[i].degree, nl[i].depth, nl[i].height, nl[i].NodeType())
	}
}
