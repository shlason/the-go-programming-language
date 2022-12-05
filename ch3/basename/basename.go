package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := basename2("1/2/3/4g/dg/wehwe/.f.fwev.ew.ef.dd")
	test2 := basename2(".f.fwev.ew.ef.dd")
	test3 := basename2("f.fwev.ew.ef.dd")

	fmt.Println(test1, test2, test3)
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slashIndex := strings.LastIndex(s, "/")
	s = s[slashIndex+1:]
	if dotIndex := strings.LastIndex(s, "."); dotIndex >= 0 {
		s = s[:dotIndex]
	}
	return s
}
