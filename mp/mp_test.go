package mp

import (
	"context"
	"testing"

	"scotty/wechat/mp/message"
)

func TestMp_SendCustomerMsg(t *testing.T) {
	ctx := context.TODO()
	mpClient, err := New()
	if err != nil {
		t.Fatal("new failed, err is ", err)
	}
	accessToken := ""
	toUser := ""
	text := "你好啊"
	param := message.NewCustomerTextMessage(toUser, text)
	err = mpClient.SendCustomerMsg(ctx, accessToken, param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
