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

package conversation_msg

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"
)

func (c *Conversation) doModifyMsg(ctx context.Context, msg *sdkws.MsgData) {
	var tips sdkws.ModifyMessageTips
	if err := utils.UnmarshalNotificationElem(msg.Content, &tips); err != nil {
		log.ZError(ctx, "unmarshal failed", err, "msg", msg)
		return
	}
	log.ZDebug(ctx, "do modifyMessage", "tips", &tips)
	c.modifyMessage(ctx, &tips)
}

func (c *Conversation) modifyMessage(ctx context.Context, tips *sdkws.ModifyMessageTips) {
	msg, err := c.db.GetMessageBySeq(ctx, tips.ConversationID, tips.Seq)
	if err != nil {
		log.ZError(ctx, "GetMessageBySeq failed", err, "tips", &tips)
		return
	}

	switch tips.ModifyType {
	case constant.MsgModifyServerRequestStatus:
	case constant.MsgModifyRedPacketStatus:
	}

	msg.Content = utils.StructToJsonString(sdkws.NotificationElem{Detail: tips.Content})
	if err = c.db.UpdateMessage(ctx, tips.ConversationID, msg); err != nil {
		log.ZError(ctx, "UpdateMessage err", err, "conversationID", tips.ConversationID, "message", msg)
	}
}
