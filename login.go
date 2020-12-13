package linego

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	api "github.com/sakura-rip/linego/lineapi"
	"log"
	"net/http"
)

func (cl *LineClient) LoginViaToken(authToken string) {
	cl.Talk = createTalkService(LineHost+TalkPath, cl.getDefaultHeader(authToken))
	cl.Poll = createTalkService(LineHost+PollPath, cl.getDefaultHeader(authToken))
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

func (cl *LineClient) getDefaultHeader(authToken string) map[string]string {
	return map[string]string{
		"X-Line-Access":      authToken,
		"X-Line-Application": cl.GetLineApp(),
		"User-Agent":         cl.GetUserAgent(),
		"x-lal":              cl.GetXLal(),
	}
}

func createTalkService(url string, header map[string]string) *api.TalkServiceClient {
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
	tStc := thrift.NewTStandardClient(pCol, pCol)
	return api.NewTalkServiceClient(tStc)
}
