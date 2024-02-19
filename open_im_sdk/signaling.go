package open_im_sdk

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
)

func SignalingInvite(callback open_im_sdk_callback.Base, operationID string, conversationID, userID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingInvite, conversationID, userID, channelID)
}

func SignalingAccept(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingAccept, conversationID, channelID)
}

func SignalingReject(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingReject, conversationID, channelID)
}

func SignalingJoin(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingJoin, conversationID, channelID)
}

func SignalingCancel(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string, cancelUserID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingCancel, conversationID, channelID, cancelUserID)
}

func SignalingHungUp(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string, hungUpType int32) {
	call(callback, operationID, UserForSDK.Signaling().SignalingHungUp, conversationID, channelID, hungUpType)
}

func SignalingClose(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SignalingClose, conversationID, channelID)
}

func UpdateMichoneStatus(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string, status int32) {
	call(callback, operationID, UserForSDK.Signaling().MichoneStatusChange, conversationID, channelID, status)
}

func UpdateSpeakStatuse(callback open_im_sdk_callback.Base, operationID string, conversationID string, channelID string) {
	call(callback, operationID, UserForSDK.Signaling().SpeakStatusChange, conversationID, channelID)
}

func GetRtcToken(callback open_im_sdk_callback.Base, operationID string, userID string, channelID string, roleType int32) {
	call(callback, operationID, UserForSDK.Signaling().GetRtcToken, userID, channelID, roleType)
}
