package main

import (
	"fmt"
	"go_mvc/controller"

	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Printf("Listening on port : 8080")
	panic(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) { controller.Router(ctx) }))
}
