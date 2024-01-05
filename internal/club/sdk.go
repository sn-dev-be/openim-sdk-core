// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package club

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"

	"github.com/OpenIMSDK/protocol/club"
)

// // deprecated use CreateGroup
// funcation (g *Group) CreateGroup(ctx context.Context, groupBaseInfo sdk_params_callback.CreateGroupBaseInfoParam, memberList sdk_params_callback.CreateGroupMemberRoleParam) (*sdkws.GroupInfo, error) {
//	req := &group.CreateGroupReq{
//		GroupInfo: &sdkws.GroupInfo{
//			GroupName:    groupBaseInfo.GroupName,
//			Notification: groupBaseInfo.Notification,
//			Introduction: groupBaseInfo.Introduction,
//			FaceURL:      groupBaseInfo.FaceURL,
//			Ex:           groupBaseInfo.Ex,
//			GroupType:    groupBaseInfo.GroupType,
//		},
//	}
//	if groupBaseInfo.NeedVerification != nil {
//		req.GroupInfo.NeedVerification = *groupBaseInfo.NeedVerification
//	}
//	for _, info := range memberList {
//		switch info.RoleLevel {
//		case constant.GroupOrdinaryUsers:
//			req.InitMembers = append(req.InitMembers, info.UserID)
//		case constant.GroupOwner:
//			req.OwnerUserID = info.UserID
//		case constant.GroupAdmin:
//			req.AdminUserIDs = append(req.AdminUserIDs, info.UserID)
//		default:
//			return nil, sdkerrs.ErrArgs.Wrap(fmt.Sprintf("CreateGroup: invalid role level %d", info.RoleLevel))
//		}
//	}
//	return g.CreateGroup(ctx, req)
// }

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

func (c *Club) JoinServer(ctx context.Context, req *club.JoinServerReq) error {
	if req.InviterUserID == "" {
		req.InviterUserID = c.loginUserID
	}
	if req.ServerID == "" {
		return sdkerrs.ErrArgs
	}

	_, err := util.CallApi[club.JoinServerResp](ctx, constant.JoinServerRouter, req)
	if err != nil {
		return err
	}
	return nil
}
