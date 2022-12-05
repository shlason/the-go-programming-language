package main

import "fmt"

type ByteCounter int

type ByteCS struct {
	count int
}

func (c ByteCS) Write(p []byte) (int, error) {
	c.count += len(p)
	return len(p), nil
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	var cs ByteCS
	fmt.Fprintf(cs, "Hello")
	fmt.Fprintf(&c, "Hello")
	fmt.Println(ByteCS.Write(cs, []byte("dwd")))
}
