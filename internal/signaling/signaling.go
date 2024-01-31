package signaling

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/internal/interaction"
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/ccontext"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/db_interface"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"

	"github.com/jinzhu/copier"
)

type Signaling struct {
	*interaction.LongConnMgr
	loginUserID string
	platformID  int32
	db          db_interface.DataBase
	listener    open_im_sdk_callback.OnSignalingListener
}

func NewSignaling(
	ctx context.Context,
	longConnMgr *interaction.LongConnMgr,
	db db_interface.DataBase,
	listener open_im_sdk_callback.OnSignalingListener,
) *Signaling {
	info := ccontext.Info(ctx)
	s := &Signaling{
		LongConnMgr: longConnMgr,
		loginUserID: info.UserID(),
		platformID:  info.PlatformID(),
		db:          db,
	}
	s.SetSignalingListener(listener)
	return s
}

func (s *Signaling) SetSignalingListener(listener open_im_sdk_callback.OnSignalingListener) {
	if s.listener != nil {
		return
	}
	s.listener = listener
}

func (s *Signaling) SendVoiceSignal(
	ctx context.Context,
	signalType int32,
	voiceElem *sdkws.SignalVoiceReq,
) (*sdk_struct.SignalingStruct, error) {
	m := sdk_struct.SignalingStruct{}
	if err := s.initBasicInfo(ctx, &m, constant.VoiceCall, signalType); err != nil {
		return nil, err
	}
	conversation, err := s.db.GetConversation(ctx, voiceElem.ConversationID)
	if err != nil {
		log.ZError(ctx, "SendSignalMessage GetConversation err", err)
		return nil, err
	}
	voiceElem.FromUserID = s.loginUserID
	voiceElem.SessionType = conversation.ConversationType
	voiceElem.ConversationID = conversation.ConversationID
	if conversation.ConversationType == constant.GroupChatType {
		voiceElem.GroupID = conversation.GroupID
	}
	m.Content = utils.StructToJsonString(voiceElem)
	log.ZInfo(ctx, "send voice signal", "signalType", signalType, "userID", s.loginUserID)
	return s.sendMessageToServer(ctx, &m)
}

func (s *Signaling) sendMessageToServer(
	ctx context.Context,
	m *sdk_struct.SignalingStruct,
) (*sdk_struct.SignalingStruct, error) {
	log.ZDebug(ctx, "send signal msg", "data", m.Content)
	var wsSignalData sdkws.SignalData
	copier.Copy(&wsSignalData, m)
	wsSignalData.Content = []byte(m.Content)
	wsSignalData.CreateTime = m.CreateTime
	m.Content = ""

	var sendSignalMsgResp sdkws.UserSendSignalMsgResp
	err := s.LongConnMgr.SendReqWaitResp(ctx, &wsSignalData, constant.SendSignalMsg, &sendSignalMsgResp)
	if err != nil {
		log.ZError(ctx, "send signal msg to server failed", err, "message", m)
		return m, err
	}
	return m, nil
}
