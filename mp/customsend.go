package mp

import (
	"context"
	"fmt"

	"scotty/wechat/mp/message"
)

//Send 发送客服消息
func (m *Mp) SendCustomerMsg(ctx context.Context, accessToken string, param *message.CustomerMessage) error {
	path := fmt.Sprintf(UrlCustomSend.AllPath(), accessToken)
	request := m.ca.NewRequest().Post().WithJSONBody(param).WithPath(path)
	resp, err := m.ca.Do(ctx, request)
	if err != nil {
		return fmt.Errorf("cast do failed, err is %w", err)
	}
	res := new(RespError)
	if err = resp.DecodeFromJSON(&res); err != nil {
		return fmt.Errorf("decode failed, err is %w", err)
	}
	res.StatusCode = resp.StatusCode()
	if !res.RequestSuccess() {
		return fmt.Errorf("SendCustomerMsg error, err is code:%v, msg:%v", *res.ErrCode, res.ErrMsg)
	}
	return nil
}
