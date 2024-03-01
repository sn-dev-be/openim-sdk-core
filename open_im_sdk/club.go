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

package open_im_sdk

import "github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"

//funcation CreateGroup(callback open_im_sdk_callback.Base, operationID string, groupBaseInfo string, memberList string) {
//	call(callback, operationID, UserForSDK.Group().CreateGroup, groupBaseInfo, memberList)
//}

func CreateServer(callback open_im_sdk_callback.Base, operationID string, serverReqInfo string) {
	call(callback, operationID, UserForSDK.Club().CreateServer, serverReqInfo)
}

func DismissServer(callback open_im_sdk_callback.Base, operationID string, serverID string) {
	call(callback, operationID, UserForSDK.Club().DismissServer, serverID)
}

func DeleteServerGroup(callback open_im_sdk_callback.Base, operationID string, serverID string, groupIDs string) {
	call(callback, operationID, UserForSDK.Club().DeleteServerGroup, serverID, groupIDs)
}

func JoinServer(callback open_im_sdk_callback.Base, operationID string, joinServerReq string) {
	call(callback, operationID, UserForSDK.Club().JoinServer, joinServerReq)
}

func QuitServer(callback open_im_sdk_callback.Base, operationID string, quitServerReq string) {
	call(callback, operationID, UserForSDK.Club().QuitServer, quitServerReq)
}

func KickServerMember(callback open_im_sdk_callback.Base, operationID string, serverID, reason string, kickedUserIDs string) {
	call(callback, operationID, UserForSDK.Club().KickServerMember, serverID, reason, kickedUserIDs)
}

func SetServerMemberInfo(callback open_im_sdk_callback.Base, operationID string, serverMembers string) {
	call(callback, operationID, UserForSDK.Club().SetServerMemberInfo, serverMembers)
}

func GetSpecifiedServersInfo(callback open_im_sdk_callback.Base, operationID string, serverIDList string) {
	call(callback, operationID, UserForSDK.Club().GetSpecifiedServersInfo, serverIDList)
}

func AcceptServerApplication(callback open_im_sdk_callback.Base, operationID string, serverID string, fromUserID string, handleMsg string, conversationID string) {
	call(callback, operationID, UserForSDK.Club().AcceptServerApplication, serverID, fromUserID, handleMsg, conversationID)
}

func RefuseServerApplication(callback open_im_sdk_callback.Base, operationID string, serverID string, fromUserID string, handleMsg string, conversationID string) {
	call(callback, operationID, UserForSDK.Club().RefuseServerApplication, serverID, fromUserID, handleMsg, conversationID)
}
