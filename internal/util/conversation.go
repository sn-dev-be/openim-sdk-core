package util

import (
	"strings"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/constant"
)

func ParseConversationID(s string) (int, string) {
	keyToNum := map[string]int{
		"si":  constant.SingleChatType,
		"sg":  constant.GroupChatType,
		"svg": constant.ServerGroupChatType,
		"sn":  constant.NotificationChatType,
	}
	parts := strings.Split(s, "_")
	k, v := parts[0], parts[1:]
	return keyToNum[k], strings.Join(v, "_")
}
