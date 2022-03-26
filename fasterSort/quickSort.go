package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	suite  string
	number int
}

const SENTINEL = 10000000000

var sentinel = &Card{suite: "*", number: SENTINEL}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func (c1 *Card) lessThanOrEqualTo(c2 *Card) bool {
	return c1.number <= c2.number
}

func (c1 *Card) sameAs(c2 *Card) bool {
	return c1.suite == c2.suite && c1.number == c2.number
}

func (c Card) String() string {
	return fmt.Sprintf("%s %d", c.suite, c.number)
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func scanCard() *Card {
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	n, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}

	return &Card{suite: a[0], number: n}
}

func equals(c1 []*Card, c2 []*Card, n int) bool {
	for i := 0; i < n; i += 1 {
		if !c1[i].sameAs(c2[i]) {
			return false
		}
	}

	return true
}

func printCards(c []*Card) {
	for i := 0; i < len(c); i += 1 {
		fmt.Fprintln(wr, *c[i])
	}
}

func main() {
	sc.Split(bufio.ScanLines)
	defer wr.Flush()

	n := scanInt()
	// n := 6

	cards1 := make([]*Card, n)
	cards2 := make([]*Card, n)
	for i := 0; i < n; i += 1 {
		card := scanCard()
		// card := []*Card{{"D", 3}, {"H", 2}, {"D", 1}, {"S", 3}, {"D", 2}, {"C", 1}}[i]
		cards1[i] = card
		cards2[i] = card
	}

	quickSort(cards1, 0, n-1)
	mergeSort(cards2, 0, n-1)

	if equals(cards1, cards2, n) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	printCards(cards1)
	// printCards(cards2)
}

func mergeSort(cards []*Card, left int, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(cards, left, mid)
	mergeSort(cards, mid, right)
	merge(cards, left, mid, right)
}

func merge(cards []*Card, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid
	l := make([]*Card, n1+1)
	r := make([]*Card, n2+1)

	for i := 0; i < n1; i += 1 {
		l[i] = cards[left+i]
	}
	l[n1] = sentinel
	for i := 0; i < n2; i += 1 {
		r[i] = cards[mid+i]
	}
	r[n2] = sentinel

	j1, j2 := 0, 0
	for i := left; i < right; i += 1 {
		if l[j1].lessThanOrEqualTo(r[j2]) {
			cards[i] = l[j1]
			j1 += 1
		} else {
			cards[i] = r[j2]
			j2 += 1
		}
	}
}

func quickSort(cards []*Card, p int, r int) {
	if p >= r {
		return
	}

	q := partition(cards, p, r)
	quickSort(cards, p, q-1)
	quickSort(cards, q+1, r)
}

func partition(cards []*Card, p int, r int) int {
	x := cards[r]
	i := p - 1
	for j := p; j < r; j += 1 {
		if cards[j].lessThanOrEqualTo(x) {
			i += 1
			cards[i], cards[j] = cards[j], cards[i]
		}
	}
	cards[i+1], cards[r] = cards[r], cards[i+1]

	return i + 1
}
