package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	c := "left leg"
	t := c
	c += ", right leg"
	fmt.Println(c)
	fmt.Println(t)
}