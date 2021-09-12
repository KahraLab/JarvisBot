package server

import (
	"errors"
	"jarvis/src/lark"
	"os"

	"jarvis/src/server/routers"

	larkSdk "github.com/chyroc/lark"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func CreateFiberApp() (*fiber.App, error) {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
		return nil, errors.New("environment variables loading failed")
	}

	// create lark cli via SDK
	larkCli := larkSdk.New(
		larkSdk.WithAppCredential(os.Getenv("LARK_BOT_APP_ID"), os.Getenv("LARK_BOT_APP_SECRET")),
		larkSdk.WithEventCallbackVerify(os.Getenv("LARK_BOT_ENCRYPT_KEY"), os.Getenv("LARK_BOT_APP_VERIFICATION_TOKEN")),
	)

	// setup every platform sdk client
	lark.SetupLarkClient(larkCli)

	// create fiber app instance
	app := fiber.New()
	// register all routers
	jarvisCtx := &routers.JarvisContext{
		LarkCli: larkCli,
	}
	for _, router := range routers.AllRouters() {
		jarvisCtx.Inject(router).Register(app)
	}

	return app, nil
}
