package action

import (
	"context"
	"github.com/pkg/errors"
	"msg/app/msg-common/enums/channelType"
	"msg/app/msg-common/enums/idType"
	"msg/app/msg-common/taskUtil"
	"msg/app/msg-common/types"
	"regexp"
)

type AfterParamCheckAction struct {
}

func NewAfterParamCheckAction() *AfterParamCheckAction {
	return &AfterParamCheckAction{}
}

func (p AfterParamCheckAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {

	// 1. 过滤掉不合法的手机号

	if sendTaskModel.TaskInfo[0].IdType == idType.Phone && sendTaskModel.TaskInfo[0].SendChannel == channelType.Sms {
		var newTask []types.TaskInfo
		for _, item := range sendTaskModel.TaskInfo {
			for _, tel := range item.Receiver {
				matched, _ := regexp.Match(taskUtil.PhoneRegex, []byte(tel))
				if matched {
					newTask = append(newTask, item)
				}
			}
		}
		if len(newTask) <= 0 {
			return errors.Wrapf(sendErr, "AfterParamCheckAction sendTaskModel:%v", sendTaskModel)
		}
		sendTaskModel.TaskInfo = newTask
	}

	return nil
}
