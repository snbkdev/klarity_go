// 6.4. Значения-методы и выражения-методы

package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q)) // 5
	fmt.Printf("%T\n", distance) // func(main.Point, main.Point) float64

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p) // {2 4}
	fmt.Printf("%T\n", scale) // func(*main.Point, float64)

	// distanceFromP := p.Distance
	// fmt.Println(distanceFromP(q)) // 5
	// var origin Point
	// fmt.Println(distanceFromP(origin)) // 2.23606797749979

	// scaleP := p.ScaleBy
	// scaleP(2)
	// scaleP(3)
	// scaleP(10)
}