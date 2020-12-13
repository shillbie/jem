package linego

func (cl *LineClient) LoginViaToken(authToken string) {
	cl.Talk = createTalkService(LineHost+TalkPath, cl.getDefaultHeader(authToken))
	cl.Poll = createTalkService(LineHost+PollPath, cl.getDefaultHeader(authToken))

	cl.BeforeLogin()
}

}
func (cl *LineClient) loginViaQrCode() {

}

func (cl *LineClient) loginViaMail(mail, passwd string) {

}
