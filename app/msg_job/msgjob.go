package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"msg/app/msg_job/internal/config"
	"msg/app/msg_job/internal/handler"
	"msg/app/msg_job/internal/svc"
	"msg/app/msg_job/internal/tasks"
	"msg/common/dbx"
)

var configFile = flag.String("f", "etc/msgjob-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	logx.DisableStat()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	serviceGroup.Add(tasks.NewAsynqServer(ctx, svcContext))

	handler.SetUp(svcContext)
	dbx.InitDb(c.Mysql)

	serviceGroup.Start()

}
