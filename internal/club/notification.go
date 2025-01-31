// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package club

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/common"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
)

func (c *Club) DoNotification(ctx context.Context, msg *sdkws.MsgData) {
	go func() {
		if err := c.doNotification(ctx, msg); err != nil {
			log.ZError(ctx, "DoGroupNotification failed", err)
		}
	}()
}

func (c *Club) doNotification(ctx context.Context, msg *sdkws.MsgData) error {

	if c.listener == nil {
		return errors.New("listener is nil")
	}
	switch msg.ContentType {
	case constant.ServerCreatedNotification:
		var detail sdkws.ServerCreatedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServer(ctx, []string{detail.Server.ServerID}, false); err != nil {
			return err
		}
		return c.group.SyncAllJoinedGroupsAndMembers(ctx)
	case constant.ServerDismissedNotification:
		var detail sdkws.ServerDissmissedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		c.listener.OnServerDismissed(detail.ServerID)
		return c.dismissServer(ctx, detail.ServerID)

	case constant.ServerInfoSetNotification:
		var detail sdkws.ServerInfoSetTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServer(ctx, []string{detail.Server.ServerID}, true); err != nil {
			return err
		}
		//todo listener
		return nil
	case constant.JoinServerApplicationNotification:

	case constant.ServerApplicationAcceptedNotification:
		var detail sdkws.ServerApplicationAcceptedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServer(ctx, []string{detail.Server.ServerID}, false); err != nil {
			return err
		}
		return c.group.SyncAllJoinedGroupsAndMembers(ctx)
	case constant.ServerApplicationRejectedNotification:

	case constant.ServerMemberEnterNotification:
		var detail sdkws.ServerMemberEnterTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if detail.User.UserID == c.loginUserID {
			if err := c.SyncServer(ctx, []string{detail.ServerID}, false); err != nil {
				return err
			}
			return c.group.SyncAllJoinedGroupsAndMembers(ctx)
		}
	case constant.ServerMemberQuitNotification:
		var detail sdkws.ServerMemberQuitTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.dismissServer(ctx, detail.ServerID); err != nil {
			return err
		}
		c.listener.OnServerDismissed(detail.ServerID)
		return nil
	case constant.ServerMemberKickedNotification:
		var detail sdkws.ServerMemberKickedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}

		for _, kickedUserID := range detail.MemberUserIDList {
			if c.loginUserID == kickedUserID {
				if err := c.dismissServer(ctx, detail.ServerID); err != nil {
					return err
				}
				c.listener.OnServerMemberKicked(detail.ServerID)
			}
		}
		return nil
	case constant.ServerMemberInfoSetNotification:
		var detail sdkws.ServerMemberInfoSetTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServerMemberByServer(ctx, []string{detail.ServerID}); err != nil {
			return err
		}
		return nil
	case constant.ServerMemberMutedNotification:
		var detail sdkws.ServerMemberMutedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServerMemberByServer(ctx, []string{detail.ServerID}); err != nil {
			return err
		}
		return nil
	case constant.ServerMemberCancelMutedNotification:
		var detail sdkws.ServerMemberCancelMutedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		if err := c.SyncServerMemberByServer(ctx, []string{detail.ServerID}); err != nil {
			return err
		}
		return nil
	case constant.ServerGroupCreatedNotification:
		var detail sdkws.ServerGroupCreatedTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}

		return c.group.SyncGroups(ctx, detail.Group.GroupID)
	case constant.ServerGroupDismissNotification:
		var detail sdkws.ServerGroupDismissTips
		if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
			return err
		}
		group, err := c.db.GetGroupInfoByGroupID(ctx, detail.GroupID)
		if err != nil {
			return err
		}
		data, err := json.Marshal(group)
		if err != nil {
			return err
		}
		c.dismissServerGroup(ctx, detail.ServerID, detail.GroupID)
		c.listener.OnServerGroupDismissed(string(data))
		common.TriggerCmdUpdateConversation(ctx, common.UpdateConNode{Action: constant.ServerTotalUnreadMessageChanged, Args: group.ServerID}, c.conversationCh)
		return nil

		//return c.group.SyncGroups(ctx, detail.Group.GroupID)
	// case constant.GroupOwnerTransferredNotification: // 1507
	// 	var detail sdkws.GroupOwnerTransferredTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	if err := g.SyncGroups(ctx, detail.Group.GroupID); err != nil {
	// 		return err
	// 	}
	// 	if detail.Group == nil {
	// 		return errors.New(fmt.Sprintf("group is nil, groupID: %s", detail.Group.GroupID))
	// 	}
	// 	return g.SyncAllGroupMember(ctx, detail.Group.GroupID)
	// case constant.MemberKickedNotification: // 1508
	// 	var detail sdkws.MemberKickedTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	var self bool
	// 	for _, info := range detail.KickedUserList {
	// 		if info.UserID == g.loginUserID {
	// 			self = true
	// 			break
	// 		}
	// 	}
	// 	if self {
	// 		members, err := g.db.GetGroupMemberListSplit(ctx, detail.Group.GroupID, 0, 0, 999999)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if err := g.db.DeleteGroupAllMembers(ctx, detail.Group.GroupID); err != nil {
	// 			return err
	// 		}
	// 		for _, member := range members {
	// 			data, err := json.Marshal(member)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			g.listener.OnGroupMemberDeleted(string(data))
	// 		}
	// 		group, err := g.db.GetGroupInfoByGroupID(ctx, detail.Group.GroupID)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		group.MemberCount = 0
	// 		data, err := json.Marshal(group)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if err := g.db.DeleteGroup(ctx, detail.Group.GroupID); err != nil {
	// 			return err
	// 		}
	// 		g.listener.OnGroupInfoChanged(string(data))
	// 		return nil
	// 	} else {
	// 		var userIDs []string
	// 		for _, info := range detail.KickedUserList {
	// 			userIDs = append(userIDs, info.UserID)
	// 		}
	// 		return g.SyncGroupMembers(ctx, detail.Group.GroupID, userIDs...)
	// 	}
	// case constant.MemberQuitNotification: // 1504
	// 	var detail sdkws.MemberQuitTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	if detail.QuitUser.UserID == g.loginUserID {
	// 		members, err := g.db.GetGroupMemberListSplit(ctx, detail.Group.GroupID, 0, 0, 999999)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if err := g.db.DeleteGroupAllMembers(ctx, detail.Group.GroupID); err != nil {
	// 			return err
	// 		}
	// 		for _, member := range members {
	// 			data, err := json.Marshal(member)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			g.listener.OnGroupMemberDeleted(string(data))
	// 		}
	// 		group, err := g.db.GetGroupInfoByGroupID(ctx, detail.Group.GroupID)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		group.MemberCount = 0
	// 		data, err := json.Marshal(group)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if err := g.db.DeleteGroup(ctx, detail.Group.GroupID); err != nil {
	// 			return err
	// 		}
	// 		g.listener.OnGroupInfoChanged(string(data))
	// 		return nil
	// 	} else {
	// 		return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.QuitUser.UserID)
	// 	}
	// case constant.MemberInvitedNotification: // 1509
	// 	var detail sdkws.MemberInvitedTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	if err := g.SyncGroups(ctx, detail.Group.GroupID); err != nil {
	// 		return err
	// 	}
	// 	var userIDs []string
	// 	for _, info := range detail.InvitedUserList {
	// 		userIDs = append(userIDs, info.UserID)
	// 	}
	//
	// 	if utils.IsContain(g.loginUserID, userIDs) {
	// 		return g.SyncAllGroupMember(ctx, detail.Group.GroupID)
	// 	} else {
	// 		return g.SyncGroupMembers(ctx, detail.Group.GroupID, userIDs...)
	// 	}
	// case constant.MemberEnterNotification: // 1510
	// 	var detail sdkws.MemberEnterTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	if detail.EntrantUser.UserID == g.loginUserID {
	// 		if err := g.SyncGroups(ctx, detail.Group.GroupID); err != nil {
	// 			return err
	// 		}
	// 		return g.SyncAllGroupMember(ctx, detail.Group.GroupID)
	// 	} else {
	// 		return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.EntrantUser.UserID)
	// 	}
	// case constant.GroupDismissedNotification: // 1511
	// 	var detail sdkws.GroupDismissedTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	g.listener.OnGroupDismissed(utils.StructToJsonString(detail.Group))
	// 	if err := g.db.DeleteGroupAllMembers(ctx, detail.Group.GroupID); err != nil {
	// 		return err
	// 	}
	// 	if err := g.db.DeleteGroup(ctx, detail.Group.GroupID); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncAllGroupMember(ctx, detail.Group.GroupID)
	// case constant.GroupMemberMutedNotification: // 1512
	// 	var detail sdkws.GroupMemberMutedTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.MutedUser.UserID)
	// case constant.GroupMemberCancelMutedNotification: // 1513
	// 	var detail sdkws.GroupMemberCancelMutedTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.MutedUser.UserID)
	// case constant.GroupMutedNotification: // 1514
	// 	return g.SyncGroups(ctx, msg.GroupID)
	// case constant.GroupCancelMutedNotification: // 1515
	// 	return g.SyncGroups(ctx, msg.GroupID)
	// case constant.GroupMemberInfoSetNotification: // 1516
	// 	var detail sdkws.GroupMemberInfoSetTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	//
	// 	return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.ChangedUser.UserID) //detail.ChangedUser.UserID
	// case constant.GroupMemberSetToAdminNotification: // 1517
	// 	var detail sdkws.GroupMemberInfoSetTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.ChangedUser.UserID)
	// case constant.GroupMemberSetToOrdinaryUserNotification: // 1518
	// 	var detail sdkws.GroupMemberInfoSetTips
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroupMembers(ctx, detail.Group.GroupID, detail.ChangedUser.UserID)
	// case constant.GroupInfoSetAnnouncementNotification: // 1519
	// 	var detail sdkws.GroupInfoSetAnnouncementTips //
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroups(ctx, detail.Group.GroupID)
	// case constant.GroupInfoSetNameNotification: // 1520
	// 	var detail sdkws.GroupInfoSetNameTips //
	// 	if err := utils.UnmarshalNotificationElem(msg.Content, &detail); err != nil {
	// 		return err
	// 	}
	// 	return g.SyncGroups(ctx, detail.Group.GroupID)
	default:
		return fmt.Errorf("unknown tips type: %d", msg.ContentType)
	}
	return nil
}
