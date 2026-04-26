// 7.4 Анализ флагов с помощью flag.value
package main

import (
	"flag"
	"fmt"
	"time"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

type Value interface {
	String() string
	Set(string) error
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	
	_, err := fmt.Sscanf(s, "%f%s", &value, &unit)
	if err != nil {
		return fmt.Errorf("invalid temperature format %q", s)
	}
	
	switch unit {
	case "C", "°C", "Celsius":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F", "Fahrenheit":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	default:
		return fmt.Errorf("wrong temperature unit %q (use C or F)", unit)
	}
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	var temp = CelsiusFlag("temp", 20.0, "temperature (e.g., '25C' or '77F')")
	
	flag.Parse()
	
	fmt.Println("Temperature:", *temp)
	fmt.Println("Period:", *period)
}