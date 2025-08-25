package src

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

var bot *linebot.Client
var client *gpt

type Push_Message_Req struct {
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

func Bot_Init() {
	b, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	client = NewGPT()
	bot = b
}

func Callback(ctx *gin.Context) {
	events, err := bot.ParseRequest(ctx.Request)

	if err == linebot.ErrInvalidSignature {
		log.Println(err)
		ctx.Status(500)
	}

	for _, evn := range events {
		if evn.Type == linebot.EventTypeMessage {
			if msg, ok := evn.Message.(*linebot.TextMessage); ok {
				fmt.Println(evn.Source.UserID)
				resp := client.Requset(msg.Text, evn.Source.UserID)
				_, err := bot.ReplyMessage(evn.ReplyToken, linebot.NewTextMessage(resp)).Do()
				if err != nil {
					log.Println(err)
					ctx.Status(500)
				}
			}
		}
	}
	ctx.Status(200)
}

func Push_Message(ctx *gin.Context) {

	var data Push_Message_Req

	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		ctx.Status(500)
	}

	bot.PushMessage(data.UserID, linebot.NewTextMessage(data.Message)).Do()
	ctx.Status(200)
}
