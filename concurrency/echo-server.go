package main

import (
	"io"
	"log"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		log.Println("connection!")
	}
}

func handleConn(c net.Conn) {
	io.Copy(c, c)
	c.Close()
}
