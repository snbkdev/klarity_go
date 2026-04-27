// 7.10 Декларации типов
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c := w.(*bytes.Buffer)

	fmt.Println(f, c)
}