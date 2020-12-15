package linego

import (
	"bytes"
	"encoding/json"
	api "github.com/sakura-rip/linego/talkservice"
	"net/http"
)

func (cl *LineClient) SendTemplate(to string, data interface{}) {
	contextChat := api.NewLiffChatContext()
	contextChat.ChatMid = to

	context := api.NewLiffContext()
	context.Chat = contextChat

	req := api.NewLiffViewRequest()
	req.LiffId = "1602687308-GXq4Vvk9"
	req.Context = context

	token, err := cl.Liff.IssueLiffView(cl.ctx, req)
	if err != nil {
		cl.SendMessage(to, "Please Allow liff for template\nNEO TEAM BOTS : line://app/1647614612-DGev40WM")
		return
	}

	url := "https://api.line.me/message/v3/share"

	data = map[string]interface{}{"messages": []interface{}{data}}
	jsonStr, _ := json.Marshal(data)
	hReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if hReq == nil {
		return
	}
	hReq.Header.Set("Authorization", "Bearer "+token.AccessToken)
	hReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(hReq)
}
