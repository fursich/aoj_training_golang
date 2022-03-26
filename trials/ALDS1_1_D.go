package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int64
	var minv, maxv int64

	reader := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	fmt.Scan(&m)
	minv, maxv = m, -m

	for i := 0; i < int(n)-1; i += 1 {
		fmt.Fscan(reader, &m)

		if maxv <= m-minv {
			maxv = m - minv
		}
		if minv >= m {
			minv = m
		}
	}

	fmt.Println(maxv)
}
