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

func printResult(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}

func insertionSort(arr []int) {
	for i, v := range arr {
		j := i - 1
		for ; j >= 0 && arr[j] > v; j -= 1 {
			arr[j+1] = arr[j]
		}

		arr[j+1] = v
		printResult(arr)
	}
}

func main() {
	var arr [100]int

	sc.Split(bufio.ScanWords)

	n := scanInt()
	for i := 0; i < n; i += 1 {
		arr[i] = scanInt()
	}

	insertionSort(arr[:n])
}
