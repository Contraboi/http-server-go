package main

import (
	"fmt"
	"os"

	myhttp "github.com/codecrafters-io/http-server-starter-go/app/my-http"
)

func main() {
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
		fmt.Println(directory + "/" + ctx.Params["file"])
		if err != nil {
			res.NotFound()
		} else {
			res.WriteHeader("Content-Type", "application/octet-stream")
			res.Send(200, string(data))
		}
	})
	myhttp.Post("/files/:file", func(res *myhttp.Response, req *myhttp.Request, ctx *myhttp.Context) {
		directory := os.Args[2]

		err := os.WriteFile(directory+"/"+ctx.Params["file"], []byte(req.Body), 0644)
		if err != nil {
			res.Send(500, "Error writing file")
		} else {
			res.WriteHeader("Content-Type", "text/plain")
			res.Send(201, req.Body)
		}

	})

	myhttp.ListenAndServe("4221")
}
