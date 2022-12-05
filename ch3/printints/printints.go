package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	s := "Testing String a A"
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(intsToString([]int{1, 2, 3, 4, 5})), intsToString([]int{1, 2, 3, 4, 5}))
}
