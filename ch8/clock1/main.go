package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("listen")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		fmt.Println("connected")
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Println("disconnected")
			return
		}
		time.Sleep(1 * time.Second)
	}
}
