package testv2

import (
	"context"
	"testing"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

func Test_SendMessageInServerGroup(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, _ := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	_, err := open_im_sdk.UserForSDK.Conversation().SendMessage(ctx, msg, "", "sg1", nil)
	if err != nil {
		t.Fatal(err)
	}
}
