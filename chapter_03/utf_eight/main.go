package main

import (
	"fmt"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}

func main() {
	s := "Hello, дорогой мир"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// for i := 0; i < len(s); i++ {
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\n", i, r)
	// 	i += size
	// }

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)
}