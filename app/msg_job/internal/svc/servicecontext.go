package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"msg/app/msg_job/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  initRedis(c),
	}
}

func initRedis(c config.Config) *redis.Redis {
	fmt.Println("connect Redis ...")
	conf := redis.RedisConf{
		Host: c.RedisConf.Address,
		Type: "node",
		Pass: c.RedisConf.Address,
		Tls:  false,
	}

	newRedis, err := redis.NewRedis(conf)
	if err != nil {
		return nil
	}

	//db := redis.NewClient(&redis.Options{
	//	Addr:     c.RedisConf.Address,
	//	Password: c.RedisConf.Password,
	//	//DB:       c.DBList.Redis.DB,
	//	//超时
	//	ReadTimeout:  2 * time.Second,
	//	WriteTimeout: 2 * time.Second,
	//	PoolTimeout:  3 * time.Second,
	//})
	//_, err := db.Ping(context.Background()).Result()
	//if err != nil {
	//	fmt.Println("connect Redis failed")
	//	panic(err)
	//}
	//fmt.Println("connect Redis success")
	return newRedis
}
