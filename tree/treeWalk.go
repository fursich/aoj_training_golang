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

	fmt.Fprintln(wr, "Preorder")
	list.WalkWithPreOrder(rid)
	fmt.Fprintln(wr)

	fmt.Fprintln(wr, "Inorder")
	list.WalkWithInOrder(rid)
	fmt.Fprintln(wr)

	fmt.Fprintln(wr, "Postorder")
	list.WalkWithPostOrder(rid)
	fmt.Fprintln(wr)
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

func (nl NodeList) WalkWithPreOrder(id int) {
	fmt.Fprintf(wr, " %d", id)
	if nl[id].left != NIL {
		nl.WalkWithPreOrder(nl[id].left)
	}
	if nl[id].right != NIL {
		nl.WalkWithPreOrder(nl[id].right)
	}
}

func (nl NodeList) WalkWithInOrder(id int) {
	if nl[id].left != NIL {
		nl.WalkWithInOrder(nl[id].left)
	}
	fmt.Fprintf(wr, " %d", id)
	if nl[id].right != NIL {
		nl.WalkWithInOrder(nl[id].right)
	}
}

func (nl NodeList) WalkWithPostOrder(id int) {
	if nl[id].left != NIL {
		nl.WalkWithPostOrder(nl[id].left)
	}
	if nl[id].right != NIL {
		nl.WalkWithPostOrder(nl[id].right)
	}
	fmt.Fprintf(wr, " %d", id)
}

func (nl NodeList) FindRoot() int {
	for i := 0; i < len(nl); i++ {
		if nl[i].parent == NIL {
			return i
		}
	}
	return NIL
}
