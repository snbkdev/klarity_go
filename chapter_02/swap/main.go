package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}

	return x
}

func main() {
	a := gcd(81, 27)
	fmt.Println(a)
}