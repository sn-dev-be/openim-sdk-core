package club

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/common"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/db_interface"
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
	c.initSyncer()
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
}

func (c *Club) initSyncer() {
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
