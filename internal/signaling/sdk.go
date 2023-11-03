package signaling

import (
	"context"

	"github.com/OpenIMSDK/protocol/constant"
	"github.com/OpenIMSDK/protocol/sdkws"

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

func (s *Signaling) SignalingInvite(ctx context.Context, userIDList []string) error {
	req := sdkws.SignalReq{
		FromUserID:    s.loginUserID,
		ChannelID:     utils.GetMsgID(s.loginUserID),
		InviteUsersID: userIDList,
		SessionType:   constant.SingleChatType,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingInviation)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingInviteInGroup(ctx context.Context, userIDList []string, groupID string) error {
	req := sdkws.SignalReq{
		FromUserID:    s.loginUserID,
		ChannelID:     utils.GetMsgID(s.loginUserID),
		InviteUsersID: userIDList,
		SessionType:   constant.GroupChatType,
		GroupID:       groupID,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingInviation)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingAccept(ctx context.Context, channelID string) error {
	req := sdkws.SignalReq{
		FromUserID: s.loginUserID,
		ChannelID:  channelID,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingAccept)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingReject(ctx context.Context, channelID string, sessionType int32) error {
	req := sdkws.SignalReq{
		FromUserID:  s.loginUserID,
		ChannelID:   channelID,
		SessionType: sessionType,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingReject)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingJoin(ctx context.Context, channelID string, groupID string) error {
	req := sdkws.SignalReq{
		FromUserID:  s.loginUserID,
		ChannelID:   channelID,
		SessionType: constant.SuperGroupChatType,
		GroupID:     groupID,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingJoin)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingCancel(ctx context.Context, channelID string, sessionType int32, cancelUserID string) error {
	req := sdkws.SignalReq{
		FromUserID:    s.loginUserID,
		ChannelID:     channelID,
		SessionType:   sessionType,
		InviteUsersID: []string{cancelUserID},
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingCancel)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingHungUp(ctx context.Context, channelID string, sessionType int32) error {
	req := sdkws.SignalReq{
		FromUserID:  s.loginUserID,
		ChannelID:   channelID,
		SessionType: sessionType,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingHungUp)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SignalingClose(ctx context.Context, channelID string, sessionType int32) error {
	req := sdkws.SignalReq{
		FromUserID:  s.loginUserID,
		ChannelID:   channelID,
		SessionType: sessionType,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingClose)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) MichoneStatusChange(ctx context.Context, channelID string, status int32) error {
	req := sdkws.SignalReq{
		FromUserID:     s.loginUserID,
		ChannelID:      channelID,
		MicphoneStatus: status,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingMicphoneStatusChange)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}

func (s *Signaling) SpeakStatusChange(ctx context.Context, channelID string) error {
	req := sdkws.SignalReq{
		FromUserID: s.loginUserID,
		ChannelID:  channelID,
	}
	m := sdk_struct.SignalingStruct{}
	err := s.initBasicInfo(ctx, &m, constant.VoiceCall, constant.SignalingSpeakStatusChange)
	if err != nil {
		return err
	}
	m.SignalReq = &req
	_, err = s.SendSignalMessage(ctx, &m)
	return err
}
