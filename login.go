package linego

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	api "github.com/sakura-rip/linego/talkservice"
	"log"
	"net/http"
)

func (cl *LineClient) LoginViaToken(authToken string, isSeq bool) {
	header := cl.getDefaultHeader(authToken)
	if isSeq {
		header["X-Line-Application"] = header["X-Line-Application"] + ";SECONDARY"
	}
	cl.Profile.Mid = authToken[:33]
	cl.Talk = createTalkService(LineHost+TalkPath, header)
	cl.Poll = createTalkService(LineHost+PollPath, header)
}

type QrLoginClient struct {
	ctx        context.Context
	login      *api.SecondaryQrcodeLoginServiceClient
	loginCheck *api.SecondaryQrCodeLoginPermitNoticeServiceClient
	sessionID  string
	appType    api.AppType
}

func NewQrLoginClient() *QrLoginClient {
	return &QrLoginClient{
		ctx: context.Background(),
	}
}
func (cl *LineClient) loginViaQrCode() {
	qrL := NewQrLoginClient()
	qrL.appType = cl.appType
	qrL.createLoginSession1()
	qrL.CreateQrSession()
	qrL.createLoginCheckSession()
	url, er := qrL.CreateQrCode()
	if er != nil {
		log.Printf("%+v\n", er)
	}
	fmt.Println("login this url on your mobile :\n" + url)

	qrL.WaitForQrCodeVerified()
	err := qrL.CertificateLogin("")
	if err != nil {
		pin, _ := qrL.CreatePinCode()
		fmt.Println("input this pin code on your mobile :\n" + pin)
		qrL.WaitForInputPinCode()
	}
	token, cert, _ := qrL.QrLogin()
	fmt.Print("cert: " + cert + "\n")
	fmt.Printf("token:" + token)

	cl.LoginViaToken(token, true)
}

func (cl *LineClient) loginViaMail(mail, passwd string) {

}

func (cl *LineClient) getDefaultHeader(authToken string) map[string]string {
	return map[string]string{
		"X-Line-Access":      authToken,
		"X-Line-Application": cl.GetLineApp(),
		"User-Agent":         cl.GetUserAgent(),
		"x-lal":              cl.GetXLal(),
	}
}

func createThriftClient(url string, header map[string]string) *thrift.TStandardClient {
	var transport thrift.TTransport

	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	transport, _ = thrift.NewTHttpClientWithOptions(url, option)
	connect := transport.(*thrift.THttpClient)
	for k, v := range header {
		connect.SetHeader(k, v)
	}
	pCol := thrift.NewTCompactProtocol(transport)
	return thrift.NewTStandardClient(pCol, pCol)
}

func createTalkService(url string, header map[string]string) *api.TalkServiceClient {
	return api.NewTalkServiceClient(createThriftClient(url, header))
}

func createSqLoginService(url string, header map[string]string) *api.SecondaryQrcodeLoginServiceClient {
	return api.NewSecondaryQrcodeLoginServiceClient(createThriftClient(url, header))

}
func createSqLoginCheckService(url string, header map[string]string) *api.SecondaryQrCodeLoginPermitNoticeServiceClient {
	return api.NewSecondaryQrCodeLoginPermitNoticeServiceClient(createThriftClient(url, header))
}
