package main

import (
	"fmt"
	"reflect"
)

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "Augest",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	summer := months[6:9]
	s := make([]int, 1)
	fmt.Println("len: ", len(s), "cap: ", cap(s))
	s = append(s, 3)
	s = append(s, 3)
	fmt.Println(reflect.TypeOf(summer), summer, s)
	fmt.Println("len: ", len(s), "cap: ", cap(s))
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
