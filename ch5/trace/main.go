package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	fmt.Println("First")
	i := 1

	defer func() { fmt.Println(i) }()
	defer trace("bigSlowOperation")
	i = 2
	time.Sleep(1 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
