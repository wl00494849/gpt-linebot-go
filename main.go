package main

import (
	"flag"
	"gpt-linebot-go/src"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	var port string
	flag.StringVar(&port, "p", ":6666", "port")
	src.Bot_Init()

	app := gin.Default()
	app.POST("/bot_callback", src.Callback)
	app.POST("/push_message", src.Push_Message)
	app.Run(port)
}
