package main

import (
	"github.com/sakura-rip/linego"
	"github.com/sakura-rip/linego/talkservice"
	"log"
	"strings"
)

var bot *linego.LineClient

func main() {
	bot = linego.NewLineClient(talkservice.AppType_ANDROIDLITE)
	bot.Login() //QRLogin
	//bot.Login("mail", "passwd") //email
	//bot.Login("Token") //token login
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
					HandleOperation(op)
				}
			}
		} else {
			log.Printf("%+v\n", err)
		}
	}
}

func HandleOperation(op *talkservice.Operation) {
	log.Printf("%+v\n", op)
	switch op.Type {
	case talkservice.OperationType_RECEIVE_MESSAGE, talkservice.OperationType_SEND_MESSAGE:
		msg := op.Message
		if msg.ContentType == talkservice.ContentType_NONE {
			if msg.Text == "test" {
				bot.SendMessage(msg.To, "Yes")
			}
		}
	case talkservice.OperationType_NOTIFIED_INVITE_INTO_CHAT:
		if strings.Contains(op.Param3, bot.Profile.Mid) {
			bot.AcceptChatInvitation(op.Param1)
			bot.SendMessage(op.Param1, "Hi there")
		}
	}
}
