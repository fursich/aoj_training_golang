package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)
var stdin = os.Stdin

// var stdin2 = bytes.NewBufferString(`
//     5
//     1 2 3 4 5
//     3 2 4 1 5
//   `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanList(n int) []int {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = scanInt()
	}
	return l
}

func printList(l []int) {
	fmt.Fprintln(wr, strings.Trim(fmt.Sprint(l), "[]"))
}

const NIL = -1

var cur int
var preOrderList, inOrderList, postOrderList []int

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	preOrderList = scanList(n)
	inOrderList = scanList(n)
	postOrderList = make([]int, 0, n)

	cur = 0
	reconstruct(0, n)
	printList(postOrderList)
}

func findIndex(x int, l []int) int {
	for i, v := range l {
		if v == x {
			return i
		}
	}
	return NIL
}

func reconstruct(l int, r int) {
	if l >= r {
		return
	}
	root := preOrderList[cur]
	cur++
	m := findIndex(root, inOrderList)

	reconstruct(l, m)
	reconstruct(m+1, r)
	postOrderList = append(postOrderList, root)
}
