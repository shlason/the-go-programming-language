package main

import "fmt"

const (
	KiB uint64 = 1 << (10 * (iota + 1))
	MiB
	GiB
	TiB
	PiB
)

func main() {
	fmt.Println(KiB, MiB, GiB, TiB, PiB)
}
