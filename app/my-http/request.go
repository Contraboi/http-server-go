package myhttp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func HandleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	req := createRequest(buf)

	if req == nil {
		fmt.Println("Error creating request")
		return
	}

	node, params := router.Search(req.Path)

	if node == nil {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	node.handler(CreateResponse(req, conn), &Context{Params: params})
}

func createRequest(buf []byte) *Request {
	reader := bufio.NewReader(strings.NewReader(string(buf)))
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading status line:", err)
		return nil
	}
	statusLine = strings.TrimSpace(statusLine)

	// Parse the headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading headers:", err)
			return nil
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Malformed header:", line)
			continue
		}
		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	for key, value := range headers {
		fmt.Println(key+":", value)
	}
	// Read the body
	var bodyBuilder strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading body:", err)
			}
			break
		}
		bodyBuilder.WriteString(line)
	}
	body := bodyBuilder.String()

	return &Request{
		Method:  strings.Fields(statusLine)[0],
		Path:    strings.Fields(statusLine)[1],
		Headers: headers,
		Body:    body,
	}
}
