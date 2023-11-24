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

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
)

func SetGroupListener(callback open_im_sdk_callback.OnGroupListener) {
	if callback == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetGroupListener(callback)
}

func SetConversationListener(listener open_im_sdk_callback.OnConversationListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetConversationListener(listener)
}
func SetAdvancedMsgListener(listener open_im_sdk_callback.OnAdvancedMsgListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetAdvancedMsgListener(listener)
}
func SetBatchMsgListener(listener open_im_sdk_callback.OnBatchMsgListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetBatchMsgListener(listener)
}

func SetUserListener(listener open_im_sdk_callback.OnUserListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetUserListener(listener)
}

func SetFriendListener(listener open_im_sdk_callback.OnFriendshipListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetFriendListener(listener)
}

func SetCustomBusinessListener(listener open_im_sdk_callback.OnCustomBusinessListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetBusinessListener(listener)
}

func SetMessageKvInfoListener(listener open_im_sdk_callback.OnMessageKvInfoListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetMessageKvInfoListener(listener)
}

func SetSignalingListener(listener open_im_sdk_callback.OnSignalingListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetSignalingListener(listener)
}

func SetClubListener(listener open_im_sdk_callback.OnClubListener) {
	if listener == nil || UserForSDK == nil {
		return
	}
	UserForSDK.SetClubListener(listener)
}
