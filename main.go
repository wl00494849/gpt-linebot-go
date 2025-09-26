package main

import (
	"flag"
	"gpt-linebot-go/src"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	line_secret    = os.Getenv("LINE_CHANNEL_SECRET")
	line_token     = os.Getenv("LINE_CHANNEL_TOKEN")
	python_gpt_url = os.Getenv("PYTHON_GPT_URL")
	port           string
)

func init() {
	flag.StringVar(&port, "p", ":6666", "port")
	if line_secret == "" || line_token == "" || python_gpt_url == "" {
		panic("load env fail")
	}
}

func main() {
	lint_bot := src.NewLine_Bot(line_token, line_secret, python_gpt_url)

	app := gin.Default()

	app.POST("/bot_callback", lint_bot.Callback)
	app.POST("/push_message", lint_bot.Push_Message)
	app.GET("/live", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	app.Run(port)
}
