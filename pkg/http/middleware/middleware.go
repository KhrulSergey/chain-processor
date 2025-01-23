package middleware

import "github.com/valyala/fasthttp"

func CORSDisablingMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, "*")
		ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowCredentials, "true")

		h(ctx)
	}
}
