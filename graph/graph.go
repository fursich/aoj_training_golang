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

// var stdin = bytes.NewBufferString(
// 	`4
// 1 2 2 4
// 2 1 4
// 3 0
// 4 1 3
//   `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanList() []int {
	sc.Scan()

	a := strings.Split(sc.Text(), " ")
	list := make([]int, len(a))

	for i, v := range a {
		n, _ := strconv.Atoi(v)
		list[i] = n
	}
	return list
}

func printList(l []int) {
	fmt.Fprintln(wr, strings.Trim(fmt.Sprint(l), "[]"))
}

func main() {
	sc.Split(bufio.ScanLines)
	defer wr.Flush()

	n := scanInt()
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)

		l := scanList()
		for _, j := range l[2:] {
			m[i][j-1] = 1
		}

		printList(m[i])
	}
}
