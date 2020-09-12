package oplatform

import (
	"scotty/wechat"

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
		cast.WithRetry(2),
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
