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
	perim := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance()) // 12
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