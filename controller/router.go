package controller

import (
	"fmt"

	"go_mvc/view"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "%s", view.Hello("Pippo"))
	default:
		fmt.Fprintf(ctx, "Do not misunderstand it, %q is not power of your creation! ", ctx.RequestURI())
	}
}
