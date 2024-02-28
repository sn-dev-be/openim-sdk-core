package club

import (
	"context"
	"fmt"
	"math"

	"github.com/OpenIMSDK/protocol/club"
	pbconstant "github.com/OpenIMSDK/protocol/constant"
	"github.com/OpenIMSDK/protocol/sdkws"

	"github.com/OpenIMSDK/tools/utils"
	localUtil "github.com/openimsdk/openim-sdk-core/v3/pkg/utils"

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

func (c *Club) DismissServer(ctx context.Context, req *club.DismissServerReq) error {
	if req.ServerID == "" {
		return sdkerrs.ErrArgs
	}

	_, err := util.CallApi[club.DismissServerResp](ctx, constant.DismissServerRouter, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Club) DeleteServerGroup(ctx context.Context, req *club.DeleteServerGroupReq) error {
	if req.ServerID == "" || req.GroupIDs == nil || len(req.GroupIDs) == 0 {
		return sdkerrs.ErrArgs
	}

	_, err := util.CallApi[club.DeleteServerGroupResp](ctx, constant.DeleteServerGroupRouter, req)
	if err != nil {
		return err
	}
	return nil
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

	return nil
}

func (c *Club) AcceptServerApplication(ctx context.Context, serverID, fromUserID, handleMsg, conversationID string) error {
	return c.HandlerServerApplication(ctx, &club.ServerApplicationResponseReq{ServerID: serverID, FromUserID: fromUserID, HandledMsg: handleMsg, HandleResult: constant.GroupResponseAgree, ConversationID: conversationID})
}

func (c *Club) RefuseServerApplication(ctx context.Context, serverID, fromUserID, handleMsg, conversationID string) error {
	return c.HandlerServerApplication(ctx, &club.ServerApplicationResponseReq{ServerID: serverID, FromUserID: fromUserID, HandledMsg: handleMsg, HandleResult: constant.GroupResponseRefuse, ConversationID: conversationID})
}

func (c *Club) HandlerServerApplication(ctx context.Context, req *club.ServerApplicationResponseReq) error {
	keywordList := []string{}
	keywordList = append(keywordList, "serverID\\\":\\\""+req.ServerID)
	keywordList = append(keywordList, "userID\\\":\\\""+req.FromUserID)
	//keywordList = append(keywordList, "handleResult\\\":\\\""+strconv.FormatInt(pbconstant.ServerResponseNotHandle, 10))
	sList, err := c.db.SearchMessageByContentTypeAndKeyword(ctx, []int{constant.JoinServerApplicationNotification}, req.ConversationID, keywordList, constant.KeywordMatchAnd, 0, math.MaxInt64)
	if err != nil {
		return err
	}

	seqs := []int64{}
	for _, msg := range sList {
		if msg.ContentType != constant.JoinServerApplicationNotification {
			continue
		}
		var detail sdkws.JoinServerApplicationTips
		if err := localUtil.UnmarshalNotificationElem([]byte(msg.Content), &detail); err != nil {
			return err
		}
		if detail.Server.ServerID != req.ServerID {
			continue
		}
		if detail.Applicant.UserID != req.FromUserID {
			continue
		}
		if detail.HandleResult != pbconstant.ServerResponseNotHandle {
			continue
		}
		seqs = append(seqs, msg.Seq)
	}
	req.Seqs = seqs

	fmt.Println(utils.StructToJsonString(req))
	if err := util.ApiPost(ctx, constant.ServerApplicationResponseRouter, req, nil); err != nil {
		return err
	}
	// SyncAdminGroupApplication todo
	return nil
}
