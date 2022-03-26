package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	// n := 5
	a := make([]int, n)
	for i := 0; i < n; i += 1 {
		a[i] = scanInt()
		// a[i] = []int{3, 5, 2, 1, 4}[i]
	}

	countInversion(a, 0, len(a))
	fmt.Println(cnt)
}

var cnt = 0

func countInversion(a []int, left int, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	countInversion(a, left, mid)
	countInversion(a, mid, right)
	merge(a, left, mid, right)
}

const SENTINEL = 10000000000

func merge(a []int, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid
	l := make([]int, n1+1)
	r := make([]int, n2+1)

	for i := 0; i < n1; i += 1 {
		l[i] = a[left+i]
	}
	l[n1] = SENTINEL

	for i := 0; i < n2; i += 1 {
		r[i] = a[mid+i]
	}
	r[n2] = SENTINEL

	j1, j2 := 0, 0
	for i := left; i < right; i += 1 {
		if l[j1] <= r[j2] {
			a[i] = l[j1]
			j1 += 1
		} else {
			a[i] = r[j2]
			j2 += 1
			cnt += n1 - j1
		}
	}
}
