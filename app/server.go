package main

import (
	"fmt"
	myhttp "github.com/codecrafters-io/http-server-starter-go/app/my-http"
)

func main() {
	myhttp.Get("/", func(res *myhttp.Response, ctx *myhttp.Context) {
		fmt.Println("Hello from /")
		res.Send(200, "<h1>Hello</h1>")
	})

	myhttp.Get("/banana", func(res *myhttp.Response, ctx *myhttp.Context) {
		res.Send(200, "")
	})
	myhttp.Get("/echo/:slug", func(res *myhttp.Response, ctx *myhttp.Context) {
		fmt.Println("Hello from /echo/:slug")
		res.WriteHeader("Content-Type", "text/plain")
		fmt.Println(ctx.Params)
		res.Send(200, ctx.Params["slug"])
	})

	myhttp.ListenAndServe("4221")
}
