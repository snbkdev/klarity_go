// 5.4.1 Конец файла (EOF)
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var EOF = errors.New("EOF")

func main() {
	if err := readUntilEOF(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}

func readUntilEOF() error {
	in := bufio.NewReader(os.Stdin)
	
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			return EOF
		}
		if err != nil {
			return fmt.Errorf("сбой чтения: %v", err)
		}
		fmt.Printf("%c", r)
	}
}