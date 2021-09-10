package routers

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func routerPingLark(fiberCtx *fiber.Ctx, jarvisCtx *JarvisContext) error {
	jarvisCtx.LarkCli.EventCallback.ListenCallback(
		fiberCtx.UserContext(),
		bytes.NewReader(fiberCtx.Body()),
		&fiberHttpResponseWriter{fiberCtx: fiberCtx},
	)
	return nil
}

type fiberHttpResponseWriter struct {
	fiberCtx *fiber.Ctx
}

func (r *fiberHttpResponseWriter) Header() http.Header {
	res := http.Header{}
	r.fiberCtx.Response().Header.VisitAll(func(key, value []byte) {
		res.Add(string(key), string(value))
	})
	return res
}

func (r *fiberHttpResponseWriter) Write(i []byte) (int, error) {
	return r.fiberCtx.Write(i)
}

func (r *fiberHttpResponseWriter) WriteHeader(statusCode int) {
	r.fiberCtx.Response().SetStatusCode(statusCode)
}
