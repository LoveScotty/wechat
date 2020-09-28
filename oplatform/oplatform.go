package oplatform

import (
	"scotty/wechat"
	"scotty/wechat/mp"

	"github.com/xiaojiaoyu100/cast"
)

type Oplatform struct {
	ca *cast.Cast
	Config
}

func New(setters ...Setter) (*Oplatform, error) {
	var err error
	c := new(Oplatform)
	c.ca, err = cast.New(
		cast.WithBaseURL(wechat.BaseUrl),
		cast.WithRetry(3),
	)
	if err != nil {
		return nil, err
	}
	for _, setter := range setters {
		if err = setter(&c.Config); err != nil {
			return nil, err
		}
	}
	return c, nil
}

type Mp struct {
	AppID string `json:"app_id"` // 公众号的app_id
	*mp.Config
}

// 获取代公众号实例
func NewMp(appId string, setters ...mp.Setter) (*Mp, error) {
	m := new(Mp)
	m.AppId = appId
	for _, setter := range setters {
		if err := setter(m.Config); err != nil {
			return nil, err
		}
	}
	return m, nil
}


