package testv2

import (
	"context"
	"testing"
	"time"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

func Test_SignalingInvite(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	channelID, err := open_im_sdk.UserForSDK.Signaling().SignalingInvite(ctx, "si_80010_90010", "90010", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(channelID)
	time.Sleep(time.Second * 5)
}

func Test_SignalingAccept(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingAccept(ctx, "si_80010_90010", "9ef9424a631146566958e4fd0eb459e5")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingReject(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	for i := 0; i < 1; i++ {
		err := open_im_sdk.UserForSDK.Signaling().SignalingReject(ctx, "si_80010_90010", "bc49a112933babb9b5501b0936ed1276")
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Second * 5)
}

func Test_SignalingJoin(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingJoin(ctx, "si_80010_90010", "94f3d1d5f39fc49b3496cdb379d0d3f3")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingCancel(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingCancel(ctx, "si_80010_90010", "97324170c11a67adc2f00fa03eda19a1", "90010")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 5)
}

func Test_SignalingHungUp(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingHungUp(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "", 1)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingClose(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingClose(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingMicphoneStatusChange(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().MichoneStatusChange(ctx, "51cc96934c80801b8bb4f3c94f395fc0", "", 1)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingSpeakStatusChange(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SpeakStatusChange(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
	if err != nil {
		t.Fatal(err)
	}
}
func Test_GetRtcToken(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	token, err := open_im_sdk.UserForSDK.Signaling().GetRtcToken(ctx, "1101", "test_channel", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
