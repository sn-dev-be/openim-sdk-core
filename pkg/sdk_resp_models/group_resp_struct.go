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

package sdk_resp_models

type GroupInfoResp struct {
	GroupID                string `json:"groupID"`
	GroupName              string `json:"groupName"`
	Notification           string `json:"notification"`
	Introduction           string `json:"introduction"`
	FaceURL                string `json:"faceURL"`
	CreateTime             int64  `json:"createTime"`
	Status                 int32  `json:"status"`
	CreatorUserID          string `json:"creatorUserID"`
	GroupType              int32  `json:"groupType"`
	OwnerUserID            string `json:"ownerUserID"`
	MemberCount            int32  `json:"memberCount"`
	Ex                     string `json:"ex"`
	AttachedInfo           string `json:"attachedInfo"`
	NeedVerification       int32  `json:"needVerification"`
	LookMemberInfo         int32  `json:"lookMemberInfo"`
	ApplyMemberFriend      int32  `json:"applyMemberFriend"`
	NotificationUpdateTime int64  `json:"notificationUpdateTime"`
	NotificationUserID     string `json:"notificationUserID"`
	Condition              string `gorm:"column:condition"`
	ConditionType          int32  `gorm:"column:condition_type"`
	SyncMode               int32  `gorm:"column:sync_mode;default:0"`
	VisitorMode            int32  `gorm:"column:visitor_mode;default:0"`
	ViewMode               int32  `gorm:"column:view_mode;default:0"`
	GroupMode              int32  `gorm:"column:group_mode;default:1"`
	GroupCategoryID        string `gorm:"column:group_category_id;index;size:65"`
	ServerID               string `gorm:"column:server_id;index;size:64"`
	ReorderWeight          int32  `gorm:"column:reorder_weight;default:0"`
	Saved                  int    `json:"saved"` // 新字段 saved
}
