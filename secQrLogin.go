package linego

import (
	api "github.com/shillbie/jem/talkservice"
	"log"
)

func (cl *QrLoginClient) createLoginSession1() {
	cl.login = createSqLoginService(LineHost+SQLogin, map[string]string{
		"X-Line-Application": cl.GetLineApp() + ";SECONDARY",
		"User-Agent":         cl.GetUserAgent(),
		"x-lal":              cl.GetXLal(),
	})
}

func (cl *QrLoginClient) CreateQrSession() {
	req := api.NewCreateQrSessionRequest()
	res, err := cl.login.CreateSession(cl.ctx, req)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	cl.sessionID = res.AuthSessionId
}

func (cl *QrLoginClient) createLoginCheckSession() {
	cl.loginCheck = createSqLoginCheckService(LineHost+SQLoginCheck, map[string]string{
		"X-Line-Application": cl.GetLineApp() + ";SECONDARY",
		"User-Agent":         cl.GetUserAgent(),
		"x-lal":              cl.GetXLal(),
		"X-Line-Access":      cl.sessionID,
	})
}

func (cl *QrLoginClient) CreateQrCode() (string, error) {
	req := api.NewCreateQrCodeRequest()
	req.AuthSessionId = cl.sessionID
	res, err := cl.login.CreateQrCode(cl.ctx, req)
	if res != nil {
		return res.CallbackUrl, err
	}
	return "", err

}
func (cl *QrLoginClient) WaitForQrCodeVerified() {
	req := api.NewCheckQrCodeVerifiedRequest()
	req.AuthSessionId = cl.sessionID
	err := cl.loginCheck.CheckQrCodeVerified(cl.ctx, req)
	if err != nil {
		panic(err)
	}
}

func (cl *QrLoginClient) CertificateLogin(cert string) error {
	req := api.NewVerifyCertificateRequest()
	req.AuthSessionId = cl.sessionID
	req.Certificate = cert
	err := cl.login.VerifyCertificate(cl.ctx, req)
	return err
}
func (cl *QrLoginClient) CreatePinCode() (string, error) {
	req := api.NewCreatePinCodeRequest()
	req.AuthSessionId = cl.sessionID
	res, err := cl.login.CreatePinCode(cl.ctx, req)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return res.PinCode, err
}
func (cl *QrLoginClient) WaitForInputPinCode() {
	req := api.NewCheckPinCodeVerifiedRequest()
	req.AuthSessionId = cl.sessionID
	err := cl.loginCheck.CheckPinCodeVerified(cl.ctx, req)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
func (cl *QrLoginClient) QrLogin() (string, string, error) {
	req := api.NewQrCodeLoginRequest()
	req.AuthSessionId = cl.sessionID
	req.SystemName = "SakuraBOT"
	req.AutoLoginIsRequired = true
	callback, err := cl.login.QrCodeLogin(cl.ctx, req)
	if callback != nil {
		return callback.AccessToken, callback.Certificate, err
	}
	return "", "", err
}
