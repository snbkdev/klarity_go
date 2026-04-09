package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) //[5 4 3 2 1 0]

	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// Циклический сдвиг влево на две позиции
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // [2 3 4 5 6 7 0 1]
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}