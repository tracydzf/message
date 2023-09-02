package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"msg/common/dbx"
)

type Config struct {
	service.ServiceConf
	Mysql      dbx.DbConfig
	RedisConf  RedisConf
	CacheRedis cache.CacheConf // redis缓存
}

type RedisConf struct {
	Address  string
	Password string
	//DB       int
}
