package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"msg/app/msg-web/msg_api/internal/config"
	"msg/app/msg-web/msg_rpc/msg"
	"msg/app/msg-web/msg_rpc/pb/msg_rpc"
)

type ServiceContext struct {
	Config  config.Config
	SendRpc msg_rpc.MsgClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SendRpc: msg.NewMsg(zrpc.MustNewClient(c.SendRpcConf)),
	}
}
