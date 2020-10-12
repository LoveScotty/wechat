package mp

import "scotty/wechat"

type Url string

const (
	UrlCustomSend Url = "cgi-bin/message/custom/send?access_token=%s"
)

func (u Url) Path() string {
	return string(u)
}

func (u Url) AllPath() string {
	return wechat.BaseUrl + u.Path()
}
