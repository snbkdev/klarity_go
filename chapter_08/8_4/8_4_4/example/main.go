// 8.4.4 Буферизованные каналы
package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "A"
	ch <- "B"
	ch <- "C"

	fmt.Println(<-ch) // A
	fmt.Println(cap(ch)) //3 
	fmt.Println(len(ch)) // 2

	fmt.Println(<-ch) // B
	fmt.Println(<-ch) // C
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()

	return <- responses
}

func request(hostname string) (response string) {
	return
}