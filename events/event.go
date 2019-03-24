package events

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func EventCPBL(client *linebot.Client, event *linebot.Event) {
	client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("you need CPBL info")).Do()
}