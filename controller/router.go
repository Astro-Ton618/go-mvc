package controller

import (
	"fmt"
	"go_mvc/view"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "%s", view.Home("go_mvc", "go_mvc"))
		ctx.SetContentType("text/html; charset=utf-8")
	case "/robots.txt":
		fmt.Fprintf(ctx, "user-agent: * allow: /")
	default:
		fmt.Fprintf(ctx, "Do not misunderstand it, %q is not power of your creation!", ctx.RequestURI())
	}
}
