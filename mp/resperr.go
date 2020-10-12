package mp

import "scotty/wechat"

type RespError struct {
	StatusCode int    `json:"-"`
	ErrCode    *int   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (err RespError) RequestSuccess() bool {
	// 因为0是默认值，所以要加200判断
	if err.ErrCode == nil {
		return true
	}

	return *err.ErrCode == wechat.MpSuccess
}
