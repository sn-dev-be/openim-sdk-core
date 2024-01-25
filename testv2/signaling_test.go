package testv2

import (
	"context"
	"testing"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

func Test_SignalingInvite(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	inviteUserIDs := []string{"6512ba81b6f8463c81a0ab6a"}
	channelID, err := open_im_sdk.UserForSDK.Signaling().SignalingInvite(ctx, "sg_10010", inviteUserIDs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(channelID)
}

func Test_SignalingAccept(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingAccept(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingReject(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingReject(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingJoin(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingJoin(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingCancel(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingCancel(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "", "10010")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SignalingHungUp(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	err := open_im_sdk.UserForSDK.Signaling().SignalingHungUp(ctx, "94f3d1d5f39fc49b3496cdb379d0d3f3", "")
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
