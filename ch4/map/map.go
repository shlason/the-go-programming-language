package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{
		"Jason": 24,
		"Lance": 25,
	}
	m2 := map[string]int{
		"Jason": 25,
		"Lance": 25,
	}
	fmt.Println(equal(m1, m2))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
