// 7.1 Интерфейсы как контракты
package main

import (
	// "bytes"
	"fmt"
	// "io"
	// "os"
)

// func Fprintf(w io.Writer, format string, args ...interface{})(int, error)

// func Printf(format string, args ...interface{}) (int, error) {
// 	return Fprintf(os.Stdout, format, args...)
// }

// func Sprintf(format string, args ...interface{}) string {
// 	var buf bytes.Buffer
// 	Fprintf(&buf, format, args...)
// 	return buf.String()
// }

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	* c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0

	var name = "John"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}

type Stringer interface {
	String() string
}