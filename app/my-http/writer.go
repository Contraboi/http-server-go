package myhttp

import (
	"fmt"
	"net"
	"strings"
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

var ACCCEPTED_ENCODINGS = [1]string{"gzip"}

func (res *Response) Send(status int, body string) {
	if len(body) > 0 {
		res.WriteHeader(`Content-Length`, fmt.Sprint(len(body)))
	}

	encoding := strings.Split(res.Request.Headers["Accept-Encoding"], ",")
	if len(encoding) > 0 {
		for _, enc := range encoding {
			for _, acc := range ACCCEPTED_ENCODINGS {
				if strings.TrimSpace(enc) == acc {
					res.WriteHeader(`Content-Encoding`, acc)
				}
			}
		}
	}

	dataToSend := "HTTP/1.1 " + statusText[status] + "\r\n"
	for key, value := range res.headers {
		dataToSend += key + ": " + value + "\r\n"
	}
	dataToSend += "\r\n" + body

	res.conn.Write([]byte(dataToSend))
}
