package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)
	rx.Set(reflect.ValueOf(3))
	rx.SetString("hello")

	rx.Set(reflect.ValueOf("hello"))

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	ry.SetInt(2)

	ry.Set(reflect.ValueOf(3))
	ry.SetString("hello")

	ry.Set(reflect.ValueOf("hello"))

	stdout := reflect.ValueOf(os.Stdout).Elem()

	fmt.Println(stdout.Type())
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.Int())
	fd.SetInt(2)

	fmt.Println(fd.CanAddr(), fd.CanSet())
}