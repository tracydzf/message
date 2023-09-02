package action

import "msg/common/werr"

var (
	sendErr         = werr.NewErrMsg("发送消息错误")
	clientParamsErr = werr.NewErrMsg("客户端参数错误")
)
