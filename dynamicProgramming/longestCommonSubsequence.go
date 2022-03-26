package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(stdin)
var wr = bufio.NewWriter(os.Stdout)

var stdin = os.Stdin

// var stdin = bytes.NewBufferString(`
// 3
// abcbdab
// bdcaba
// abc
// abc
// abc
// bc
// `)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())

	return i
}

func scanChars() []byte {
	sc.Scan()
	return []byte(sc.Text())
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := scanInt()
	for i := 0; i < n; i++ {
		a := scanChars()
		b := scanChars()
		fmt.Fprintln(wr, calcLongestCommonSequence(a, b))
	}
}

func calcLongestCommonSequence(a []byte, b []byte) int {
	n, m := len(a), len(b)
	var lcs [][]int = make([][]int, n+1)

	for i := 0; i <= n; i++ {
		lcs[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				if lcs[i-1][j] > lcs[i][j-1] {
					lcs[i][j] = lcs[i-1][j]
				} else {
					lcs[i][j] = lcs[i][j-1]
				}
			}
		}
	}

	return lcs[n][m]
}
