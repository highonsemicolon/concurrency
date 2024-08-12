package main

import (
	"fmt"
	"log"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	conn.Read(request)

	conn.Write([]byte("HTTP/1.1 200 ok \r\n\r\nhello there!\r\n"))
	fmt.Println("done processing")

}

func main() {
	addr := "localhost:3000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("listening on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go process(conn)
	}
}
