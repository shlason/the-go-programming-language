package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("Send start")
	go receive(ch)
	ch <- 1
	fmt.Println("main func done.")
}

func receive(ch <-chan int) {
	fmt.Println("In receive func")
	fmt.Println("Start sleep")
	time.Sleep(30 * time.Second)
	fmt.Println("Sleep end")
	<-ch
}
