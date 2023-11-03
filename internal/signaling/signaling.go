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

func (s *Signaling) SendSignalMessage(
	ctx context.Context,
	m *sdk_struct.SignalingStruct,
) (*sdk_struct.SignalingStruct, error) {

	// switch m.SignalType {
	// case constant.SignalingInviation:
	// default:
	// 	return nil, sdkerrs.ErrMsgContentTypeNotSupport
	// }
	m.Content = utils.StructToJsonString(m.SignalReq)
	return s.sendMessageToServer(ctx, m)
}

func (s *Signaling) sendMessageToServer(
	ctx context.Context,
	m *sdk_struct.SignalingStruct,
) (*sdk_struct.SignalingStruct, error) {

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
