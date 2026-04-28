// 7.14 Пример: XML-декодирование на основе лексем
package main

import "io"

type Name struct {
	Local string
}

type Attr struct {
	Name Name
	Value string
}

type Token interface{}

type StartElement struct {
	Name Name
	Attr []Attr
}

type EndElement struct { Name Name }
type CharData []byte
type Comment []byte

type Decoder struct {}
func NewDecoder(io.Reader) *Decoder
func (*Decoder) Token() (Token, error)