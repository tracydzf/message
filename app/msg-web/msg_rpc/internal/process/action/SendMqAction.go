package action

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"msg/app/msg-common/types"
	"msg/app/msg-web/msg_rpc/internal/svc"
)

type SendMqAction struct {
	svcCtx *svc.ServiceContext
}

func NewSendMqAction(svcCtx *svc.ServiceContext) *SendMqAction {
	return &SendMqAction{svcCtx: svcCtx}
}

func (p SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {
	marshal, err := jsonx.Marshal(sendTaskModel.TaskInfo)
	fmt.Println(marshal)
	if err != nil {
		return err
	}
	//channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	//msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	//return p.svcCtx.MqClient.Publish(marshal, taskUtil.GetMqKey(channel, msgType))
	return nil
}
