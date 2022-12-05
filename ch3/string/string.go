package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	sc := "Hello, 世界"
	s1 := s[:5]
	s2 := s[7:]

	fmt.Println(&s, &s1, &s2)
	fmt.Println(len(sc))
}
