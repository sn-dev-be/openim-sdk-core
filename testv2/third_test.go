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

package testv2

import (
	"testing"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

func Test_UpdateFcmToken(t *testing.T) {
	err := open_im_sdk.UserForSDK.Third().UpdateFcmToken(
		ctx,
		"e1bgoVRIT2GZpXECRy9Lw6:APA91bHRiN6Muyvz-_brNlR_HHjJL0Jm_7q2LnmxpH0fcQeo_lhqW4zbeaWpS-X0e_IjiVgjiwz1PkwBZlJPC_Tb29nYgYSpUAL5EmB8AAJKUKqOa-OWEooaTWIbPWrvaxoWmAGfSO6u",
		100000,
	)
	if err != nil {
		t.Fatal(err)
	}
}
