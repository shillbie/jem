package linego

import (
	"context"
	"fmt"
	api "github.com/sakura-rip/linego/talkservice"
)

type LineClient struct {
	Talk *api.TalkServiceClient
	Poll *api.TalkServiceClient
	Liff *api.LiffServiceClient

	ctx           context.Context
	appType       api.AppType
	Profile       *api.Profile
	SaveData      *BotSavedData
	reqSeq        int32
	reqSeqMessage map[string]int32
}

func NewLineClient(appType api.AppType) *LineClient {
	return &LineClient{
		ctx:           context.Background(),
		appType:       appType,
		Profile:       &api.Profile{},
		reqSeqMessage: map[string]int32{},
	}
}

func (cl *LineClient) Login(attribute ...string) {
	switch len(attribute) {
	case 0:
		cl.loginViaQrCode()
	case 1:
		cl.LoginViaToken(attribute[0], false)
	case 2:
		cl.loginViaMail(attribute[0], attribute[1])
	}
	fmt.Println("Login Done")
	cl.BeforeLogin()
}
