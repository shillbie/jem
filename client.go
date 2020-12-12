package linego

import (
	"context"
	api "github.com/sakura-rip/linego/lineapi"
)

type LineClient struct {
	Talk *api.TalkServiceClient
	Poll *api.TalkServiceClient

	ctx     context.Context
	appType api.AppType
	Profile *api.Profile
}

func NewLineClient(appType api.AppType) *LineClient {
	return &LineClient{
		Talk:    nil,
		Poll:    nil,
		appType: appType,
		ctx:     context.Background(),
	}
}
