package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Card struct {
	suite  string
	number int
}

func printList(list []Card) {
	fmt.Println(strings.Trim(fmt.Sprint(list), "[]"))
}

func (c Card) String() string {
	return fmt.Sprintf("%s%d", c.suite, c.number)
}

func (c Card) IsLessThan(other Card) bool {
	return c.number < other.number
}

func (c Card) Equals(other Card) bool {
	return c.number == other.number
}

func (c Card) Identical(other Card) bool {
	return c.suite == other.suite && c.number == other.number
}

func (c Card) Clone() Card {
	return Card{suite: c.suite, number: c.number}
}

func printStability(src []Card, dst []Card) {
	for i, s1 := range src {
		for _, s2 := range src[i+1:] {
			if !s1.Equals(s2) {
				continue
			}

			for j, d1 := range dst {
				if !d1.Identical(s2) {
					continue
				}
				for _, d2 := range dst[j+1:] {
					if !d2.Identical(s1) {
						continue
					}
					fmt.Println("Not stable")
					return
				}
			}
		}
	}

	fmt.Println("Stable")
	return
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func scanCard() Card {
	sc.Scan()
	s := strings.Split(sc.Text(), "")
	i, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}

	return Card{suite: s[0], number: i}
}

func bubbleSort(list []Card, orig []Card) {
	for i := 0; i < len(list); i += 1 {
		for j := len(list) - 1; j > i; j -= 1 {
			if list[j].IsLessThan(list[j-1]) {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	printList(list)
	printStability(orig, list)
}

func selectionSort(list []Card, orig []Card) {
	for i := 0; i < len(list); i += 1 {
		minIdx := i
		for j := i; j < len(list); j += 1 {
			if list[j].IsLessThan(list[minIdx]) {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	printList(list)
	printStability(orig, list)
}

func main() {
	var orig, arr1, arr2 [36]Card

	sc.Split(bufio.ScanWords)

	n := scanInt()
	for i := 0; i < n; i += 1 {
		orig[i] = scanCard()
		arr1[i] = orig[i].Clone()
		arr2[i] = orig[i].Clone()
	}

	bubbleSort(arr1[:n], orig[:n])
	selectionSort(arr2[:n], orig[:n])
}
