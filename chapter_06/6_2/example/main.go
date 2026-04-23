package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main() {
	r := &Point{1 ,2}
	r.ScaleBy(2)
	fmt.Println(*r) // {2 4}

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // {2 4}

	d := Point{1, 2}
	(&d).ScaleBy(2)
	fmt.Println(d) // {2 4}
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := 1; i < len(path); i++ {
		sum += path[i-1].DistanceTo(path[i])
	}
	return sum
}

func (p Point) DistanceTo(q Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}