package linego

func (cl *LineClient) LoginViaToken(authToken string) {
	cl.Talk = createTalkService(LineHost+TalkPath, cl.getDefaultHeader(authToken))
	cl.Poll = createTalkService(LineHost+PollPath, cl.getDefaultHeader(authToken))

	cl.BeforeLogin()
}

type QrLoginClient struct {
	ctx        context.Context
	login      *api.SecondaryQrcodeLoginServiceClient
	loginCheck *api.SecondaryQrCodeLoginPermitNoticeServiceClient
	sessionID  string
}

func NewQrLoginClient() *QrLoginClient {
	return &QrLoginClient{
		ctx: context.Background(),
	}
}
func (cl *LineClient) loginViaQrCode() {

}

func (cl *LineClient) loginViaMail(mail, passwd string) {

}
