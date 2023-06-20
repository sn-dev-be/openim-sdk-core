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

//go:build js && wasm
// +build js,wasm

package indexdb

import (
	"context"
	"encoding/json"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/wasm/indexdb/temp_struct"
)

type LocalChatLogs struct {
	loginUserID string
}

// NewLocalChatLogs creates a new LocalChatLogs
func NewLocalChatLogs(loginUserID string) *LocalChatLogs {
	return &LocalChatLogs{loginUserID: loginUserID}
}

// GetMessage gets the message from the database
func (i *LocalChatLogs) GetMessage(ctx context.Context, conversationID, clientMsgID string) (*model_struct.LocalChatLog, error) {
	msg, err := Exec(conversationID, clientMsgID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msg.(string); ok {
			result := model_struct.LocalChatLog{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetSendingMessageList gets the list of messages that are being sent
func (i *LocalChatLogs) GetSendingMessageList(ctx context.Context) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// UpdateMessage updates the message in the database
func (i *LocalChatLogs) UpdateMessage(ctx context.Context, conversationID string, c *model_struct.LocalChatLog) error {
	if c.ClientMsgID == "" {
		return PrimaryKeyNull
	}
	tempLocalChatLog := temp_struct.LocalChatLog{
		ServerMsgID:          c.ServerMsgID,
		SendID:               c.SendID,
		RecvID:               c.RecvID,
		SenderPlatformID:     c.SenderPlatformID,
		SenderNickname:       c.SenderNickname,
		SenderFaceURL:        c.SenderFaceURL,
		SessionType:          c.SessionType,
		MsgFrom:              c.MsgFrom,
		ContentType:          c.ContentType,
		Content:              c.Content,
		IsRead:               c.IsRead,
		Status:               c.Status,
		Seq:                  c.Seq,
		SendTime:             c.SendTime,
		CreateTime:           c.CreateTime,
		AttachedInfo:         c.AttachedInfo,
		Ex:                   c.Ex,
		IsReact:              c.IsReact,
		IsExternalExtensions: c.IsExternalExtensions,
		MsgFirstModifyTime:   c.MsgFirstModifyTime,
	}
	_, err := Exec(conversationID, c.ClientMsgID, utils.StructToJsonString(tempLocalChatLog))
	return err
}

// UpdateMessageStatus updates the message status in the database
func (i *LocalChatLogs) BatchInsertMessageList(ctx context.Context, conversationID string, messageList []*model_struct.LocalChatLog) error {
	_, err := Exec(conversationID, utils.StructToJsonString(messageList))
	return err
}

// InsertMessage inserts a message into the local chat log.
func (i *LocalChatLogs) InsertMessage(ctx context.Context, conversationID string, message *model_struct.LocalChatLog) error {
	_, err := Exec(conversationID, utils.StructToJsonString(message))
	return err
}

// UpdateColumnsMessageList updates multiple columns of a message in the local chat log.
func (i *LocalChatLogs) UpdateColumnsMessageList(ctx context.Context, clientMsgIDList []string, args map[string]interface{}) error {
	_, err := Exec(utils.StructToJsonString(clientMsgIDList), args)
	return err
}

// UpdateColumnsMessage updates a column of a message in the local chat log.
func (i *LocalChatLogs) UpdateColumnsMessage(ctx context.Context, clientMsgID string, args map[string]interface{}) error {
	_, err := Exec(clientMsgID, utils.StructToJsonString(args))
	return err
}

// DeleteAllMessage deletes all messages from the local chat log.
func (i *LocalChatLogs) DeleteAllMessage(ctx context.Context) error {
	_, err := Exec()
	return err
}

// UpdateMessageStatusBySourceID updates the status of a message in the local chat log by its source ID.
func (i *LocalChatLogs) UpdateMessageStatusBySourceID(ctx context.Context, sourceID string, status, sessionType int32) error {
	_, err := Exec(sourceID, status, sessionType, i.loginUserID)
	return err
}

// UpdateMessageTimeAndStatus updates the time and status of a message in the local chat log.
func (i *LocalChatLogs) UpdateMessageTimeAndStatus(ctx context.Context, conversationID, clientMsgID string, serverMsgID string, sendTime int64, status int32) error {
	_, err := Exec(conversationID, clientMsgID, serverMsgID, sendTime, status)
	return err
}

// GetMessageList retrieves a list of messages from the local chat log.
func (i *LocalChatLogs) GetMessageList(ctx context.Context, conversationID string, count int, startTime int64, isReverse bool) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(conversationID, count, startTime, isReverse, i.loginUserID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetMessageListNoTime retrieves a list of messages from the local chat log without specifying a start time.
func (i *LocalChatLogs) GetMessageListNoTime(ctx context.Context, conversationID string, count int, isReverse bool) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(conversationID, count, isReverse, i.loginUserID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// UpdateSingleMessageHasRead updates the hasRead field of a single message in the local chat log.
func (i *LocalChatLogs) UpdateSingleMessageHasRead(ctx context.Context, sendID string, msgIDList []string) error {
	_, err := Exec(sendID, utils.StructToJsonString(msgIDList))
	return err
}

// SearchMessageByContentType searches for messages in the local chat log by content type.
func (i *LocalChatLogs) SearchMessageByContentType(ctx context.Context, contentType []int, conversationID string, startTime, endTime int64, offset, count int) (messages []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(conversationID, utils.StructToJsonString(contentType), startTime, endTime, offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				messages = append(messages, &v1)
			}
			return messages, err
		} else {
			return nil, ErrType
		}
	}
}

//func (i *LocalChatLogs) SuperGroupSearchMessageByContentType(ctx context.Context, contentType []int, sourceID string, startTime, endTime int64, sessionType, offset, count int) (messages []*model_struct.LocalChatLog, err error) {
//	msgList, err := Exec(utils.StructToJsonString(contentType), sourceID, startTime, endTime, sessionType, offset, count)
//	if err != nil {
//		return nil, err
//	} else {
//		if v, ok := msgList.(string); ok {
//			var temp []model_struct.LocalChatLog
//			err := utils.JsonStringToStruct(v, &temp)
//			if err != nil {
//				return nil, err
//			}
//			for _, v := range temp {
//				v1 := v
//				messages = append(messages, &v1)
//			}
//			return messages, err
//		} else {
//			return nil, ErrType
//		}
//	}
//}

// SearchMessageByKeyword searches for messages in the local chat log by keyword.
func (i *LocalChatLogs) SearchMessageByContentTypeAndKeyword(ctx context.Context, contentType []int, conversationID string, keywordList []string, keywordListMatchType int, startTime, endTime int64) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(conversationID, utils.StructToJsonString(contentType), utils.StructToJsonString(keywordList), keywordListMatchType, startTime, endTime)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// MessageIfExists check if message exists
func (i *LocalChatLogs) MessageIfExists(ctx context.Context, clientMsgID string) (bool, error) {
	isExist, err := Exec(clientMsgID)
	if err != nil {
		return false, err
	} else {
		if v, ok := isExist.(bool); ok {
			return v, nil
		} else {
			return false, ErrType
		}
	}
}

// IsExistsInErrChatLogBySeq check if message exists in error chat log by seq
func (i *LocalChatLogs) IsExistsInErrChatLogBySeq(ctx context.Context, seq int64) bool {
	isExist, err := Exec(seq)
	if err != nil {
		return false
	} else {
		if v, ok := isExist.(bool); ok {
			return v
		} else {
			return false
		}
	}
}

// MessageIfExistsBySeq check if message exists by seq
func (i *LocalChatLogs) MessageIfExistsBySeq(ctx context.Context, seq int64) (bool, error) {
	isExist, err := Exec(seq)
	if err != nil {
		return false, err
	} else {
		if v, ok := isExist.(bool); ok {
			return v, nil
		} else {
			return false, ErrType
		}
	}
}

// GetMultipleMessage gets multiple messages from the local chat log.
func (i *LocalChatLogs) GetMultipleMessage(ctx context.Context, msgIDList []string) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(utils.StructToJsonString(msgIDList))
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetLostMsgSeqList gets lost message seq list.
func (i *LocalChatLogs) GetLostMsgSeqList(ctx context.Context, minSeqInSvr uint32) (result []uint32, err error) {
	l, err := Exec(minSeqInSvr)
	if err != nil {
		return nil, err
	} else {
		if v, ok := l.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetTestMessage gets test message.
func (i *LocalChatLogs) GetTestMessage(ctx context.Context, seq uint32) (*model_struct.LocalChatLog, error) {
	msg, err := Exec(seq)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msg.(model_struct.LocalChatLog); ok {
			return &v, err
		} else {
			return nil, ErrType
		}
	}
}

// Update the sender's nickname in the chat logs
func (i *LocalChatLogs) UpdateMsgSenderNickname(ctx context.Context, sendID, nickname string, sType int) error {
	_, err := Exec(sendID, nickname, sType)
	return err
}

// Update the sender's face URL in the chat logs
func (i *LocalChatLogs) UpdateMsgSenderFaceURL(ctx context.Context, sendID, faceURL string, sType int) error {
	_, err := Exec(sendID, faceURL, sType)
	return err
}

// Update the sender's face URL and nickname in the chat logs
func (i *LocalChatLogs) UpdateMsgSenderFaceURLAndSenderNickname(ctx context.Context, conversationID, sendID, faceURL, nickname string) error {
	_, err := Exec(conversationID, sendID, faceURL, nickname)
	return err
}

// Get the message sequence number by client message ID
func (i *LocalChatLogs) GetMsgSeqByClientMsgID(ctx context.Context, clientMsgID string) (uint32, error) {
	result, err := Exec(clientMsgID)
	if err != nil {
		return 0, err
	}
	if v, ok := result.(float64); ok {
		return uint32(v), nil
	}
	return 0, ErrType
}

// Search all messages by content type
func (i *LocalChatLogs) SearchAllMessageByContentType(ctx context.Context, contentType int) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(contentType)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []*model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// Get the message sequence number list by group ID
func (i *LocalChatLogs) GetMsgSeqListByGroupID(ctx context.Context, groupID string) (result []uint32, err error) {
	l, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := l.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// Get the message sequence number list by peer user ID
func (i *LocalChatLogs) GetMsgSeqListByPeerUserID(ctx context.Context, userID string) (result []uint32, err error) {
	l, err := Exec(userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := l.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// Get the message sequence number list by self user ID
func (i *LocalChatLogs) GetMsgSeqListBySelfUserID(ctx context.Context, userID string) (result []uint32, err error) {
	l, err := Exec(userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := l.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// Get the abnormal message sequence number
func (i *LocalChatLogs) GetAbnormalMsgSeq(ctx context.Context) (int64, error) {
	result, err := Exec()
	if err != nil {
		return 0, err
	}
	if v, ok := result.(float64); ok {
		return int64(v), nil
	}
	return 0, ErrType
}

// Get the list of abnormal message sequence numbers
func (i *LocalChatLogs) GetAbnormalMsgSeqList(ctx context.Context) (result []int64, err error) {
	l, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := l.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// Batch insert exception messages into the chat logs
func (i *LocalChatLogs) BatchInsertExceptionMsg(ctx context.Context, MessageList []*model_struct.LocalErrChatLog) error {
	_, err := Exec(utils.StructToJsonString(MessageList))
	return err
}

// Update the message status to read in the chat logs
func (i *LocalChatLogs) UpdateGroupMessageHasRead(ctx context.Context, msgIDList []string, sessionType int32) error {
	_, err := Exec(utils.StructToJsonString(msgIDList), sessionType)
	return err
}

// Get the message by message ID
func (i *LocalChatLogs) SearchMessageByKeyword(ctx context.Context, contentType []int, keywordList []string, keywordListMatchType int, conversationID string, startTime, endTime int64, offset, count int) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(conversationID, utils.StructToJsonString(contentType), utils.StructToJsonString(keywordList), keywordListMatchType, startTime, endTime, offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetSuperGroupAbnormalMsgSeq get super group abnormal msg seq
func (i *LocalChatLogs) GetSuperGroupAbnormalMsgSeq(ctx context.Context, groupID string) (uint32, error) {
	isExist, err := Exec(groupID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := isExist.(uint32); ok {
			return v, nil
		} else {
			return 0, ErrType
		}
	}
}

// GetAlreadyExistSeqList get already exist seq list
func (i *LocalChatLogs) GetAlreadyExistSeqList(ctx context.Context, conversationID string, lostSeqList []int64) ([]int64, error) {
	seqList, err := Exec(conversationID, lostSeqList)
	if err != nil {
		return nil, err
	} else {
		if v, ok := seqList.([]int64); ok {
			result := make([]int64, len(v))
			for i, value := range v {
				result[i] = value.(int64)
			}
			return result, nil
		} else {
			return nil, ErrType
		}
	}
}

// GetMessagesBySeq get message by seq
func (i *LocalChatLogs) GetMessageBySeq(ctx context.Context, conversationID string, seq int64) (*model_struct.LocalChatLog, error) {
	msg, err := Exec(conversationID, seq)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msg.(string); ok {
			result := model_struct.LocalChatLog{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

// UpdateMessageBySeq update message
func (i *LocalChatLogs) UpdateMessageBySeq(ctx context.Context, conversationID string, c *model_struct.LocalChatLog) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = Exec(conversationID, c.Seq, string(data))
	return err
}

func (i *LocalChatLogs) UpdateMessageByClientMsgID(ctx context.Context, clientMsgID string, data map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func (i *LocalChatLogs) MarkConversationMessageAsRead(ctx context.Context, conversationID string, msgIDs []string) (rowsAffected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *LocalChatLogs) MarkConversationMessageAsReadBySeqs(ctx context.Context, conversationID string, seqs []int64) (rowsAffected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *LocalChatLogs) GetUnreadMessage(ctx context.Context, conversationID string) (result []*model_struct.LocalChatLog, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *LocalChatLogs) MarkConversationAllMessageAsRead(ctx context.Context, conversationID string) (rowsAffected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *LocalChatLogs) GetMessagesByClientMsgIDs(ctx context.Context, conversationID string, msgIDs []string) ([]*model_struct.LocalChatLog, error) {
	msgs, err := Exec(conversationID, msgIDs)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgs.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			result := make([]*model_struct.LocalChatLog, len(temp))
			for i, v := range temp {
				v1 := v
				result[i] = &v1
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetMessagesBySeqs gets messages by seqs
func (i *LocalChatLogs) GetMessagesBySeqs(ctx context.Context, conversationID string, seqs []int64) (result []*model_struct.LocalChatLog, err error) {
	msgs, err := Exec(conversationID, seqs)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgs.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

// GetConversationNormalMsgSeq gets the maximum seq of the session
func (i *LocalChatLogs) GetConversationNormalMsgSeq(ctx context.Context, conversationID string) (int64, error) {
	seq, err := Exec(conversationID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(float64); ok {
			var result int64
			result = int64(v)
			return result, err
		} else {
			return 0, ErrType
		}
	}
}

// GetConversationPeerNormalMsgSeq gets the maximum seq of the peer in the session
func (i *LocalChatLogs) GetConversationPeerNormalMsgSeq(ctx context.Context, conversationID string) (int64, error) {
	seq, err := Exec(conversationID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(int64); ok {
			return v, nil
		} else {
			return 0, ErrType
		}
	}
}

// GetConversationAbnormalMsgSeq gets the maximum abnormal seq of the session
func (i *LocalChatLogs) GetConversationAbnormalMsgSeq(ctx context.Context, groupID string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteConversationAllMessages deletes all messages of the session
func (i *LocalChatLogs) DeleteConversationAllMessages(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// MarkDeleteConversationAllMessages marks all messages of the session as deleted
func (i *LocalChatLogs) MarkDeleteConversationAllMessages(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteConversationMsgs deletes messages of the session
func (i *LocalChatLogs) DeleteConversationMsgs(ctx context.Context, conversationID string, msgIDs []string) error {
	_, err := Exec(conversationID, msgIDs)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteConversationMsgsBySeqs deletes messages of the session
func (i *LocalChatLogs) DeleteConversationMsgsBySeqs(ctx context.Context, conversationID string, seqs []int64) error {
	_, err := Exec(conversationID, seqs)
	if err != nil {
		return err
	} else {
		return nil
	}
}
