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
	numbers := make([]int, n)
	for i := 0; i < n; i += 1 {
		numbers[i] = scanInt()
	}

	q := scanInt()
	targets := make([]int, q)
	for i := 0; i < q; i += 1 {
		targets[i] = scanInt()
	}

	// numbers := []int{1, 5, 7, 10, 21}
	// targets := []int{2, 4, 17, 8}
	for _, v := range targets {
		if isComposable(v, numbers, 0) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func isComposable(v int, targets []int, k int) bool {
	if v == 0 {
		return true
	}
	if k >= len(targets) || v < 0 {
		return false
	}

	return isComposable(v, targets, k+1) || isComposable(v-targets[k], targets, k+1)
}
