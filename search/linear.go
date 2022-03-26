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

func scanIntList(sz int) []int {
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	a := make([]int, sz)

	for i := 0; i < len(s); i += 1 {
		a[i], _ = strconv.Atoi(s[i])
	}
	return a
}

func countSearch(targets *[]int, searchables *[]int) int {
	cnt := 0

	for _, v := range *targets {
		if search(v, searchables) >= 0 {
			cnt += 1
		}
	}

	return cnt
}

func search(target int, searchables *[]int) int {
	(*searchables)[len(*searchables)-1] = target

	i := 0
	for ; (*searchables)[i] != target; i += 1 {
	}

	if i == len(*searchables)-1 {
		return -1
	}
	return i
}

func main() {
	sc.Split(bufio.ScanLines)

	n := scanInt()
	searchables := scanIntList(n + 1) // +1 room for sentinel
	q := scanInt()
	targets := scanIntList(q)

	// n, q := 5, 3
	// searchables := []int{1, 2, 3, 4, 5}
	// targets := []int{3, 4, 1}

	cnt := countSearch(&targets, &searchables)
	fmt.Println(cnt)
}
