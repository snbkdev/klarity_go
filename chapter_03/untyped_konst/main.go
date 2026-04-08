package main

import (
	"fmt"
	"math"
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	const Pi64 float64 = math.Pi

	var x1 float32 = float32(Pi64)
	var y1 float64 = Pi64
	var z1 complex128 = complex128(Pi64)
	fmt.Println(x1, y1, z1)

	var f float64 = 212
	fmt.Println((f-32)*5/9)
	fmt.Println(5/9*(f-32))
	fmt.Println(5.0/9.0*(f-32))

	// var f1 float64 = 3 + 0i
	// f1 = 2
	// f1 = 1e123
	// f1 = 'a'

	// var f1 float64 = float64(3 + 0i)
	// f1 = float64(2)
	// f1 = float64(1e123)
	// f1 = float64('a')

	const (
		deadbeef = 0xdeadbeef
		a = uint32(deadbeef)
		b = float32(deadbeef)
		// c = float64(deadbeef)
		// d = int32(deadbeef)
		// e = float64(1e309)
		// f = uint(-1)
	)

	i := 0
	r := '\000'
	g := 0.0
	h := 0i
	fmt.Println(i, r, g, h)

	var i1 = int8(0)
	var i2 int8 = 0
	fmt.Println(i1, i2)

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}