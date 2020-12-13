package linego

import api "github.com/sakura-rip/linego/lineapi"

const LineHost = "https://legy-jp-addr-long.line.naver.jp"
const TalkPath = "/S4"
const PollPath = "/P4"

var systemVersion = map[string]string{
	"LITE":   "10.0",
	"MAC":    "10.15.1",
	"CHROME": "1",
	"IOS":    "13.4.1",
}

var appVersion = map[string]string{
	"LITE":   "2.14.0",
	"MAC":    "5.24.1",
	"CHROME": "2.3.9",
	"IOS":    "10.9.0",
}

func (cl *LineClient) GetLineApp() string {
	switch cl.appType {
	case api.AppType_ANDROIDLITE:
		return "ANDROIDLITE\t" + appVersion["LITE"] + "\tAndroid OS\t" + systemVersion["LITE"]
	case api.AppType_IOS:
		return "IOS\t" + appVersion["IOS"] + "\tiOS\t" + systemVersion["IOS"]
	case api.AppType_CHROMEOS:
		return "CHROMEOS\t" + appVersion["CHROME"] + "\tChrome_OS\t" + systemVersion["CHROME"]
	case api.AppType_DESKTOPMAC:
		return "DESKTOPMAC\t" + appVersion["MAC"] + "\tOS X\t" + systemVersion["MAC"]

	case api.AppType_DESKTOPWIN:
	case api.AppType_ANDROID:
	case api.AppType_IOSIPAD:
	default:
		panic("invalid app type")
	}
	return ""
}

func (cl *LineClient) GetUserAgent() string {
	switch cl.appType {
	case api.AppType_ANDROIDLITE:
		return "LLA/" + systemVersion["LITE"] + " Galaxy Note 10+ " + systemVersion["LITE"]
	case api.AppType_IOS:
		return "Line/" + appVersion["IOS"] + " iPhone8,1 " + systemVersion["IOS"]
	case api.AppType_CHROMEOS:
		return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"
	case api.AppType_DESKTOPMAC:
		return "Line/" + systemVersion["MAC"]

	case api.AppType_DESKTOPWIN:
	case api.AppType_IOSIPAD:
	case api.AppType_ANDROID:
	default:
		panic("invalid app type")
	}
	return ""
}
func (cl *LineClient) GetXLal() string {
	switch cl.appType {
	case api.AppType_CHROMEOS:
		return "ja"
	default:
		return "jp_ja"
	}
}
