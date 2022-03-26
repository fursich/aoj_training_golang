package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func printList(list []int) {
	fmt.Println(strings.Trim(fmt.Sprint(list), "[]"))
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func selectionSort(list []int) {
	cnt := 0

	for i := 0; i < len(list); i += 1 {
		min := i
		for j := i; j < len(list); j += 1 {
			if list[j] < list[min] {
				min = j
			}
		}
		if i != min {
			cnt += 1
			list[i], list[min] = list[min], list[i]
		}
	}
	printList(list)
	fmt.Println(cnt)
}

func main() {
	var arr [100]int

	sc.Split(bufio.ScanWords)

	n := scanInt()
	for i := 0; i < n; i += 1 {
		arr[i] = scanInt()
	}

	selectionSort(arr[:n])
}
