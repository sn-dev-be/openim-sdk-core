package conversation_msg

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"

	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"
)

func (c *Conversation) doRedPacketMsgStatusChange(ctx context.Context, msg *sdkws.MsgData) {
	var tips sdkws.RedPacketTips
	if err := utils.UnmarshalNotificationElem(msg.Content, &tips); err != nil {
		log.ZError(ctx, "unmarshal failed", err, "msg", msg)
		return
	}
	log.ZDebug(ctx, "do redPacket message", "tips", &tips)

	redPacketMsg, err := c.db.GetMessage(ctx, tips.ConversationID, tips.ClientMsgID)
	if err != nil {
		log.ZError(ctx, "GetMessageBySeq failed", err, "tips", &tips)
		return
	}

	elem := sdkws.RedPacketElem{}
	utils.JsonStringToStruct(redPacketMsg.Content, &elem)
	if tips.ClaimUser != nil && c.loginUserID == tips.ClaimUser.UserID {
		elem.Status = constant.RedPacketClaimedByUser
	} else if elem.Status != constant.RedPacketClaimedByUser {
		elem.Status = tips.Status
	}

	redPacketMsg.Content = utils.StructToJsonString(&elem)
	if err = c.db.UpdateMessage(ctx, tips.ConversationID, redPacketMsg); err != nil {
		log.ZError(ctx, "UpdateMessage err", err, "conversationID", tips.ConversationID, "message", redPacketMsg)
	}

	c.msgListener.OnRecvRedPacketStatusChanged(utils.StructToJsonString(&tips))
}
