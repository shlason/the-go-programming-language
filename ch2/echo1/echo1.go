package echo1

import (
	"fmt"
	"os"
	"time"
)

func echo1() {
	var s, sep string = "", ""
	start := time.Now()
	count := 0
	for count <= 1000 {
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		count++
	}
	end := time.Now()
	fmt.Println(s)
	fmt.Println(end.Sub(start))
}
