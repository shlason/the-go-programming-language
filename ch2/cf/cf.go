package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}
		c := Celsius(t)
		f := Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
	}
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g˚F", f)
}
