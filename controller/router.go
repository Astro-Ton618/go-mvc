package controller

import (
	"fmt"
	"go_mvc/view/css"
	"go_mvc/view/html"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		ctx.SetContentType("text/html")
		html.WriteIndex(ctx, "Index")
	case "/css":
		ctx.SetContentType("text/css")
		css.WriteIndex_style(ctx)
	case "/favicon.ico":
		ctx.SetContentType("image/x-icon")
		fmt.Fprintf(ctx, "%s", get_ico())
	case "/robots.txt":
		fmt.Fprintf(ctx, "user-agent: * allow: /")
	default:
		ctx.SetContentType("text/html")
		html.WriteIndex(ctx, "Not found")
	}
}

func get_ico() string {
	abs, err := filepath.Abs("./view/assets/favicon.ico")
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadFile(abs)
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}
