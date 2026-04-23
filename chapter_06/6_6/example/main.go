// 6.6. Инкапсуляция
package main

import (
	"fmt"
	"time"
)

type IntSet uint64

type Buffer struct {
	buf     []byte
	initial [64]byte
}

func (b *Buffer) Len() int {
	if b.buf == nil {
		return 0
	}
	return len(b.buf)
}

func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0]
	}

	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
}

type Counter struct {
	n int
}

func (c *Counter) N() int {
	return c.n
}

func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Reset() {
	c.n = 0
}

type Logger struct {
	flags  int
	prefix string
}

// func (l *Logger) Flags() int
// func (l *Logger) SetFlags(flag int)
// func (l *Logger) Prefix() string
// func (l *Logger) SetPrefix(prefix string)

func main() {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds()) // 86400
}
