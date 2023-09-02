package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"msg/app/msg-common/interfaces"
	"msg/app/msg-common/types"
	"msg/app/msg-web/msg_rpc/internal/process"
	action2 "msg/app/msg-web/msg_rpc/internal/process/action"
	"msg/app/msg-web/msg_rpc/internal/svc"
	"msg/app/msg-web/msg_rpc/pb/msg_rpc"
	"msg/common/utils"
	"msg/common/werr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(in *msg_rpc.SendRequest) (*msg_rpc.SendResponse, error) {
	if in.MessageParam == nil {
		return nil, errors.Wrapf(werr.NewErrMsg("客户端参数错误"), "in:%v", in)
	}
	variables := make(map[string]interface{})
	extra := make(map[string]interface{})
	var err error
	err = jsonx.Unmarshal([]byte(in.MessageParam.Variables), &variables)
	if err != nil {
		return nil, errors.Wrapf(werr.NewErrMsg("客户端参数错误"), "in:%v err:%v", in, err)
	}
	err = jsonx.Unmarshal([]byte(in.MessageParam.Extra), &extra)
	if err != nil {
		return nil, errors.Wrapf(werr.NewErrMsg("客户端参数错误"), "in:%v err:%v", in, err)
	}
	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList: []types.MessageParam{
			{
				Receiver:  in.MessageParam.Receiver,
				Variables: variables,
				Extra:     extra,
			},
		},
	}
	utils.Print(sendTaskModel)
	businessProcess := process.NewBusinessProcess()
	_ = businessProcess.AddProcess([]interfaces.Process{
		action2.NewPreParamCheckAction(), //前置参数校验
		//action2.NewAssembleAction(l.svcCtx), //拼装参数
		action2.NewAfterParamCheckAction(), //后置参数检查
		action2.NewSendMqAction(l.svcCtx),  //发送到mq
	}...)

	err = businessProcess.Process(l.ctx, sendTaskModel)
	if err != nil {
		return nil, err
	}
	return &msg_rpc.SendResponse{}, nil
}
