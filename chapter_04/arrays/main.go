package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q)

	symbol := [...]string{USD: "$", EUR:"E", GBP:"£", RUB:"P"}
	fmt.Println(RUB, symbol[RUB])
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUB
)