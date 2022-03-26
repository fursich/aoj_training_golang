package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var cnt = 0

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func calculateIntervals(n int) (int, []int) {
	gs := make([]int, 0, n)

	for g := 1; ; {
		gs = append(gs, g)
		g = 3*g + 1
		if g > n {
			break
		}
	}

	for i, j := 0, len(gs)-1; i <= j; i, j = i+1, j-1 {
		gs[i], gs[j] = gs[j], gs[i]
	}

	return len(gs), gs
}

func printStatus(list []int) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func insertionSort(list []int, n int, g int) {
	for i := g; i < n; i += 1 {
		v := list[i]
		j := i - g
		for ; j >= 0 && list[j] > v; j -= g {
			list[j+g] = list[j]
			cnt += 1
		}
		list[j+g] = v
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	list := make([]int, n)
	for i := 0; i < n; i += 1 {
		list[i] = scanInt()
	}

	m, gs := calculateIntervals(n)
	for i := 0; i < m; i += 1 {
		insertionSort(list, n, gs[i])
	}
	fmt.Println(m)
	fmt.Println(strings.Trim(fmt.Sprint(gs), "[]"))
	fmt.Println(cnt)
	printStatus(list)
}
