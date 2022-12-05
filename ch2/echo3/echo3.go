package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	count := 0
	start := time.Now()
	for count <= 1000 {
		fmt.Println(strings.Join(os.Args, " "))
		count++
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
