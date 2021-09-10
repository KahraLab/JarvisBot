package routers

import (
	"github.com/chyroc/lark"
	"github.com/gofiber/fiber/v2"
)

type RouterMethodType string

type JarvisContext struct {
	LarkCli *lark.Lark
}

func (jctx *JarvisContext) Inject(router *RouterConfig) *RouterConfig {
	router.jarvisCtx = jctx
	return router
}

type RouterConfig struct {
	Method    string
	Path      string
	Handler   func(*fiber.Ctx, *JarvisContext) error
	jarvisCtx *JarvisContext
}

func (r *RouterConfig) Register(app *fiber.App) {
	handler := r.FiberHandler()
	path := r.Path

	switch r.Method {
	case fiber.MethodGet:
		app.Get(path, handler)
	case fiber.MethodHead:
		app.Head(path, handler)
	case fiber.MethodPost:
		app.Post(path, handler)
	case fiber.MethodPut:
		app.Put(path, handler)
	case fiber.MethodPatch:
		app.Patch(path, handler)
	case fiber.MethodDelete:
		app.Delete(path, handler)
	case fiber.MethodConnect:
		app.Connect(path, handler)
	case fiber.MethodOptions:
		app.Options(path, handler)
	case fiber.MethodTrace:
		app.Trace(path, handler)
	}
}

func (r *RouterConfig) FiberHandler() func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		return r.Handler(fiberCtx, r.jarvisCtx)
	}
}

func AllRouters() []*RouterConfig {
	return []*RouterConfig{
		// * 在此添加路由
		{Method: fiber.MethodGet, Path: "/", Handler: routerPing},
	}
}
