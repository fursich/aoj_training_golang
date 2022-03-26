package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func printSlice(s []int) {
	fmt.Print(strings.Trim(fmt.Sprint(s), "[]"))
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	a := make([]int, n)
	for i := 0; i < n; i += 1 {
		a[i] = scanInt()
	}
	// n := 12
	// a := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}

	q := partition(a, 0, n-1)
	if q > 0 {
		printSlice(a[:q])
	}
	fmt.Printf(" [%d] ", a[q])
	if q < n-1 {
		printSlice(a[q+1:])
	}
	fmt.Println()
}

func partition(a []int, n0 int, n1 int) int {
	x := a[n1]
	i := n0 - 1

	for j := n0; j < n1; j += 1 {
		if a[j] <= x {
			i = i + 1
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[n1] = a[n1], a[i+1]
	return i + 1
}
