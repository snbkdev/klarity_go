package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return z
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func main() {
	fmt.Printf("%T\n", add(3, 4))
	fmt.Printf("%T\n", sub(4, 3))
	fmt.Printf("%T\n", first(5, 6))
	fmt.Printf("%T\n", zero(1, 3))
}