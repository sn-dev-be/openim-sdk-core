package club

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/common"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/db_interface"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/model_struct"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/syncer"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
)

func NewClub(
	loginUserID string,
	db db_interface.DataBase,
	conversationCh chan common.Cmd2Value) *Club {
	c := &Club{
		loginUserID:    loginUserID,
		db:             db,
		conversationCh: conversationCh,
	}
	// c.initSyncer()
	return c
}

type Club struct {
	loginUserID string
	loginTime   int64
	db          db_interface.DataBase

	conversationCh chan common.Cmd2Value
	heartbeatCmdCh chan common.Cmd2Value

	listener           open_im_sdk_callback.OnClubListener
	listenerForService open_im_sdk_callback.OnListenerForService

	serverRequestSyncer      *syncer.Syncer[*model_struct.LocalServerRequest, [2]string]
	serverAdminRequestSyncer *syncer.Syncer[*model_struct.LocalAdminServerRequest, [2]string]
}

func (c *Club) initSyncer() {
	c.serverRequestSyncer = syncer.New(func(ctx context.Context, value *model_struct.LocalServerRequest) error {
		return c.db.InsertServerRequest(ctx, value)
	}, func(ctx context.Context, value *model_struct.LocalServerRequest) error {
		return c.db.DeleteServerRequest(ctx, value.ServerID, value.UserID)
	}, func(ctx context.Context, server, local *model_struct.LocalServerRequest) error {
		return c.db.UpdateServerRequest(ctx, server)
	}, func(value *model_struct.LocalServerRequest) [2]string {
		return [...]string{value.ServerID, value.UserID}
	}, nil, func(ctx context.Context, state int, server, local *model_struct.LocalServerRequest) error {
		switch state {
		case syncer.Insert:
			c.listener.OnClubApplicationAdded(utils.StructToJsonString(server))
		case syncer.Update:
			switch server.HandleResult {
			case constant.FriendResponseAgree:
				c.listener.OnClubApplicationAccepted(utils.StructToJsonString(server))
			case constant.FriendResponseRefuse:
				c.listener.OnClubApplicationRejected(utils.StructToJsonString(server))
			default:
				c.listener.OnClubApplicationAdded(utils.StructToJsonString(server))
			}
		}
		return nil
	})

	c.serverAdminRequestSyncer = syncer.New(func(ctx context.Context, value *model_struct.LocalAdminServerRequest) error {
		return c.db.InsertAdminServerRequest(ctx, value)
	}, func(ctx context.Context, value *model_struct.LocalAdminServerRequest) error {
		return c.db.DeleteAdminServerRequest(ctx, value.ServerID, value.UserID)
	}, func(ctx context.Context, server, local *model_struct.LocalAdminServerRequest) error {
		return c.db.UpdateAdminServerRequest(ctx, server)
	}, func(value *model_struct.LocalAdminServerRequest) [2]string {
		return [...]string{value.ServerID, value.UserID}
	}, nil, func(ctx context.Context, state int, server, local *model_struct.LocalAdminServerRequest) error {
		switch state {
		case syncer.Insert:
			c.listener.OnClubApplicationAdded(utils.StructToJsonString(server))
		case syncer.Update:
			switch server.HandleResult {
			case constant.FriendResponseAgree:
				c.listener.OnClubApplicationAccepted(utils.StructToJsonString(server))
			case constant.FriendResponseRefuse:
				c.listener.OnClubApplicationRejected(utils.StructToJsonString(server))
			default:
				c.listener.OnClubApplicationAdded(utils.StructToJsonString(server))
			}
		}
		return nil
	})

}

func (c *Club) SetClubListener(callback open_im_sdk_callback.OnClubListener) {
	if callback == nil {
		return
	}
	c.listener = callback
}

func (c *Club) SetListenerForService(listener open_im_sdk_callback.OnListenerForService) {
	c.listenerForService = listener
}
