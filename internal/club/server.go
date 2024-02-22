package club

import (
	"context"

	"github.com/OpenIMSDK/protocol/club"
	"github.com/OpenIMSDK/tools/utils"
	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/model_struct"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
)

func (c *Club) CreateServer(ctx context.Context, req *club.CreateServerReq) (string, error) {
	if req.OwnerUserID == "" {
		req.OwnerUserID = c.loginUserID
	}
	if req.Icon == "" {
		return "", sdkerrs.ErrArgs
	}
	if req.ServerName == "" {
		return "", sdkerrs.ErrArgs
	}

	resp, err := util.CallApi[club.CreateServerResp](ctx, constant.CreateServerRouter, req)
	if err != nil {
		return "", err
	}
	return resp.ServerID, nil
}

func (c *Club) GetSpecifiedServersInfo(ctx context.Context, serverIDs []string) ([]*model_struct.LocalServer, error) {
	serverList, err := c.db.GetServers(ctx, serverIDs)
	if err != nil {
		return nil, err
	}

	dbServerIDs := utils.Slice(serverList, func(s *model_struct.LocalServer) string { return s.ServerID })
	missServerIDs := utils.SliceSub(serverIDs, dbServerIDs)
	if len(missServerIDs) > 0 {
		c.SyncServer(ctx, missServerIDs, false)
		syncServerList, err := c.db.GetServers(ctx, missServerIDs)
		if err != nil {
			return nil, err
		}
		serverList = append(serverList, syncServerList...)
	}
	return serverList, nil
}

func (c *Club) GetJoinedServersInfo(ctx context.Context) ([]*model_struct.LocalServer, error) {
	serverList, err := c.db.GetAllServers(ctx)
	if err != nil {
		return nil, err
	}
	return serverList, nil
}

func (c *Club) dismissServer(ctx context.Context, serverID string) error {
	c.db.DeleteServer(ctx, serverID)
	c.db.DeleteServerMemberByServerIDAndUserID(ctx, serverID, c.loginUserID)
	c.db.DeleteGroupCategoryByServers(ctx, []string{serverID})
	c.db.DeleteGroupByServers(ctx, []string{serverID})

	//todo listener
	return nil
}

func (c *Club) dismissServerGroup(ctx context.Context, serverID, groupID string) error {
	err := c.db.DeleteGroup(ctx, groupID)
	if err != nil {
		return err
	}

	// totalUnreadCount, err := c.db.GetServerTotalUnreadCountByServerID(ctx, serverID)
	// if err != nil {
	// 	return err
	// }
	// c.conversationListener.OnServerUnreadMessageCountChanged(serverID, totalUnreadCount)
	//

	c.listener.OnServerGroupDismissed(groupID)
	return nil
}
