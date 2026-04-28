// 7.11 Распознавание ошибок с помощью деклараций типов
package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	// "strings"
)

func IsExist(err error) bool

func IsPermission(err error) bool

// func IsNotExist(err error) bool {
// 	return strings.Contains(err.Error(), "файл не существует")
// }

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

var ErrNotExist = errors.New("файл не существует")

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}

	return err == syscall.ENOENT || err == ErrNotExist
}

func main() {
	_, err := os.Open("note.txt")
	fmt.Println(os.IsNotExist(err))
}