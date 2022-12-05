package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

type point struct {
	x, y int
}

func main() {
	s1 := []int{2, 5, 6, 34, 6, 78, 3, 65, 7, 4}
	pp := new(point)

	Sort(s1)
	fmt.Println(s1, pp)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
