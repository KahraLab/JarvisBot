package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberHttpResponseWriter struct {
	FiberCtx *fiber.Ctx
}

func (r *FiberHttpResponseWriter) Header() http.Header {
	res := http.Header{}
	r.FiberCtx.Response().Header.VisitAll(func(key, value []byte) {
		res.Add(string(key), string(value))
	})
	return res
}

func (r *FiberHttpResponseWriter) Write(i []byte) (int, error) {
	return r.FiberCtx.Write(i)
}

func (r *FiberHttpResponseWriter) WriteHeader(statusCode int) {
	r.FiberCtx.Response().SetStatusCode(statusCode)
}
