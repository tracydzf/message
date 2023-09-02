package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"msg/app/msg-web/msg_api/internal/svc"
	"msg/app/msg-web/msg_api/internal/types"
	"msg/app/msg-web/msg_rpc/pb/msg_rpc"
	"msg/common/werr"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req types.SendRequest) (resp *types.Response, err error) {
	variables, err := jsonx.Marshal(req.MessageParam.Variables)
	if err != nil {
		return nil, errors.Wrapf(werr.NewErrMsg("客户端参数错误"), "send err:%v", err)
	}

	extra, err := jsonx.Marshal(req.MessageParam.Extra)
	if err != nil {
		return nil, errors.Wrapf(werr.NewErrMsg("客户端参数错误"), "send err:%v", err)
	}

	send, err := l.svcCtx.SendRpc.Send(l.ctx, &msg_rpc.SendRequest{
		Code:              req.Code,
		MessageTemplateId: req.MessageTemplateId,
		MessageParam: &msg_rpc.MessageParam{
			Receiver:  req.MessageParam.Receiver,
			Variables: string(variables),
			Extra:     string(extra),
		},
	})
	if err != nil {
		return nil, errors.Wrapf(werr.NewErrMsg("发送消息错误"), "send err:%v", err)
	}
	return &types.Response{Message: send.Code}, err
}
