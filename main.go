package main

import (
	"go_mvc/controller"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Printf("Listening on port : 8080")
	panic(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) { controller.Router(ctx) }))
}
