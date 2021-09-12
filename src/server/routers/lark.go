package routers

import (
	"bytes"
	"jarvis/src/utils"

	"github.com/gofiber/fiber/v2"
)

func routerStartListenLarkEventCallback(fiberCtx *fiber.Ctx, jarvisCtx *JarvisContext) error {
	jarvisCtx.LarkCli.EventCallback.ListenCallback(
		fiberCtx.UserContext(),
		bytes.NewReader(fiberCtx.Body()),
		&utils.FiberHttpResponseWriter{FiberCtx: fiberCtx},
	)
	return nil
}
