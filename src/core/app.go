package core

import (
	"context"
	"errors"
	"os"

	"jarvis/src/core/routers"

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

	// 这里监听文本消息，并回复
	larkCli.EventCallback.HandlerEventV1ReceiveMessage(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV1, event *lark.EventV1ReceiveMessage) (string, error) {
		_, _, err := cli.Message.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{
			ChatID:  event.OpenChatID,
			RootID:  &event.RootID,
			MsgType: lark.MsgTypeText,
			Content: &lark.SendRawMessageOldReqContent{
				Text: "resp: " + event.TextWithoutAtBot,
			},
		})
		return "", err
	})

	// register all routers
	jarvisCtx := &routers.JarvisContext{
		LarkCli: larkCli,
	}
	for _, router := range routers.AllRouters() {
		jarvisCtx.Inject(router).Register(app)
	}

	return app, nil
}
