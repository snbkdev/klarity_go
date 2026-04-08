package main

import (
	"fmt"
	"time"
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		d = 2
		c
	)

	fmt.Println(a, b, c, d)
}