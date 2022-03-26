package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)

var stdin = os.Stdin

// var stdin = bytes.NewBufferString(
// 	`
//   6
//   30 35
//   35 15
//   15 5
//   5 10
//   10 20
//   20 25
// `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())

	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	p := make([]int, n+1)

	for i := 0; i < n; i++ {
		c, r := scanInt(), scanInt()
		if i == 0 {
			p[0] = c
		}
		p[i+1] = r
	}

	fmt.Println(calcMinMcr(n, p))
}

func calcMinMcr(n int, p []int) int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		// m[i][i] = 0 // always guaranteed
	}

	for diff := 1; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff

			min := 1000000000000
			for k := i; k < j; k++ {
				v := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1]
				if v < min {
					min = v
				}
			}

			m[i][j] = min
		}
	}

	return m[0][n-1]
}
