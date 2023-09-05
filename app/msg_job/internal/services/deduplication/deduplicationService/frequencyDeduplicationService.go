package deduplicationService

import (
	"context"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/services/deduplication/limit"
	"msg/app/msg_job/internal/services/deduplication/srv"
	"msg/app/msg_job/internal/services/deduplication/structs"
	"msg/app/msg_job/internal/svc"
)

type frequencyDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &frequencyDeduplicationService{svcCtx: svcCtx}
}

func (c frequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return srv.NewFrequencyDeduplicationService(c.svcCtx, limit.NewSlideWindowLimitService(c.svcCtx)).
		Deduplication(ctx, taskInfo, param)
}
