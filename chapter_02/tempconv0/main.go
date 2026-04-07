package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)		// true
	fmt.Println(f >= 0) 	// true
	//fmt.Println(c == f)
	fmt.Println(c == Celsius(f)) 	// true

	d := FToC(212.0)
	fmt.Println(d.String())
	fmt.Printf("%v\n", d)
	fmt.Printf("%s\n", d)
	fmt.Println(d)
	fmt.Printf("%g\n", d)
	fmt.Println(float64(d))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}