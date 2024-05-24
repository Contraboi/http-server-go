package main

import (
	"fmt"
	myhttp "github.com/codecrafters-io/http-server-starter-go/app/my-http"
)

func main() {
	myhttp.Get("/", func(res *myhttp.Response) {
		fmt.Println("Hello from /")
		res.Send(200, "<h1>Hello</h1>")
	})
	myhttp.Get("/about", func(res *myhttp.Response) {
		fmt.Println("Hello from /about")
		res.Send(200, "<h1>About</h1>")
	})
	myhttp.Get("/echo/:slug", func(res *myhttp.Response) {
		fmt.Println("Hello from /contact")
		res.WriteHeader("Content-Type", "text/plain")
		res.Send(200, "abc")
	})

	myhttp.ListenAndServe("4221")
}
