package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_WORD_LENGTH = 12
const DICT_CAPACITY = 2000003 // arbitary prime number that can contain <=1_000_000 numbers

type Dict [DICT_CAPACITY]string

var dict Dict
var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func getKey(w *string) int {
	var chrs []byte = []byte(*w)

	var sum int = 0
	for _, v := range chrs {
		sum *= 5
		sum += getChar(v)
	}

	return sum
}

func h1(key int) int {
	return key % DICT_CAPACITY
}

func h2(key int) int {
	return 1 + (key % (DICT_CAPACITY - 1))
}

func (d *Dict) insert(str *string) bool {

	key := getKey(str)
	hash1, hash2 := h1(key), h2(key)
	h := hash1 % DICT_CAPACITY

	for {
		if dict[h] == *str {
			return false
		} else if dict[h] == "" {
			dict[h] = *str
			return true
		}
		h = (h + hash2) % DICT_CAPACITY
	}
}

func (d *Dict) find(str *string) bool {
	key := getKey(str)
	hash1, hash2 := h1(key), h2(key)
	h := hash1 % DICT_CAPACITY

	for {
		if dict[h] == *str {
			return true
		} else if dict[h] == "" {
			return false
		}
		h = (h + hash2) % DICT_CAPACITY
	}
}

func getChar(c byte) int {
	switch c {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		return 0
	}
}

func scanInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return n
}

func scanCommand() (string, string) {
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	return s[0], s[1]
}

func main() {
	sc.Split(bufio.ScanLines)
	defer wr.Flush()

	n := scanInt()
	for i := 0; i < n; i += 1 {
		op, w := scanCommand()

		switch op {
		case "insert":
			dict.insert(&w)
		case "find":
			if dict.find(&w) {
				fmt.Fprintln(wr, "yes")
			} else {
				fmt.Fprintln(wr, "no")
			}
		}
	}
}
