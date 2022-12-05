package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	count := 0
	for count <= 1000 {
		for _, arg := range os.Args {
			s += sep + arg
			sep = " "
		}
		count++
	}
	end := time.Now()
	fmt.Println(s)
	fmt.Println(end.Sub(start))
}
