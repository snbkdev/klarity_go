package main

import "fmt"

func main() {
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 14

	var ages map[string]int
	// names := make([]string, 0, len(ages))
	fmt.Println(ages == nil) // true
	fmt.Println(len(ages) == 0) // true

	a := equal(map[string]int{"A": 0}, map[string]int{"B":12})
	fmt.Println(a) // false

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}