package src

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type PushMessageReq struct {
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

type LintBot struct {
	bot *linebot.Client
	gpt *gpt
}

func NewLineBot(line_token string, line_secret string, gpt_url string) *LintBot {
	bot, err := linebot.New(
		line_secret,
		line_token,
	)

	gpt := NewGPT(gpt_url)

	if err != nil {
		log.Fatal(err)
	}

	return &LintBot{
		bot: bot,
		gpt: gpt,
	}
}

func (line_bot *LintBot) Callback(ctx *gin.Context) {
	events, err := line_bot.bot.ParseRequest(ctx.Request)

	if err == linebot.ErrInvalidSignature {
		log.Println(err)
		ctx.Status(500)
		return
	}

	for _, evn := range events {
		if evn.Type == linebot.EventTypeMessage {
			if msg, ok := evn.Message.(*linebot.TextMessage); ok {
				fmt.Println(evn.Source.UserID)
				resp := line_bot.gpt.Requset(msg.Text, evn.Source.UserID)
				_, err := line_bot.bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage(resp)).Do()
				if err != nil {
					log.Println(err)
					ctx.Status(500)
					return
				}
			}
		}
	}
	ctx.Status(200)
}

func (line_bot *LintBot) Push_Message(ctx *gin.Context) {

	var data PushMessageReq

	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		ctx.Status(500)
		return
	}

	line_bot.bot.PushMessage(data.UserID, linebot.NewTextMessage(data.Message)).Do()
	ctx.Status(200)
}
