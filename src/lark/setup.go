package lark

import (
	"github.com/chyroc/lark"
)

func SetupLarkClient(larkCli *lark.Lark) {
	// handle text messages and reply
	larkCli.EventCallback.HandlerEventV1ReceiveMessage(HandlerReceiveTextMessage)
}
