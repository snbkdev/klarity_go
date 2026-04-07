package main

import (
	"fmt"
	"program_lang/chapter_02/tempconv_pack/pack"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", pack.AbsoluteZeroC)
	fmt.Println(pack.CToF(pack.BoilingC))
}