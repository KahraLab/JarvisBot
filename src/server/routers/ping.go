package routers

import "github.com/gofiber/fiber/v2"

func routerPing(fiberCtx *fiber.Ctx, jarvisCtx *JarvisContext) error {
	return fiberCtx.SendString("Hello Sir, Jarvis Here.")
}
