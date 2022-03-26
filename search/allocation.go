package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAXLOAD = 10000

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return n
}

func calcMinMax(loads []int) (int, int) {
	max := 0
	min := 0
	for _, v := range loads {
		max += v
		if v > min {
			min = v
		}
	}
	return min - 1, max
}

func conveyablePackets(n int, k int, capacity int, loads []int) int {

	i := 0
	for j := 0; j < k; j += 1 {
		remaining := capacity

		for remaining >= loads[i] {
			remaining -= loads[i]
			i += 1
			if i == n {
				return n
			}
		}
	}

	return i
}

func search(n int, k int, loads []int) int {
	left, right := calcMinMax(loads)
	// fmt.Println(left, right)

	for left < right-1 {
		mid := (left + right) / 2
		// fmt.Println(left, mid, right, conveyablePackets(n, k, mid, loads), n)
		if conveyablePackets(n, k, mid, loads) < n {
			left = mid
		} else {
			right = mid
		}
	}

	return right
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	k := scanInt()
	// n, k := 5, 3

	loads := make([]int, n)
	for i := 0; i < n; i += 1 {
		loads[i] = scanInt()
		// loads[i] = []int{8, 1, 7, 3, 9}[i]
	}

	fmt.Println(search(n, k, loads))
}
