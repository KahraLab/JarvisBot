package lark

import (
	"context"
	"github.com/chyroc/lark"
)

func HandlerReceiveTextMessage(
	ctx context.Context,
	cli *lark.Lark,
	schema string,
	header *lark.EventHeaderV1,
	event *lark.EventV1ReceiveMessage,
) (string, error) {
	// parse text content as Command
	textContent := event.TextWithoutAtBot
	// TODO:
	// command := parseCommand(textContent)
	// replyMsg, commandErr := command.Execute()
	// if commandErr != nil {
	//   log.Fatalln("run command fail, ", commandErr)
	// }

	// send reply
	_, _, cliErr := cli.Message.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{
		ChatID:  event.OpenChatID,
		RootID:  &event.RootID,
		MsgType: lark.MsgTypeText,
		Content: &lark.SendRawMessageOldReqContent{
			Text: "resp: " + textContent,
		},
	})
	return "", cliErr
}
