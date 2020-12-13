package main

import (
	"github.com/sakura-rip/linego"
	"github.com/sakura-rip/linego/talkservice"
)

var bot *linego.LineClient

func main() {
	bot = linego.NewLineClient(talkservice.AppType_ANDROIDLITE)
	bot.Login()
	bot.SendMessage("ud5b1c4b637d684f5714fc4dd2e585565", "login success")
}
