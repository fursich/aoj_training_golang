package main

import "fmt"

type Set map[interface{}]bool

func (st *Set) insert(k interface{}) {
	(*st)[k] = true
}

func (st *Set) remove(k interface{}) {
	delete(*st, k)
}

func (st *Set) clear(k interface{}) {
	for k := range *st {
		delete(*st, k)
	}
}

func (st *Set) size(k interface{}) int {
	return len(*st)
}

func (st *Set) find(k interface{}) bool {
	return (*st)[k]
}

func (st *Set) merge(other *Set) {
	for k := range *other {
		(*st)[k] = true
	}
}

func main() {
	set := Set{}

	fmt.Println(set)
	set.insert(1)
	fmt.Println(set)
	set.insert(4)
	fmt.Println(set)
	set.remove(3)
	fmt.Println(set)
	set.insert(3)
	fmt.Println(set)
	set.insert(4)
	fmt.Println(set)

	other := Set{}
	other.insert(2)
	other.insert(1)
	other.insert(3)
	fmt.Println(set, other)
	set.merge(&other)
	fmt.Println(set, other)
}
