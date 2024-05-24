package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	defer conn.Close()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	path := strings.Split(string(buf), " ")[1]

	if path != "/" {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return

	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
