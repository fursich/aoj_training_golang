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

func searchCount(targets *[]int, searchables *[]int) int {
	cnt := 0
	for _, t := range *targets {
		n := search(t, searchables)
		if n >= 0 {
			cnt += 1
		}
	}

	return cnt
}

func search(target int, searchables *[]int) int {
	left, right := 0, len(*searchables)

	for left < right {
		mid := (left + right) / 2
		v := (*searchables)[mid]

		if target == v {
			return mid
		} else if target < v {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	searchables := make([]int, n)
	for i := 0; i < n; i += 1 {
		searchables[i] = scanInt()
	}
	q := scanInt()
	targets := make([]int, q)
	for i := 0; i < q; i += 1 {
		targets[i] = scanInt()
	}

	// n, q := 5, 3
	// searchables := []int{1, 2, 4, 8}
	// targets := []int{5}

	cnt := searchCount(&targets, &searchables)
	fmt.Println(cnt)
}
