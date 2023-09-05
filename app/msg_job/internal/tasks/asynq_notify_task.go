package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"msg/app/msg-common/types"
)

const MainTask = "main:task"

func (l *AsynqServer) MainTask(ctx context.Context, t *asynq.Task) error {

	var taskList []types.TaskInfo
	_ = jsonx.Unmarshal(t.Payload(), &taskList)
	for _, taskInfo := range taskList {
		logx.WithContext(ctx).Infow("消息接收成功,开始消费", logx.Field("task_info", taskInfo))
		//channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		//msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		err := NewTask(taskInfo, l.svcCtx)
		if err != nil {
			logx.WithContext(ctx).Errorw("submit task err",
				logx.Field("内容", string(t.Payload())),
				logx.Field("err", err))
		}
	}
	return nil

}
