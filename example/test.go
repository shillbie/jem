package main

import (
	"github.com/sakura-rip/linego"
	"github.com/sakura-rip/linego/talkservice"
	"log"
)

var bot *linego.LineClient

func main() {
	bot = linego.NewLineClient(talkservice.AppType_ANDROIDLITE)
	bot.Login()
	bot.SendMessage("ud5b1c4b637d684f5714fc4dd2e585565", "login success")
	for {
		ops, err := bot.FetchOps()
		if err == nil {
			for _, op := range ops {
				if op.Type == talkservice.OperationType_END_OF_OPERATION {
					bot.SetGlobalRev(op)
					bot.SetIndividualRev(op)
				} else {
					bot.SetRevision(op.Revision)
					log.Printf("%+v\n", op)
				}
			}
		} else {
			log.Printf("%+v\n", err)
		}
	}
}
