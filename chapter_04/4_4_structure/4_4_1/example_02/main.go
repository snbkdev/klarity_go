package main

import (
	"fmt"
	"time"
)

// import "program_lang/chapter_04/4_4_structure/4_4_1/example_02/p"

// var _ = p.T{a: 1, b:2}
// var _ = p.T{1, 2}

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func main() {
	fmt.Println(Scale(Point{1, 2}, 5))

	pp := Point{1, 2}
	fmt.Println(pp)
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}