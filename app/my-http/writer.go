package myhttp

import (
	"fmt"
	"net"
)

type Context struct {
	Params map[string]string
}

type Response struct {
	Request *Request
	conn    net.Conn
	headers map[string]string
	body    string
}

const (
	OK        = 200
	NOT_FOUND = 400
	CREATED   = 201
)

var statusText = map[int]string{
	OK:        "200 OK",
	NOT_FOUND: "404 NOT_FOUND",
	CREATED:   "201 Created",
}

func CreateResponse(req *Request, conn net.Conn) *Response {
	return &Response{
		Request: req,
		conn:    conn,
		headers: make(map[string]string),
	}
}

func (res *Response) WriteHeader(key string, value string) {
	res.headers[key] = value
}

func (res *Response) NotFound() {
	res.conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
}
func (res *Response) Send(status int, body string) {
	if len(body) > 0 {
		res.WriteHeader(`Content-Length`, fmt.Sprint(len(body)))
	}

	dataToSend := "HTTP/1.1 " + statusText[status] + "\r\n"
	for key, value := range res.headers {
		dataToSend += key + ": " + value + "\r\n"
	}
	dataToSend += "\r\n" + body

	res.conn.Write([]byte(dataToSend))
}
