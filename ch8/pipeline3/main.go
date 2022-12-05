package main

import "fmt"

func counter(out chan<- int) {
	fmt.Println("In couter")
	for x := 0; x <= 10; x++ {
		fmt.Println("counter start out")
		out <- x
		fmt.Println("counter end out")
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	fmt.Println("In squarer")
	for x := range in {
		fmt.Println("Squarer in")
		out <- x * x
		fmt.Println("Squarer out")
	}
	close(out)
}

func printer(in <-chan int) {
	fmt.Println("In printer")
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
