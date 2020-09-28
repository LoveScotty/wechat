package mp

import (
	"scotty/wechat"

	"github.com/xiaojiaoyu100/cast"
)

type Mp struct {
	ca *cast.Cast
	Config
}

func New(setters ...Setter) (*Mp, error) {
	var err error
	c := new(Mp)
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
