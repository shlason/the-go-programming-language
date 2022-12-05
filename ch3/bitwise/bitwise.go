package main

import "fmt"

func main() {
	var x uint8 = 128<<1 - 1
	var y uint8 = 1 << 2
	var f float32 = 16777216
	var f64 float64
	newLine := '\n'

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%v\n", int(f))
	fmt.Println(f == f+1)
	fmt.Println(f64 / f64)
	fmt.Printf("%d, %[1]c, %[1]q\n", newLine)
}
