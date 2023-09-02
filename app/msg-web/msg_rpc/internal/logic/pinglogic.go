package logic

import (
	"context"

	"msg/app/msg-web/msg_rpc/internal/svc"
	"msg/app/msg-web/msg_rpc/pb/msg_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *msg_rpc.PingRequest) (*msg_rpc.PingResponse, error) {
	// todo: add your logic here and delete this line

	return &msg_rpc.PingResponse{Pong: "pong"}, nil
}
