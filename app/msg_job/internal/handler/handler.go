package handler

import (
	"context"
	"msg/app/msg-common/enums/channelType"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/svc"
	"sync"
)

var (
	once          sync.Once
	handlerHolder map[int]IHandler
)

const flowControlEmail = "flow_control_email"

// SetUp 初始化所有handler
func SetUp(svcCtx *svc.ServiceContext) {
	once.Do(func() {
		handlerHolder = map[int]IHandler{
			channelType.Email: NewEmailHandler(svcCtx),
		}
	})

}

func GetHandler(sendChannel int) IHandler {
	return handlerHolder[sendChannel]
}

type BaseHandler struct {
}

func (b BaseHandler) Limit(ctx context.Context, taskInfo types.TaskInfo) bool {
	return true
}
