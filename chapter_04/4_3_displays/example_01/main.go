package main

import "fmt"

func main() {
	// ages := map[string]int{
	// 	"alice": 31,
	// 	"charlie": 14,
	// }

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 14

	// Удаление элемента ages["alice"]
	// delete(ages, "alice")

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}