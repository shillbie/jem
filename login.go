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
	qrL := NewQrLoginClient()
	qrL.createLoginSession1()
	qrL.CreateQrSession()
	qrL.createLoginSession2()
	url, er := qrL.CreateQrCode()
	if er != nil {
		log.Printf("%+v\n", er)
	}
	fmt.Println("login this url on your mobile : " + url)

	qrL.WaitForQrCodeVerified()
	err := qrL.CertificateLogin("")
	if err != nil {
		pin, _ := qrL.CreatePinCode()
		fmt.Println("input this pin code on your mobile : " + pin)
		qrL.WaitForInputPinCode()
	}
	token, _, _ := qrL.QrLogin()

	cl.LoginViaToken(token)
}

func (cl *LineClient) loginViaMail(mail, passwd string) {

}
