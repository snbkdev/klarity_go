// 6.2.1. Значение n il является корректным получателем

package main

import "fmt"

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

// values отображает строковый ключ на список значений
type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}

	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	//m.Add("item", "3") // error
}