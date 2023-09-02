package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"msg/app/msg-web/msg_rpc/internal/svc"
	"msg/app/msg-web/msg_rpc/pb/msg_rpc"
)

type BatchSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchSendLogic {
	return &BatchSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchSendLogic) BatchSend(in *msg_rpc.BatchSendRequest) (*msg_rpc.SendResponse, error) {
	// todo: add your logic here and delete this line

	return &msg_rpc.SendResponse{}, nil
}
