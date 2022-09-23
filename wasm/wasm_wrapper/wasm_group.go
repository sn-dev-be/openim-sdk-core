package wasm_wrapper

import (
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/wasm/event_listener"
	"syscall/js"
)

func GetGroupsInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), commonFunc)
	checker(callback, &args, 2)
	callback.EventData().SetOperationID(args[0].String())
	open_im_sdk.GetGroupsInfo(callback, args[0].String(), args[1].String())
	return nil
}
