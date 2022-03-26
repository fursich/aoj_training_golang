package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const VMAX = 10000

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
	min := VMAX

	for i := 0; i < n; i++ {
		w := scanInt()
		// w := []int{1, 5, 3, 4, 2}[i]
		a[i] = w
		if w < min {
			min = w
		}
	}

	fmt.Println(minCostSort(a, min))
}

func minCostSort(a []int, min int) int {
	n := len(a)
	b := make([]int, n)
	copy(b, a)

	sort.Ints(b)
	nextIndex := make(map[int]int, n)
	for i, v := range b {
		nextIndex[v] = i
	}

	visited := make([]bool, n)
	cost := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		cur := i
		sum := 0
		localMin := VMAX
		cnt := 0
		for !visited[cur] {
			cnt++
			visited[cur] = true
			v := a[cur]

			sum += v
			if v < localMin {
				localMin = v
			}

			cur = nextIndex[v]
		}

		ct1 := sum + (cnt-2)*localMin
		ct2 := sum + localMin + (cnt+1)*min
		if ct1 <= ct2 {
			cost += ct1
		} else {
			cost += ct2
		}
	}

	return cost
}
