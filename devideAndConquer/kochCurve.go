package main

import (
	"fmt"
	"math"
)

type Vector struct {
	x float64
	y float64
}

type Line struct {
	base   *Vector
	vector *Vector
}

var frac3 float64 = 1.0 / 3.0
var cos60 float64 = 1.0 / 2.0
var sin60 float64 = math.Sqrt(3) / 2

func (v *Vector) add(v2 *Vector) *Vector {
	return &Vector{x: v.x + v2.x, y: v.y + v2.y}
}

func (v *Vector) sub(v2 *Vector) *Vector {
	return &Vector{x: v.x - v2.x, y: v.y - v2.y}
}

func (v *Vector) mul(t float64) *Vector {
	return &Vector{x: t * v.x, y: t * v.y}
}

func (v *Vector) rotate60(sign bool) *Vector {
	if sign {
		return &Vector{x: cos60*v.x - sin60*v.y, y: sin60*v.x + cos60*v.y}
	} else {
		return &Vector{x: cos60*v.x + sin60*v.y, y: -sin60*v.x + cos60*v.y}
	}
}

func (l *Line) endPoint() *Vector {
	return &Vector{l.base.x + l.vector.x, l.base.y + l.vector.y}
}

func newLine(base *Vector, vector *Vector) *Line {
	return &Line{base: base, vector: vector}
}

func newLine2(endPoint *Vector, vector *Vector) *Line {
	base := endPoint.sub(vector)
	return &Line{base: base, vector: vector}
}

func main() {
	var n int
	fmt.Scan(&n)

	line := Line{base: &Vector{x: 0, y: 0}, vector: &Vector{x: 100, y: 0}}
	printEdge(line.base)

	generateKoch(n, &line)
}

func generateKoch(n int, line *Line) {
	if n == 0 {
		printEdge(line.endPoint())
		return
	}

	b1 := line.base
	v1 := line.vector.mul(frac3)
	l1 := newLine(b1, v1)
	generateKoch(n-1, l1)

	b2 := l1.endPoint()
	v2 := v1.rotate60(true)
	l2 := newLine(b2, v2)
	generateKoch(n-1, l2)

	ve3 := line.vector.mul(2 * frac3)
	e3 := line.base.add(ve3)
	v3 := v1.rotate60(false)
	l3 := newLine2(e3, v3)
	generateKoch(n-1, l3)

	b4 := e3
	v4 := v1
	l4 := newLine(b4, v4)
	generateKoch(n-1, l4)
}

func printEdge(v *Vector) {
	fmt.Printf("%.8f %.8f\n", v.x, v.y)
}
