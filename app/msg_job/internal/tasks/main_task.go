package tasks

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/handler"
	"msg/app/msg_job/internal/services/deduplication"
	"msg/app/msg_job/internal/svc"
)

type Task struct {
	TaskInfo types.TaskInfo
	svcCtx   *svc.ServiceContext
}

func NewTask(taskInfo types.TaskInfo, svcCtx *svc.ServiceContext) *Task {
	return &Task{TaskInfo: taskInfo, svcCtx: svcCtx}
}

type TaskRun interface {
	Run(ctx context.Context)
}

func (t Task) Run(ctx context.Context) {
	if len(t.TaskInfo.Receiver) > 0 {
		deduplication.NewDeduplicationRuleService(t.svcCtx).Duplication(ctx, &t.TaskInfo)
	}

	// 3. 真正发送消息
	if len(t.TaskInfo.Receiver) > 0 {
		h := handler.GetHandler(t.TaskInfo.SendChannel)
		for {
			if h.Limit(ctx, t.TaskInfo) {
				err := h.DoHandler(ctx, t.TaskInfo)
				if err != nil {
					logx.Errorw("DoHandler err", logx.Field("task_info", t.TaskInfo), logx.Field("err", err))
				}
				return
			}
		}
	}

}
