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

package constant

const (
	GetSelfUserInfoRouter         = "/user/get_self_user_info"
	GetUsersInfoRouter            = "/user/get_users_info"
	UpdateSelfUserInfoRouter      = "/user/update_user_info"
	SetUserSettingRouter          = "/user/set_user_setting"
	GetUsersInfoFromCacheRouter   = "/user/get_users_info_from_cache"
	AccountCheck                  = "/user/account_check"
	UserRegister                  = "/user/user_register"
	SubscribeUsersStatusRouter    = "/user/subscribe_users_status"
	GetSubscribeUsersStatusRouter = "/user/get_subscribe_users_status"
	GetUserStatusRouter           = "/user/get_users_status"
	GetUserSetttingRouter         = "/user/get_user_setting"
	SetUserSetttingRouter         = "/user/set_user_setting"

	AddFriendRouter                    = "/friend/add_friend"
	DeleteFriendRouter                 = "/friend/delete_friend"
	GetFriendApplicationListRouter     = "/friend/get_friend_apply_list"      //recv
	GetSelfFriendApplicationListRouter = "/friend/get_self_friend_apply_list" //send
	ImportFriendListRouter             = "/friend/import_friend"
	IsFriendRouter                     = "/friend/is_friend"

	GetDesignatedFriendsApplyRouter = "/friend/get_designated_friend_apply"
	GetFriendListRouter             = "/friend/get_friend_list"
	GetDesignatedFriendsRouter      = "/friend/get_designated_friends"
	AddFriendResponse               = "/friend/add_friend_response"
	SetFriendRemark                 = "/friend/set_friend_remark"

	AddBlackRouter     = "/friend/add_black"
	RemoveBlackRouter  = "/friend/remove_black"
	GetBlackListRouter = "/friend/get_black_list"

	PullUserMsgRouter      = "/chat/pull_msg"
	PullUserMsgBySeqRouter = "/chat/pull_msg_by_seq"
	NewestSeqRouter        = "/chat/newest_seq"

	// msg
	ClearConversationMsgRouter             = RouterMsg + "/clear_conversation_msg" // Clear the message of the specified conversation
	ClearAllMsgRouter                      = RouterMsg + "/user_clear_all_msg"     // Clear all messages of the current user
	DeleteMsgsRouter                       = RouterMsg + "/delete_msgs"            // Delete the specified message
	RevokeMsgRouter                        = RouterMsg + "/revoke_msg"
	SetMessageReactionExtensionsRouter     = RouterMsg + "/set_message_reaction_extensions"
	AddMessageReactionExtensionsRouter     = RouterMsg + "/add_message_reaction_extensions"
	MarkMsgsAsReadRouter                   = RouterMsg + "/mark_msgs_as_read"
	GetConversationsHasReadAndMaxSeqRouter = RouterMsg + "/get_conversations_has_read_and_max_seq"

	MarkConversationAsRead    = RouterMsg + "/mark_conversation_as_read"
	MarkMsgsAsRead            = RouterMsg + "/mark_msgs_as_read"
	SetConversationHasReadSeq = RouterMsg + "/set_conversation_has_read_seq"
	SendMsgRouter             = RouterMsg + "/send_msg"
	GetServerTimeRouter       = RouterMsg + "/get_server_time"

	GetMessageListReactionExtensionsRouter = RouterMsg + "/get_message_list_reaction_extensions"
	DeleteMessageReactionExtensionsRouter  = RouterMsg + "/delete_message_reaction_extensions"

	TencentCloudStorageCredentialRouter = "/third/tencent_cloud_storage_credential"
	AliOSSCredentialRouter              = "/third/ali_oss_credential"
	MinioStorageCredentialRouter        = "/third/minio_storage_credential"
	AwsStorageCredentialRouter          = "/third/aws_storage_credential"

	// group
	CreateGroupRouter                 = RouterGroup + "/create_group"
	SetGroupInfoRouter                = RouterGroup + "/set_group_info"
	JoinGroupRouter                   = RouterGroup + "/join_group"
	QuitGroupRouter                   = RouterGroup + "/quit_group"
	GetGroupsInfoRouter               = RouterGroup + "/get_groups_info"
	GetGroupMemberListRouter          = RouterGroup + "/get_group_member_list"
	GetGroupAllMemberListRouter       = RouterGroup + "/get_group_all_member_list"
	GetGroupMembersInfoRouter         = RouterGroup + "/get_group_members_info"
	InviteUserToGroupRouter           = RouterGroup + "/invite_user_to_group"
	GetJoinedGroupListRouter          = RouterGroup + "/get_joined_group_list"
	KickGroupMemberRouter             = RouterGroup + "/kick_group"
	TransferGroupRouter               = RouterGroup + "/transfer_group"
	GetRecvGroupApplicationListRouter = RouterGroup + "/get_recv_group_applicationList"
	GetSendGroupApplicationListRouter = RouterGroup + "/get_user_req_group_applicationList"
	AcceptGroupApplicationRouter      = RouterGroup + "/group_application_response"
	RefuseGroupApplicationRouter      = RouterGroup + "/group_application_response"
	DismissGroupRouter                = RouterGroup + "/dismiss_group"
	MuteGroupMemberRouter             = RouterGroup + "/mute_group_member"
	CancelMuteGroupMemberRouter       = RouterGroup + "/cancel_mute_group_member"
	MuteGroupRouter                   = RouterGroup + "/mute_group"
	CancelMuteGroupRouter             = RouterGroup + "/cancel_mute_group"
	SetGroupMemberNicknameRouter      = RouterGroup + "/set_group_member_nickname"
	SetGroupMemberInfoRouter          = RouterGroup + "/set_group_member_info"
	GetGroupAbstractInfoRouter        = RouterGroup + "/get_group_abstract_info"

	SaveGroupRouter         = RouterGroup + "/save_group"
	UnsaveGroupRouter       = RouterGroup + "/unsave_group"
	GetSavedGroupListRouter = RouterGroup + "/get_saved_group_list"

	SetReceiveMessageOptRouter         = "/conversation/set_receive_message_opt"
	GetReceiveMessageOptRouter         = "/conversation/get_receive_message_opt"
	GetAllConversationMessageOptRouter = "/conversation/get_all_conversation_message_opt"
	SetConversationOptRouter           = ConversationGroup + "/set_conversation"
	GetConversationsRouter             = ConversationGroup + "/get_conversations"
	GetAllConversationsRouter          = ConversationGroup + "/get_all_conversations"
	GetConversationRouter              = ConversationGroup + "/get_conversation"
	BatchSetConversationRouter         = ConversationGroup + "/batch_set_conversation"
	ModifyConversationFieldRouter      = ConversationGroup + "/modify_conversation_field"
	SetConversationsRouter             = ConversationGroup + "/set_conversations"

	// organization
	GetSubDepartmentRouter    = RouterOrganization + "/get_sub_department"
	GetDepartmentMemberRouter = RouterOrganization + "/get_department_member"
	ParseTokenRouter          = RouterAuth + "/parse_token"

	// super_group
	GetJoinedSuperGroupListRouter = RouterSuperGroup + "/get_joined_group_list"
	GetSuperGroupsInfoRouter      = RouterSuperGroup + "/get_groups_info"

	// third
	FcmUpdateTokenRouter = RouterThird + "/fcm_update_token"
	SetAppBadgeRouter    = RouterThird + "/set_app_badge"
	UploadLogsRouter     = RouterThird + "/logs/upload"
	GetRtcTokenRouter    = RouterThird + "/rtc/token"

	// auth
	GetUsersToken = RouterAuth + "/user_token"

	// club
	GetJoinedServerGroupListRouter     = RouterClub + "/get_joined_server_group_list"
	GetRecvServerApplicationListRouter = RouterClub + "/get_recv_server_application_list"
	GetSendServerApplicationListRouter = RouterClub + "/get_user_req_server_application_list"
	GetServerInfoListRouter            = RouterClub + "/get_servers_info"
	GetGroupCategoryListRouter         = RouterClub + "/get_categories"
	GetGroupsByServerRouter            = RouterClub + "/get_groups_by_server"
	GetJoinedServerListRouter          = RouterClub + "/get_joined_server_list"
	CreateServerRouter                 = RouterClub + "/create_server"
	DismissServerRouter                = RouterClub + "/dismiss_server"
	JoinServerRouter                   = RouterClub + "/join_server"
	QuitServerRouter                   = RouterClub + "/quit_server"
	KickServerMemberRouter             = RouterClub + "/kick_server_member"
	ServerApplicationResponseRouter    = RouterClub + "/server_application_response"

	DeleteServerGroupRouter = RouterClub + "/delete_server_group"

	SetServerMemberInfoRouter  = RouterClub + "/set_server_member_info"
	GetServerMembersInfoRouter = RouterClub + "/get_server_members_info"

	CronSetClearMsgJobRouter = RouterCron + "/set_clear_msg_job"
	CronGetClearMsgJobRouter = RouterCron + "/get_clear_msg_job"
)
const (
	RouterGroup        = "/group"
	ConversationGroup  = "/conversation"
	RouterOrganization = "/organization"
	RouterAuth         = "/auth"
	RouterSuperGroup   = "/super_group"
	RouterMsg          = "/msg"
	RouterThird        = "/third"
	RouterClub         = "/club"
	RouterCron         = "/cron"
)

const (
	ObjectPartLimit               = "/object/part_limit"
	ObjectPartSize                = "/object/part_size"
	ObjectInitiateMultipartUpload = "/object/initiate_multipart_upload"
	ObjectAuthSign                = "/object/auth_sign"
	ObjectCompleteMultipartUpload = "/object/complete_multipart_upload"
	ObjectAccessURL               = "/object/access_url"
)
