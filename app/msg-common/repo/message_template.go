package repo

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"msg/app/msg-common/dao/model"
	"msg/common/dbx"
	"msg/common/gormc"
)

type MessageTemplateRepo struct {
	cache gormc.CachedConn
}

func NewMessageTemplateRepo(c cache.CacheConf, opts ...cache.Option) *MessageTemplateRepo {
	return &MessageTemplateRepo{
		cache: gormc.NewConn(c, opts...),
	}
}

func (a *MessageTemplateRepo) getModel(ctx context.Context) *gorm.DB {
	return dbx.GetDb(ctx).Model(&model.MessageTemplate{})
}

func (a *MessageTemplateRepo) One(ctx context.Context, id int64) (item model.MessageTemplate, err error) {
	key := fmt.Sprintf("messagetemplate_%d", id)
	err = a.cache.QueryRowCtx(ctx, &item, key, func(ctx context.Context, v interface{}) error {
		return a.getModel(ctx).Where("id", id).Limit(1).Find(&item).Error
	})
	return
}
