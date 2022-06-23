package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
	default:
		fmt.Fprintf(ctx, "Do not misunderstand it, %q is not power of your creation! ", ctx.RequestURI())
	}
}
