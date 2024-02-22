package club

import (
	"context"

	"github.com/OpenIMSDK/protocol/club"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/openimsdk/openim-sdk-core/v3/internal/group"
	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
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
	conversationCh chan common.Cmd2Value,
	group *group.Group) *Club {
	c := &Club{
		loginUserID:    loginUserID,
		db:             db,
		conversationCh: conversationCh,
		group:          group,
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

	serverRequestSyncer      *syncer.Syncer[*model_struct.LocalServerRequest, [2]string]
	serverAdminRequestSyncer *syncer.Syncer[*model_struct.LocalAdminServerRequest, [2]string]
	serverSyncer             *syncer.Syncer[*model_struct.LocalServer, string]
	groupCategorySyncer      *syncer.Syncer[*model_struct.LocalGroupCategory, string]
	serverMemberSyncer       *syncer.Syncer[*model_struct.LocalServerMember, [2]string]

	group *group.Group
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
			c.listener.OnServerApplicationAdded(utils.StructToJsonString(server))
		case syncer.Update:
			switch server.HandleResult {
			case constant.FriendResponseAgree:
				c.listener.OnServerApplicationAccepted(utils.StructToJsonString(server))
			case constant.FriendResponseRefuse:
				c.listener.OnServerApplicationRejected(utils.StructToJsonString(server))
			default:
				c.listener.OnServerApplicationAdded(utils.StructToJsonString(server))
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
			c.listener.OnServerApplicationAdded(utils.StructToJsonString(server))
		case syncer.Update:
			switch server.HandleResult {
			case constant.FriendResponseAgree:
				c.listener.OnServerApplicationAccepted(utils.StructToJsonString(server))
			case constant.FriendResponseRefuse:
				c.listener.OnServerApplicationRejected(utils.StructToJsonString(server))
			default:
				c.listener.OnServerApplicationAdded(utils.StructToJsonString(server))
			}
		}
		return nil
	})

	c.serverSyncer = syncer.New(func(ctx context.Context, value *model_struct.LocalServer) error {
		return c.db.InsertServer(ctx, value)
	}, func(ctx context.Context, value *model_struct.LocalServer) error {
		return c.db.DeleteServer(ctx, value.ServerID)
	}, func(ctx context.Context, server, local *model_struct.LocalServer) error {
		return c.db.UpdateServer(ctx, server)
	}, func(value *model_struct.LocalServer) string {
		return value.ServerID
	}, nil, nil)

	c.groupCategorySyncer = syncer.New(func(ctx context.Context, value *model_struct.LocalGroupCategory) error {
		return c.db.InsertGroupCategory(ctx, value)
	}, func(ctx context.Context, value *model_struct.LocalGroupCategory) error {
		return c.db.DeleteGroupCategory(ctx, value.ServerID)
	}, func(ctx context.Context, server, local *model_struct.LocalGroupCategory) error {
		return c.db.UpdateGroupCategory(ctx, server)
	}, func(value *model_struct.LocalGroupCategory) string {
		return value.CategoryID
	}, nil, nil)

	c.serverMemberSyncer = syncer.New(func(ctx context.Context, value *model_struct.LocalServerMember) error {
		return c.db.InsertServerMember(ctx, value)
	}, func(ctx context.Context, value *model_struct.LocalServerMember) error {
		return c.db.DeleteServerMemberByServerIDAndUserID(ctx, value.ServerID, value.UserID)
	}, func(ctx context.Context, server, local *model_struct.LocalServerMember) error {
		return c.db.UpdateServerMember(ctx, server)
	}, func(value *model_struct.LocalServerMember) [2]string {
		return [...]string{value.ServerID, value.UserID}
	}, nil, nil)

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

func (c *Club) getServersInfoFromSvr(ctx context.Context, serverIDs []string) ([]*sdkws.ServerInfo, error) {
	resp, err := util.CallApi[club.GetServersInfoResp](ctx, constant.GetServerInfoListRouter, &club.GetServersInfoReq{ServerIDs: serverIDs})
	if err != nil {
		return nil, err
	}
	servers := []*sdkws.ServerInfo{}
	for _, serverResp := range resp.Servers {
		servers = append(servers, serverResp.Server)
	}
	return servers, nil
}

func (c *Club) getGroupCategoriesFromSvr(ctx context.Context, categoryIDs []string) ([]*sdkws.GroupCategoryInfo, error) {
	resp, err := util.CallApi[club.GetGroupCategoriesResp](ctx, constant.GetGroupCategoryListRouter, &club.GetGroupCategoriesReq{CategoryIDs: categoryIDs})
	if err != nil {
		return nil, err
	}
	return resp.GroupCategories, nil
}

func (c *Club) getGroupCategoriesFromSvrByServer(ctx context.Context, serverIDs []string) ([]*sdkws.GroupCategoryInfo, error) {
	resp, err := util.CallApi[club.GetServersInfoResp](ctx, constant.GetServerInfoListRouter, &club.GetServersInfoReq{ServerIDs: serverIDs})
	if err != nil {
		return nil, err
	}
	categories := []*sdkws.GroupCategoryInfo{}
	for _, serverResp := range resp.Servers {
		for _, categoryResp := range serverResp.CategoryList {
			categories = append(categories, categoryResp.CategoryInfo)
		}
	}
	return categories, nil
}

func (c *Club) getGroupsFromSvrByServer(ctx context.Context, serverIDs []string) ([]*sdkws.GroupCategoryInfo, error) {
	resp, err := util.CallApi[club.GetServersInfoResp](ctx, constant.GetServerInfoListRouter, &club.GetServersInfoReq{ServerIDs: serverIDs})
	if err != nil {
		return nil, err
	}
	categories := []*sdkws.GroupCategoryInfo{}
	for _, serverResp := range resp.Servers {
		for _, categoryResp := range serverResp.CategoryList {
			categories = append(categories, categoryResp.CategoryInfo)
		}
	}
	return categories, nil
}

func (c *Club) getJoinedServerList(ctx context.Context) ([]*sdkws.ServerInfo, error) {
	resp, err := util.CallApi[club.GetJoinedServerListResp](ctx, constant.GetJoinedServerListRouter, &club.GetJoinedServerListReq{FromUserID: c.loginUserID, Pagination: &sdkws.RequestPagination{PageNumber: 1, ShowNumber: 1000}})
	if err != nil {
		return nil, err
	}
	return resp.Servers, nil
}

func (c *Club) getServerMemberList(ctx context.Context, serverID string) ([]*sdkws.ServerMemberFullInfo, error) {
	req := &club.GetServerMembersInfoReq{UserIDs: []string{c.loginUserID}}
	if serverID != "" {
		req.ServerID = serverID
	}
	resp, err := util.CallApi[club.GetServerMembersInfoResp](ctx, constant.GetServerMembersInfoRouter, req)
	if err != nil {
		return nil, err
	}
	return resp.Members, nil
}
