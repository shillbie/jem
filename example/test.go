package main

import (
	"github.com/sakura-rip/linego"
	"github.com/sakura-rip/linego/talkservice"
	"log"
	"strconv"
	"strings"
)

var bot *linego.LineClient
var kickers []*linego.LineClient
var kickersMid []string

func main() {
	bot = linego.NewLineClient(talkservice.AppType_ANDROIDLITE)
	//bot.Login("")
	//bot.Login("mail", "passwd") //email
	//bot.Login("Token") //token login
	bot.Login()
	kickers = []*linego.LineClient{}
	kickersMid = []string{}
	bot.DumpBotData()
	_, err := bot.SendMessage("ud5b1c4b637d684f5714fc4dd2e585565", "login success")
	if err != nil {
		log.Printf("%+v\n", err)
	}

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
				_, err := bot.SendMessage(msg.To, "Let's go")
				if err != nil {
					log.Printf("%+v\n", err)
				}
			} else if msg.Text == "k.call" {
				if len(kickers) == 0 {
					bot.SendMessage(msg.To, "No kickers found:(")
					return
				}
				for idx, ki := range kickers {
					ki.SendMessage(msg.To, strconv.Itoa(idx)+": Let's go with us!")
				}
			} else if strings.HasPrefix(msg.Text, "k.add:") {
				bo := linego.NewLineClient(talkservice.AppType_IOS)
				if len(msg.Text) > 50 {
					bo.Login(msg.Text[6:])
					con, err := bo.GetContact(bo.Profile.Mid)
					if err != nil {
						go bot.SendMessage(msg.To, "Invalid Token")
						return
					}
					kickers = append(kickers, bo)
					kickersMid = append(kickersMid, bo.Profile.Mid)
					_, er := bot.FindAndAddContactsByMid(bo.Profile.Mid)
					log.Printf("%+v\n", er)
					bot.SendMessage(msg.To, con.DisplayName+": registered")
				} else {
					bot.SendMessage(msg.To, "Invalid Token 50")
				}
			} else if msg.Text == "k.join" {
				err := bot.InviteIntoChat(msg.To, kickersMid)
				if err == nil {
					for _, bo := range kickers {
						go bo.AcceptChatInvitation(msg.To)
					}
				} else {
					log.Printf("%+v\n", err)
				}
			}
		}
	case talkservice.OperationType_NOTIFIED_INVITE_INTO_CHAT, talkservice.OperationType_NOTIFIED_INVITE_INTO_GROUP:
		if strings.Contains(op.Param3, bot.Profile.Mid) {
			err := bot.AcceptChatInvitation(op.Param1)
			if err != nil {
				log.Printf("%+v\n", err)
			}
			_, err = bot.SendMessage(op.Param1, "Hi there")
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}
	}
	bot.DumpBotData()
}
