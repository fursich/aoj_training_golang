package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

var stack = make(Stack, 0, 300)
var sc = bufio.NewScanner(os.Stdin)

func (st *Stack) push(n int) {
	*st = append(*st, n)
}

func (st *Stack) pop() int {
	n := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return n
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func run(inputs []string) {
	for _, v := range inputs {
		processOne(v)
	}
}

func processOne(s string) {
	if isNumber(s) {
		stack.push(toInt(s))
	} else {
		operateOne(s)
	}
}

func operateOne(s string) {
	a, b := stack.pop(), stack.pop()

	switch s {
	case "+":
		stack.push(b + a)
	case "-":
		stack.push(b - a)
	case "*":
		stack.push(b * a)
	}
}

func main() {
	sc.Split(bufio.ScanLines)

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	run(inputs)

	fmt.Println(stack.pop())
}
