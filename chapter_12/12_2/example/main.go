// 12.2 reflect.Type reflect.Value
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	a := v.Type()
	fmt.Println(a.String())

	b := reflect.ValueOf(3)
	c := b.Interface()
	d := c.(int)
	fmt.Printf("%d\n", d)
}