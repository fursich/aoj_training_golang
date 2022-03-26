package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX_DEPTH = 20000

type Stack []interface{}

var stack = make(Stack, 0, MAX_DEPTH)
var stack2 = make(Stack, 0, MAX_DEPTH)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func (s *Stack) pop() interface{} {
	if s.isEmpty() {
		panic("Cannot pop any further")
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *Stack) push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) top() interface{} {
	return (*s)[len(*s)-1]
}

func (s *Stack) String() string {
	a := make([]string, len(*s))

	for i := 0; i < len(*s); i += 1 {
		a[i] = fmt.Sprint((*s)[i])
	}

	return strings.Trim(strings.Join(a, " "), "[]")
}

func result() string {
	a := make([]string, len(stack2))

	for i := 0; i < len(stack2); i += 1 {
		s := (stack2)[i].([]int)[1]
		a[i] = fmt.Sprint(s)
	}

	return strings.Join(a, " ")
}

func calculateDepth(angles []string) {
	var totalArea int

	for i, x := range angles {
		switch x {
		case "\\":
			stack.push(i)
		case "_":
			// do nothing
		case "/":
			if !stack.isEmpty() {
				i0 := stack.pop().(int)
				dS := i - i0
				totalArea += dS

				for !stack2.isEmpty() {
					i1 := stack2.top().([]int)[0]
					if i1 < i0 {
						break
					}

					dS += stack2.pop().([]int)[1]
				}
				stack2.push([]int{i0, dS})
			}
		default:
			panic("Could not interpret the charactor :" + x + " in " + strings.Join(angles, ""))
		}
	}

	fmt.Println(totalArea)
	if stack2.isEmpty() {
		fmt.Println("0")
	} else {
		fmt.Println(len(stack2), result())
	}
}

func main() {
	sc.Split(bufio.ScanLines)
	defer wr.Flush()

	// sc.Scan()
	// angles := strings.Split(sc.Text(), "")
	// angles := strings.Split("\\\\///\\_/\\/\\\\\\\\/_/\\\\///__\\\\\\_\\\\/_\\/_/\\", "")
	angles := strings.Split("\\\\\\___", "")

	calculateDepth(angles)
}
