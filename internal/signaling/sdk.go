package signaling

import (
	"context"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/protocol/third"

	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
)

func (s *Signaling) initBasicInfo(
	ctx context.Context,
	m *sdk_struct.SignalingStruct,
	mediaType,
	signalType int32,
) error {
	m.CreateTime = utils.GetCurrentTimestampByMill()
	m.SendID = s.loginUserID
	m.SenderPlatformID = s.platformID
	m.MediaType = mediaType
	m.SignalType = signalType
	userInfo, err := s.db.GetLoginUser(ctx, s.loginUserID)
	if err != nil {
		return err
	}
	m.SenderFaceURL = userInfo.FaceURL
	m.SenderNickname = userInfo.Nickname
	return nil
}

func (s *Signaling) SignalingInvite(
	ctx context.Context,
	conversationID string,
	userIDList []string,
) (string, error) {
	channelID := utils.GetMsgID(conversationID)
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		InviteUsersID:  userIDList,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingInviation, &req)
	return channelID, err
}

func (s *Signaling) SignalingAccept(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingAccept, &req)
	return err
}

func (s *Signaling) SignalingReject(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingReject, &req)
	return err
}

func (s *Signaling) SignalingJoin(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingJoin, &req)
	return err
}

func (s *Signaling) SignalingCancel(
	ctx context.Context,
	conversationID string,
	channelID string,
	canncelUserID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
		InviteUsersID:  []string{canncelUserID},
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingCancel, &req)
	return err
}

func (s *Signaling) SignalingHungUp(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingHungUp, &req)
	return err
}

func (s *Signaling) SignalingClose(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingClose, &req)
	return err
}

func (s *Signaling) MichoneStatusChange(
	ctx context.Context,
	conversationID string,
	channelID string,
	status int32,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
		MicphoneStatus: status,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingMicphoneStatusChange, &req)
	return err
}

func (s *Signaling) SpeakStatusChange(
	ctx context.Context,
	conversationID string,
	channelID string,
) error {
	req := sdkws.SignalVoiceReq{
		ConversationID: conversationID,
		ChannelID:      channelID,
	}
	_, err := s.SendVoiceSignal(ctx, constant.SignalingSpeakStatusChange, &req)
	return err
}

func (s *Signaling) GetRtcToken(
	ctx context.Context,
	userID string,
	channelID string,
	roleType int32,
) (string, error) {
	req := third.InitiateRtcTokenReq{
		UserID:    userID,
		ChannelID: channelID,
		RoleType:  roleType,
	}
	resp, err := util.CallApi[third.InitiateRtcTokenResp](ctx, constant.GetRtcTokenRouter, &req)
	if err != nil {
		return "", err
	}
	return resp.Token, err
}
