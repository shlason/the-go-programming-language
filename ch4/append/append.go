package main

import (
	"fmt"
)

func main() {
	var x []int
	for i := 0; i < 20; i++ {
		x = appendInt(x, i, i+1, i+2, i+3)
		fmt.Printf("%d cap=%d\tlen=%d\t%v\n", i, cap(x), len(x), x)
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < cap(x)*2 {
			zcap = cap(x) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}
