package club

import (
	"context"

	"github.com/OpenIMSDK/protocol/club"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/utils"
	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
)

func (c *Club) SyncAllSelfServerApplication(ctx context.Context) error {
	list, err := c.GetServerSelfServerApplication(ctx)
	if err != nil {
		return err
	}
	localData, err := c.db.GetSendServerApplication(ctx)
	if err != nil {
		return err
	}
	if err := c.serverRequestSyncer.Sync(ctx, util.Batch(ServerServerRequestToLocalServerRequest, list), localData, nil); err != nil {
		return err
	}
	return nil
}

func (c *Club) SyncSelfServerApplications(ctx context.Context, serverIDs ...string) error {
	return c.SyncAllSelfServerApplication(ctx)
}

func (c *Club) SyncAllAdminServerApplication(ctx context.Context) error {
	requests, err := c.GetServerAdminServerApplicationList(ctx)
	if err != nil {
		return err
	}
	localData, err := c.db.GetAdminServerApplication(ctx)
	if err != nil {
		return err
	}
	return c.serverAdminRequestSyncer.Sync(ctx, util.Batch(ServerServerRequestToLocalAdminServerRequest, requests), localData, nil)
}

func (c *Club) SyncAdminServerApplications(ctx context.Context, serverIDs ...string) error {
	return c.SyncAllAdminServerApplication(ctx)
}

func (c *Club) GetServerSelfServerApplication(ctx context.Context) ([]*sdkws.ServerRequest, error) {
	fn := func(resp *club.GetServerApplicationListResp) []*sdkws.ServerRequest { return resp.ServerRequests }
	req := &club.GetUserReqApplicationListReq{UserID: c.loginUserID, Pagination: &sdkws.RequestPagination{}}
	return util.GetPageAll(ctx, constant.GetSendServerApplicationListRouter, req, fn)
}

func (c *Club) GetServerAdminServerApplicationList(ctx context.Context) ([]*sdkws.ServerRequest, error) {
	fn := func(resp *club.GetServerApplicationListResp) []*sdkws.ServerRequest { return resp.ServerRequests }
	req := &club.GetServerApplicationListReq{FromUserID: c.loginUserID, Pagination: &sdkws.RequestPagination{}}
	return util.GetPageAll(ctx, constant.GetRecvServerApplicationListRouter, req, fn)
}

func (c *Club) SyncServer(ctx context.Context, serverIDs []string) error {
	list, err := c.getServersInfoFromSvr(ctx, serverIDs)
	if err != nil {
		return err
	}
	localData, err := c.db.GetServers(ctx, serverIDs)
	if err != nil {
		return err
	}
	if err := c.serverSyncer.Sync(ctx, util.Batch(ServerToLocalServer, list), localData, nil); err != nil {
		return err
	}
	return c.SyncGroupCategoryByServer(ctx, serverIDs)
}

func (c *Club) SyncGroupCategoryByID(ctx context.Context, categoryIDs []string) error {
	list, err := c.getGroupCategoriesFromSvr(ctx, categoryIDs)
	if err != nil {
		return err
	}
	localData, err := c.db.GetGroupCategories(ctx, categoryIDs)
	if err != nil {
		return err
	}
	if err := c.groupCategorySyncer.Sync(ctx, util.Batch(ServerGroupCategoryToLocalGroupCategory, list), localData, nil); err != nil {
		return err
	}
	return nil
}

func (c *Club) SyncGroupCategoryByServer(ctx context.Context, serverIDs []string) error {
	list, err := c.getGroupCategoriesFromSvrByServer(ctx, serverIDs)
	if err != nil {
		return err
	}
	localData, err := c.db.GetGroupCategoriesByServer(ctx, serverIDs)
	if err != nil {
		return err
	}
	if err := c.groupCategorySyncer.Sync(ctx, util.Batch(ServerGroupCategoryToLocalGroupCategory, list), localData, nil); err != nil {
		return err
	}
	return nil
}

func (c *Club) SyncJoinedServerList(ctx context.Context) error {
	list, err := c.getJoinedServerList(ctx)
	if err != nil {
		return err
	}
	localData, err := c.db.GetAllServers(ctx)
	if err != nil {
		return err
	}
	if err := c.serverSyncer.Sync(ctx, util.Batch(ServerToLocalServer, list), localData, nil); err != nil {
		return err
	}

	serverIDs := utils.Slice(list, func(e *sdkws.ServerInfo) string { return e.ServerID })
	c.SyncGroupCategoryByServer(ctx, serverIDs)
	return nil
}
