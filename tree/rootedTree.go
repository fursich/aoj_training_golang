package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent  int
	child   int
	sibling int
	depth   int
}

type NodeList []Node

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)

var stdin = os.Stdin

// var stdin = bytes.NewBufferString(`
//   13
//   0 3 1 4 10
//   1 2 2 3
//   2 0
//   3 0
//   4 3 5 6 7
//   5 0
//   6 0
//   7 2 8 9
//   8 0
//   9 0
//   10 2 11 12
//   11 0
//   12 0
// `)

const MAXNODES = 100000
const NONE = -1

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (l NodeList) printList() {
	buf := make([]string, len(l))

	for i := 0; i < len(l); i++ {
		fmt.Fprintf(wr, "node %d: parent = %d, depth = %d, %s, [%s]\n", i, l[i].parent, l[i].depth, l[i].nodeType(), l.childrenOf(&l[i], buf))
	}
}

func (nl NodeList) childrenOf(nd *Node, buf []string) string {
	n := 0
	for c := nd.child; c != NONE; c = (nl)[c].sibling {
		buf[n] = fmt.Sprint(c)
		n++
	}
	return strings.Join(buf[0:n], ", ")
}

func (nd *Node) nodeType() string {
	if nd.parent == NONE {
		return "root"
	}
	if nd.child == NONE {
		return "leaf"
	}
	return "internal node"
}

func (nl NodeList) findRootIndex() int {
	for i := 0; i < len(nl); i++ {
		if (nl)[i].parent == NONE {
			return i
		}
	}
	return NONE
}

func (nl NodeList) registerDepth(idx int, depth int) {
	(nl)[idx].depth = depth
	if nl[idx].child != NONE {
		nl.registerDepth((nl)[idx].child, depth+1)
	}
	if nl[idx].sibling != NONE {
		nl.registerDepth((nl)[idx].sibling, depth)
	}
}

func newNodeList(n int) *NodeList {
	nl := make(NodeList, n)
	for i := 0; i < len(nl); i++ {
		nl[i] = Node{parent: NONE, child: NONE, sibling: NONE}
	}

	return &nl
}

func (nl NodeList) buildNodes() {
	id := scanInt()
	n := scanInt()

	l := NONE
	for i := 0; i < n; i++ {
		c := scanInt()
		if i == 0 {
			(nl)[id].child = c
		} else {
			(nl)[l].sibling = c
		}
		(nl)[c].parent = id
		l = c
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 1000000)
	defer wr.Flush()

	n := scanInt()
	list := newNodeList(n)
	for i := 0; i < n; i++ {
		list.buildNodes()
	}

	rootIndex := list.findRootIndex()
	list.registerDepth(rootIndex, 0)

	list.printList()
}
