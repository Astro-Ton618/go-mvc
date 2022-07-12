package main

import (
	"go-mvc/controller"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Println("\nListening on port : 8080")
	panic(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) { controller.Router(ctx) }))
}
