package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type Node struct {
	prev *Node
	next *Node
	key  int
}

var list = Node{prev: nil, next: nil, key: -1}

func (list *Node) head() *Node {
	return list.next
}

func (list *Node) last() *Node {
	return list.prev
}

func (list *Node) Length() int {
	length := 0
	for node := list.head(); node != nil && node.key != -1; node = node.next {
		length += 1
	}
	return length
}

func (list Node) String() string {
	var keys = make([]string, list.Length())

	idx := 0
	for node := list.head(); node != nil && node.key != -1; node = node.next {
		keys[idx] = fmt.Sprint(node.key)
		idx += 1
	}

	return strings.Join(keys, " ")
}

func (list *Node) Find(key int) *Node {
	for node := list.head(); node != nil && node.key != -1; node = node.next {
		if node.key == key {
			return node
		}
	}
	return nil
}

func newNode(prev *Node, next *Node, key int) *Node {
	return &Node{prev: prev, next: next, key: key}
}

func (node *Node) Insert(key int) {
	next := node.next
	node.next = newNode(node, next, key)

	if next != nil {
		next.prev = node.next
	}
}

func (node *Node) deleteNode() {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}

func (list *Node) Delete(key int) {
	node := list.Find(key)
	if node != nil {
		node.deleteNode()
	}
}

func (list *Node) DeleteFirst() {
	head := list.head()
	if head != nil {
		head.deleteNode()
	}
}

func (list *Node) DeleteLast() {
	last := list.last()
	if last != nil {
		last.deleteNode()
	}
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func parseCommand() (string, int) {
	sc.Scan()
	commands := strings.Split(sc.Text(), " ")

	switch len(commands) {
	case 1:
		return commands[0], 0
	case 2:
		key, _ := strconv.Atoi(commands[1])
		return commands[0], key
	default:
		panic("Cannot parse the commands: " + fmt.Sprint(commands))
	}
}

func main() {
	sc.Split(bufio.ScanLines)
	defer wr.Flush()
	list.next = &list
	list.prev = &list

	n := scanInt()
	for i := 0; i < n; i += 1 {
		command, key := parseCommand()

		switch command {
		case "insert":
			list.Insert(key)
		case "delete":
			list.Delete(key)
		case "deleteFirst":
			list.DeleteFirst()
		case "deleteLast":
			list.DeleteLast()
		default:
			panic("Invalid command: " + command)
		}
	}
	fmt.Println(list)
}
