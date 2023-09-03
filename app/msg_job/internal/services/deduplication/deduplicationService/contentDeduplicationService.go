package deduplicationService

import (
	"context"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/services/deduplication/limit"
	"msg/app/msg_job/internal/services/deduplication/srv"
	"msg/app/msg_job/internal/services/deduplication/structs"
	"msg/app/msg_job/internal/svc"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx}
}

func (c contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return srv.NewContentDeduplicationService(c.svcCtx, limit.NewSimpleLimitService(c.svcCtx)).
		Deduplication(ctx, taskInfo, param)
}
