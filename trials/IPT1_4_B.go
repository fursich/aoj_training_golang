package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Scan(&r)

	s := math.Pi * r * r
	l := 2 * math.Pi * r
	fmt.Printf("%f %f\n", s, l)
}
