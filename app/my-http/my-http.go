package myhttp

import (
	"fmt"
	"net"
	"os"
)

var routes = make(map[string]func(res *Response))

func Get(route string, handler func(res *Response)) {
	routes[route] = handler
}

func Post(route string, handler func(res *Response)) {
	routes[route] = handler
}

func ListenAndServe(port string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)

	fmt.Println("Listening on: " + port)

	if err != nil {
		fmt.Println("Failed to bind to port " + port)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleRequest(conn)
	}
}