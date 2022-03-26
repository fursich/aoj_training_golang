package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fib := make([]int, n+1)
	fib[0], fib[1] = 1, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}

	fmt.Println(fib[n])
}
