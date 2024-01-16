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
	"errors"

	"github.com/OpenIMSDK/protocol/club"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/openimsdk/openim-sdk-core/v3/internal/util"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"gorm.io/gorm"
)

func (c *Club) JoinServer(ctx context.Context, req *club.JoinServerReq) error {
	if req.InviterUserID == "" {
		req.InviterUserID = c.loginUserID
	}
	if req.ServerID == "" {
		return sdkerrs.ErrArgs
	}

	_, err := c.db.GetServerMemberByServerIDAndUserID(ctx, req.ServerID, c.loginUserID)
	if err != nil {
		return err
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	_, err = util.CallApi[club.JoinServerResp](ctx, constant.JoinServerRouter, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Club) QuitServer(ctx context.Context, req *club.QuitServerReq) error {
	if req.UserID == "" {
		req.UserID = c.loginUserID
	}
	if req.ServerID == "" {
		return sdkerrs.ErrArgs
	}

	_, err := c.db.GetServerMemberByServerIDAndUserID(ctx, req.ServerID, c.loginUserID)
	if err != nil {
		return err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	_, err = util.CallApi[club.QuitServerResp](ctx, constant.QuitServerRouter, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Club) SetServerMemberInfo(ctx context.Context, req *club.SetServerMemberInfoReq) error {
	if req.Members == nil || len(req.Members) == 0 {
		return sdkerrs.ErrArgs
	}

	for _, member := range req.Members {
		member.UserID = c.loginUserID
	}

	_, err := util.CallApi[club.SetServerMemberInfoResp](ctx, constant.SetServerMemberInfoRouter, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Club) updateServerMember(ctx context.Context, serverMember *sdkws.ServerMemberFullInfo) error {
	if serverMember == nil {
		return sdkerrs.ErrArgs
	}
	localServerMember := ServerMemberToLocalServerMember(serverMember)

	err := c.db.UpdateServerMember(ctx, localServerMember)
	if err != nil {
		return err
	}
	//todo listener
	return nil
}
