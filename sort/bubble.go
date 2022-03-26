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

func printStatus(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}

func bubbleSort(arr []int) {
	cnt := 0

	for i := 0; i < len(arr); i += 1 {
		flag := false
		for j := len(arr) - 1; j > i; j -= 1 {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				cnt += 1
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	printStatus(arr)
	fmt.Println(cnt)
}

func main() {
	var arr [100]int
	sc.Split(bufio.ScanWords)

	n := scanInt()
	for i := 0; i < n; i += 1 {
		arr[i] = scanInt()
	}

	bubbleSort(arr[:n])
}
