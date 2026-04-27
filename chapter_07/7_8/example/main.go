// 7.8 Интерфейс error
package main

import (
	"errors"
	"fmt"
	"syscall"
)

type error interface {
	Error() string
}

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	text string
}

func (e *errorString) Error() string {
	return e.text
}

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))

	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}

func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

type Errno uintptr

// var errors = [...]string{
// 	1: "операция не разрешена",
// 	2: "нет такого файла или каталога",
// 	3: "нет такого процесса",
// }

// func (e Errno) Error() string {
// 	if 0 <= int(e) && int(e) < len(errors) {
// 		return errors[e]
// 	}

// 	return fmt.Sprintf("errno %d", e)
// }

