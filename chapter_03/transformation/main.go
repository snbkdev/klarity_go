package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))

	fmt.Sprintf("x=%b", x)
	x1, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	y1, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(x1)
	fmt.Println(y1)
}