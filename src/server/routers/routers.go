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
	Group     string
	Method    string
	Path      string
	Handler   func(*fiber.Ctx, *JarvisContext) error
	jarvisCtx *JarvisContext
}

func (r *RouterConfig) Register(reciever fiber.Router) {
	handler := r.FiberHandler()
	path := r.Path

	switch r.Method {
	case fiber.MethodGet:
		reciever.Get(path, handler)
	case fiber.MethodHead:
		reciever.Head(path, handler)
	case fiber.MethodPost:
		reciever.Post(path, handler)
	case fiber.MethodPut:
		reciever.Put(path, handler)
	case fiber.MethodPatch:
		reciever.Patch(path, handler)
	case fiber.MethodDelete:
		reciever.Delete(path, handler)
	case fiber.MethodConnect:
		reciever.Connect(path, handler)
	case fiber.MethodOptions:
		reciever.Options(path, handler)
	case fiber.MethodTrace:
		reciever.Trace(path, handler)
	}
}

func (r *RouterConfig) FiberHandler() func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		return r.Handler(fiberCtx, r.jarvisCtx)
	}
}

func defineRouterGroup(groupName string, routers []*RouterConfig) []*RouterConfig {
	for _, router := range routers {
		router.Group = groupName
	}
	return routers
}

// * export all routers here
func AllRouters() []*RouterConfig {
	allRouters := []*RouterConfig{
		{Method: fiber.MethodGet, Path: "/", Handler: routerPing},
	}

	larkRouters := defineRouterGroup("lark", []*RouterConfig{
		// * export all Lark routers - path: "/lark/..."
		{Method: fiber.MethodPost, Path: "/callback", Handler: routerStartListenLarkEventCallback},
	})
	allRouters = append(allRouters, larkRouters...)

	return allRouters
}
