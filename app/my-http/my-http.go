package myhttp

import (
	"fmt"
	"net"
	"os"
)

var router = NewRouter()

func Get(route string, handler HandlerFunc) {
	router.Insert(route, "GET", handler)
}

func Post(route string, handler HandlerFunc) {
	router.Insert(route, "POST", handler)
}

func ListenAndServe(port string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)

	fmt.Println("Listening on: " + port)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
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
