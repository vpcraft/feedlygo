package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening on port 8080...")
	listener, err := net.Listen("tcp", "localhost:8080")

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

		go handle_connection(conn)
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted connection from ", conn.RemoteAddr())

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Connection closed")
				return
			}
			fmt.Println("Couldn't read from connection: ", err.Error())
			return
		}
		fmt.Println("Received: ", string(buf[:n]))

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Couldn't write to connection: ", err.Error())
			return
		}
	}
}
