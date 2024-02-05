package signaling

import (
	"context"
	"errors"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"
)

func (s *Signaling) DoNotification(ctx context.Context, msg *sdkws.MsgData) {
	go func() {
		if err := s.doNotification(ctx, msg); err != nil {
			log.ZError(ctx, "DoSignalingNotification failed", err)
		}
	}()
}

func (s *Signaling) doNotification(ctx context.Context, msg *sdkws.MsgData) error {
	if s.listener == nil {
		return errors.New("listener is nil")
	}
	tips := sdkws.SignalVoiceTips{}
	if err := utils.UnmarshalNotificationElem(msg.Content, &tips); err != nil {
		log.ZError(ctx, "comm.UnmarshalTips failed", err, "msg", msg.Content)
		return err
	}
	if msg.ContentType != constant.SignalingClosedNotification && (tips.OpUser == nil || tips.OpUser.UserID == s.loginUserID) {
		return nil
	}
	notificationMsg := utils.StructToJsonString(&tips)
	switch msg.ContentType {
	case constant.SignalingInvitedNotification:
		s.listener.OnReceiveNewInvitation(notificationMsg)
	case constant.SignalingAcceptedNotification:
		s.listener.OnInviteeAccepted(notificationMsg)
	case constant.SignalingRejectedNotification:
		s.listener.OnInviteeRejected(notificationMsg)
	case constant.SignalingJoinedNotification:
		s.listener.OnJoined(notificationMsg)
	case constant.SignalingCanceledNotification:
		s.listener.OnInvitationCancelled(notificationMsg)
	case constant.SignalingHungUpNotification:
		s.listener.OnHangUp(notificationMsg)
	case constant.SignalingClosedNotification:
		s.listener.OnClosed(notificationMsg)
	case constant.SignalingMicphoneStatusChangedNotification:
		s.listener.OnMicphoneStatusChanged(notificationMsg)
	case constant.SignalingSpeakStatusChangedNotification:
		s.listener.OnSpeakStatusChanged(notificationMsg)
	}
	return nil
}
