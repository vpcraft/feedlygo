package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on port 8080...")
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Couldn't listen on port 8080: ", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Couldn't accept connection: ", err.Error())
			return
		}
		defer conn.Close()

		handle_connection(conn)
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted connection from ", conn.RemoteAddr())

	conn.Write([]byte("Hello!"))
}
