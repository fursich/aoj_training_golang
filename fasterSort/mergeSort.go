package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SENTINEL = 10000000000

var cnt int = 0
var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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
	defer wr.Flush()

	n := scanInt()
	s := make([]int, n)
	for i := 0; i < n; i += 1 {
		s[i] = scanInt()
	}

	// n := 10
	// s := []int{8, 5, 9, 2, 6, 3, 7, 1, 10, 4}

	mergeSort(s, 0, n)

	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
	fmt.Println(cnt)
}

func mergeSort(s []int, left int, right int) {
	if left >= right-1 {
		return
	}
	mid := (left + right) / 2
	mergeSort(s, left, mid)
	mergeSort(s, mid, right)
	merge(s, left, mid, right)
}

func merge(s []int, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid
	l := make([]int, n1+1)
	r := make([]int, n2+1)

	for i := 0; i < n1; i += 1 {
		l[i] = s[left+i]
	}
	l[n1] = SENTINEL

	for i := 0; i < n2; i += 1 {
		r[i] = s[mid+i]
	}
	r[n2] = SENTINEL

	// merge r and l back into s
	k0, k1 := 0, 0
	for i := left; i < right; i += 1 {
		cnt += 1
		if l[k0] <= r[k1] {
			s[i] = l[k0]
			k0 += 1
		} else {
			s[i] = r[k1]
			k1 += 1
		}
	}
}
