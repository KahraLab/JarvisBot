package core

import (
	"errors"
	"jarvis/src/core/routers"
	"os"

	"github.com/chyroc/lark"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func CreateFiberApp() (*fiber.App, error) {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
		return nil, errors.New("Env loading failed!")
	}

	// create fiber app instance
	app := fiber.New()
	// create lark cli via SDK
	larkCli := lark.New(
		lark.WithAppCredential(os.Getenv("LARK_BOT_APP_ID"), os.Getenv("LARK_BOT_APP_SECRET")),
		lark.WithEventCallbackVerify(os.Getenv("LARK_BOT_ENCRYPT_KEY"), os.Getenv("LARK_BOT_APP_VERIFICATION_TOKEN")),
	)

	// register all routers
	jarvisCtx := &routers.JarvisContext {
		LarkCli: larkCli,
	}
	for _, router := range routers.AllRouters() {
		jarvisCtx.Inject(router).Register(app)
	}

	return app, nil
}
