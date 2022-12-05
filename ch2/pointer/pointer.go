package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := " "
	p := new(int)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(p)
}
