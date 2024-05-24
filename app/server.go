package main

import (
	"fmt"
	"os"

	myhttp "github.com/codecrafters-io/http-server-starter-go/app/my-http"
)

func main() {
	myhttp.Get("/", func(res *myhttp.Response, req *myhttp.Request, ctx *myhttp.Context) {
		fmt.Println("Hello from /")
		res.Send(200, "<h1>Hello</h1>")
	})
	myhttp.Get("/echo/:slug", func(res *myhttp.Response, req *myhttp.Request, ctx *myhttp.Context) {
		fmt.Println("Hello from /echo/:slug")
		res.WriteHeader("Content-Type", "text/plain")
		fmt.Println(ctx.Params)
		res.Send(200, ctx.Params["slug"])
	})
	myhttp.Get("/user-agent", func(res *myhttp.Response, req *myhttp.Request, ctx *myhttp.Context) {
		fmt.Println("Hello from /echo/:slug/:slug2")
		res.WriteHeader("Content-Type", "text/plain")
		res.Send(200, req.Headers["User-Agent"])
	})
	myhttp.Get("/files/:file", func(res *myhttp.Response, req *myhttp.Request, ctx *myhttp.Context) {
		directory := os.Args[2]
		data, err := os.ReadFile(directory + "/" + ctx.Params["file"])
		if err != nil {
			res.Send(404, "")
		} else {
			res.WriteHeader("Content-Type", "application/octet-stream")
			res.WriteHeader("Content-Length", fmt.Sprint(len(data)))
			res.Send(200, string(data))
		}
	})

	myhttp.ListenAndServe("4221")
}
