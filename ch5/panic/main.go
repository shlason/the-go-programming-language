package main

import "fmt"

func main() {
	recp()
	fmt.Println("Complete recp")
}

func recp() {
	defer func() {
		p := recover()
		fmt.Print(p)
	}()
	for i := 3; i >= 1; i-- {
		fmt.Println(1 / i)
	}
}
