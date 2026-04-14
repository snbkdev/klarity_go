// 5.7 Вариативные функции

package main

import (
	"fmt"
	"os"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func main() {
	fmt.Println(sum()) // 0
	fmt.Println(sum(3)) // 3
	fmt.Println(sum(1, 2, 3, 4)) // 10

	values := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(values...)) // 21

	linenum, name := 12, "count"
	errorf(linenum, "not identified %s", name)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, " page %d ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}