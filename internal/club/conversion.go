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
	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/model_struct"

	"github.com/OpenIMSDK/protocol/sdkws"
)

func ServerServerRequestToLocalServerRequest(info *sdkws.ServerRequest) *model_struct.LocalServerRequest {
	return &model_struct.LocalServerRequest{
		ServerID:    info.ServerInfo.ServerID,
		ServerName:  info.ServerInfo.ServerName,
		Icon:        info.ServerInfo.Icon,
		Description: info.ServerInfo.Description,
		Banner:      info.ServerInfo.Banner,
		CreateTime:  info.ServerInfo.CreateTime,
		Status:      info.ServerInfo.Status,
		OwnerUserID: info.ServerInfo.OwnerUserID,
		MemberNum:   int32(info.ServerInfo.MemberNumber),
		UserID:      info.UserInfo.UserID,
		Nickname:    info.UserInfo.Nickname,
		UserFaceURL: info.UserInfo.FaceURL,
		//Gender:        info.UserInfo.Gender,
		HandleResult: info.HandleResult,
		ReqMsg:       info.ReqMsg,
		HandledMsg:   info.HandleMsg,
		ReqTime:      info.ReqTime,
		HandleUserID: info.HandleUserID,
		HandledTime:  info.HandleTime,
		Ex:           info.Ex,
		//AttachedInfo:  info.AttachedInfo,
		JoinSource:    info.JoinSource,
		InviterUserID: info.InviterUserID,
	}
}

func ServerServerRequestToLocalAdminServerRequest(info *sdkws.ServerRequest) *model_struct.LocalAdminServerRequest {
	return &model_struct.LocalAdminServerRequest{
		LocalServerRequest: *ServerServerRequestToLocalServerRequest(info),
	}
}
