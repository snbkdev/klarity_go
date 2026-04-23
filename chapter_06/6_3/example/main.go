// 6.3 Создание типов путем встраивания структур

package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct {
	X, Y float64
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
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // 5
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // {2 2} {2 2}

	// var cp ColoredPoint
	// cp.X = 1
	// fmt.Println(cp.Point.X) // 1

	// cp.Point.Y = 2
	// fmt.Println(cp.Y) // 2

	// red := color.RGBA{255, 0, 0, 255}
	// blue := color.RGBA{0, 0, 255, 255}
	// var p1 = ColoredPoint{Point{1, 1}, red}
	// var q1 = ColoredPoint{Point{5, 4}, blue}

	// fmt.Println(p1.Distance(q1.Point)) // 5

	// p1.ScaleBy(2)
	// q1.ScaleBy(2)
	// fmt.Println(p1.Distance(q1.Point)) //
}

var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}