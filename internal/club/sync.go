package club

import (
	"context"

	"github.com/OpenIMSDK/protocol/club"
	"github.com/OpenIMSDK/protocol/sdkws"
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
