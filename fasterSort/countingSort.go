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
	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	// n := 7
	a := make([]int, n)

	max := 0
	for i := 0; i < n; i += 1 {
		m := scanInt()
		// m := []int{2, 5, 1, 3, 2, 3, 0}[i]

		a[i] = m
		if m > max {
			max = m
		}
	}

	b := countSort(a, max)
	printSlice(b)
}

func countSort(a []int, max int) []int {
	counter := make([]int, max+1)
	result := make([]int, len(a))

	for _, v := range a {
		counter[v] += 1
	}

	for i := 1; i <= max; i += 1 {
		counter[i] += counter[i-1]
	}

	for i := len(a) - 1; i >= 0; i -= 1 {
		num := a[i]
		result[counter[num]-1] = num
		counter[num] -= 1
	}

	return result
}
