package controller

import (
	"fmt"
	"go_mvc/view/css"
	"go_mvc/view/html"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		ctx.SetContentType("text/html; charset=utf-8")
		fmt.Fprintf(ctx, "%s", html.Home("go_mvc", "go_mvc"))
	case "/css":
		ctx.SetContentType("text/css; charset=utf-8")
		fmt.Fprintf(ctx, "%s", css.Home_style())
	case "/robots.txt":
		fmt.Fprintf(ctx, "user-agent: * allow: /")
	default:
		ctx.SetContentType("text/html; charset=utf-8")
		fmt.Fprintf(ctx, "%s", html.Not_found("not_found", "not_found"))
	}
}
