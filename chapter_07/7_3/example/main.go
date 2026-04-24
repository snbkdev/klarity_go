// 7.3 Соответствие интерфейсу
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
	w = new(bytes.Buffer)
	// w = time.Second // Ошибка: у time.Duration нет метода Write

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rws = new(bytes.Buffer) // Ошибка: у *bytes.Buffer нет метода Close

	w = rwc
	// rwc = w // Ошибка: io.Writer не имеет метода Close

	fmt.Println(w)

	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var v io.Writer
	v = os.Stdout
	v.Write([]byte("hello"))
	// v.Close() // Ошибка: io.Writer не имеет метода Close

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)

	fmt.Println(any)
}

type IntSet struct {}
// func (*IntSet) String() string
// var _ = IntSet{}.String() // Ошибка: String требует получатель *IntSet

// var s IntSet
// var _ = s.String()
// var _ fmt.Stringer = &s
// var _ fmt.Stringer = s // Ошибка: у IntSet нет метода String

