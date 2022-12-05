package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

type Currency int

const (
	USD Currency = iota + 10
	EUR
	GBP
	RMB
)

func main() {
	var a = [3]string{"12"}
	var b = [32]byte{4, 3, 14, 5}
	c1 := sha256.Sum256([]byte{})
	s1 := "testing"
	p := &s1
	p2 := p
	symbol := [...]string{EUR: "e", USD: "u", GBP: "g", RMB: "r"}
	a[1] = "2"
	fmt.Println(a, symbol[0], symbol, USD, c1, b)
	zero(&b)
	*p = "ch"
	*p2 = "ch22"
	fmt.Println(b, &b, s1, *p, *p2)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
	fmt.Println(reflect.TypeOf(ptr))
}
