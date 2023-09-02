package handler

import (
	"context"
	"msg/app/msg-common/types"
)

type ILimit interface {
	Limit(ctx context.Context, taskInfo types.TaskInfo) bool
}

type IHandler interface {
	ILimit
	DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error)
}
